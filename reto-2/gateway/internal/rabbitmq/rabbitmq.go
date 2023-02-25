package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQService struct {
	Channel    *amqp.Channel
	Connection *amqp.Connection
}

type Response struct {
	Payload []string `json:"payload"`
	Key     string   `json:"key"`
}

func (service *RabbitMQService) Publish(body []byte, key string) error {
	err := service.Channel.Publish(
		"",        // exchange
		"request", // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "encoding/json",
			Body:        body,
		})
	if err != nil {
		log.Println("Could not publish message")
		return err
	}
	log.Printf(" [x] Sent %s\n", body)
	return err
}

func (service *RabbitMQService) Receive(name string) (<-chan amqp.Delivery, error) {
	args := make(amqp.Table)
	args["x-expires"] = int32(20000)
	responseQueue, err := service.Channel.QueueDeclare(
		name,  // name
		false, // durable
		true,  // delete when unused
		false, // exclusive
		false, // no-wait
		args,  // arguments
	)
	if err != nil {
		return nil, err
	}

	return service.Channel.Consume(
		responseQueue.Name, // queue
		"",                 // consumer
		true,               // auto-ack
		false,              // exclusive
		false,              // no-local
		false,              // no-wait
		nil,                // args
	)
}

func Setup(rabbitmqURL string) (*RabbitMQService, error) {

	amqpConn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		return nil, err
	}

	channel, err := amqpConn.Channel()
	return &RabbitMQService{Channel: channel, Connection: amqpConn}, err
}
