package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type request struct {
	Name  string `json:"name"`
	Query string `json:"query"`
}

func listFiles(directory string) []string {
	var files []string

	dirEntries, _ := os.ReadDir(directory)
	for _, file := range dirEntries {
		files = append(files, file.Name())
	}
	return files
}

func searchFiles(directory, name string) []string {
	files := listFiles(directory)
	var ret []string
	if name == "*" {
		return files
	}
	for _, file := range files {
		if strings.Contains(file, name) {
			ret = append(ret, file)
		}
	}
	return ret
}

func requestHandler(body []byte, directory string) []string {
	var req request
	err := json.Unmarshal(body, &req)
	if err != nil {
		return []string{err.Error()}
	}
	if req.Name == "list" {
		return listFiles(directory)
	} else if req.Name == "search" {
		return searchFiles(directory, req.Query)
	} else {
		return []string{"Invalid request"}
	}
}

func publish(responseQueue amqp.Queue, ch *amqp.Channel, body []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := ch.PublishWithContext(ctx,
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
	rabbitmqURL := os.Getenv("RABBITMQ_URL")
	amqpConn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		panic(err)
	}
	defer amqpConn.Close()
	ch, err := amqpConn.Channel()
	requestQueue, err := ch.QueueDeclare(
		"request", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		panic(err)
	}
	responseQueue, err := ch.QueueDeclare(
		"response", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		panic(err)
	}

	messages, err := ch.Consume(
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
	directory := os.Getenv("MICROSERVICE_TWO_DIRECTORY")
	for msg := range messages {
		l := requestHandler(msg.Body, directory)
		body := &bytes.Buffer{}
		enc := json.NewEncoder(body)
		enc.Encode(l)
		err = publish(responseQueue, ch, body.Bytes())
	}

}
