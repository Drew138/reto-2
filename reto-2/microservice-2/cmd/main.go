package main

import (
	"asalaza5-st0263/reto-2/microservice-2/internal/files"
	"bytes"
	"encoding/json"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Request struct {
	Name      string `json:"name"`
	Query     string `json:"query"`
	QueueName string `json:"queue_name"`
}

func parseRequest(request Request, directory string, ch *amqp.Channel) []string {
	if request.Name == "list" {
		return files.ListFiles(directory)
	} else if request.Name == "search" {
		return files.SearchFiles(directory, request.Query)
	} else {
		return []string{"Invalid request"}
	}
}

func publish(request Request, channel *amqp.Channel, body []byte) error {
	args := make(amqp.Table)
	args["x-expires"] = int32(20000)
	responseQueue, err := channel.QueueDeclare(
		request.QueueName, // name
		false,             // durable
		true,              // delete when unused
		false,             // exclusive
		false,             // no-wait
		args,              // arguments
	)
	if err != nil {
		return err
	}
	err = channel.Publish(
		"",                 // exchange
		responseQueue.Name, // routing key
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			ContentType: "encoding/json",
			Body:        body,
		})
	log.Printf(" [x] Sent %s\n", body)
	return err
}

func main() {
	time.Sleep(60 * time.Second)
	rabbitmqURL := os.Getenv("RABBITMQ_URL")
	directory := os.Getenv("MICROSERVICE_TWO_DIRECTORY")

	amqpConn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		panic(err)
	}
	defer amqpConn.Close()

	channel, err := amqpConn.Channel()
	requestQueue, err := channel.QueueDeclare(
		"request",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	messages, err := channel.Consume(
		requestQueue.Name, // queue
		"",                // consumer
		true,              // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)

	if err != nil {
		panic(err)
	}

	for msg := range messages {
		var request Request
		err := json.Unmarshal(msg.Body, &request)
		if err != nil {
			continue
		}

		dirFiles := parseRequest(request, directory, channel)
		body := &bytes.Buffer{}
		enc := json.NewEncoder(body)
		enc.Encode(dirFiles)
		err = publish(request, channel, body.Bytes())
	}
}
