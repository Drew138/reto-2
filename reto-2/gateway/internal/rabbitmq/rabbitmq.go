package rabbitmq

import (
	"context"
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (s *ServiceAMQP) Publish(body []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := s.Ch.PublishWithContext(ctx,
		"",                  // exchange
		s.RequestQueue.Name, // routing key
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			ContentType: "encoding/json",
			Body:        body,
		})
	log.Printf(" [x] Sent %s\n", body)
	return err
}

type ServiceAMQP struct {
	C            chan []string
	RequestQueue amqp.Queue
	Ch           *amqp.Channel
	Conn         *amqp.Connection
}

func Setup(rabbitmqURL string) (*ServiceAMQP, error) {

	amqpConn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		panic(err)
	}
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
		return nil, err
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
		return nil, err
	}

	messages, err := ch.Consume(
		responseQueue.Name, // queue
		"",                 // consumer
		true,               // auto-ack
		false,              // exclusive
		false,              // no-local
		false,              // no-wait
		nil,                // args
	)
	if err != nil {
		return nil, err
	}
	forever := make(chan []string, 100)
	go func() {
		for msg := range messages {
			var response []string
			json.Unmarshal(msg.Body, &response)
			forever <- response
		}
	}()
	return &ServiceAMQP{C: forever, RequestQueue: requestQueue, Ch: ch, Conn: amqpConn}, err
}
