package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Service interface {
	Connect() error
	Pusblish(msg string) error
	Consume()
}

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewRabbitMQService() *RabbitMQ {
	return &RabbitMQ{}
}

func (r *RabbitMQ) Connect() error {
	fmt.Println("Connecting to RabbitMQ")

	var err error

	r.Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}

	fmt.Println("Successfully connected")

	r.Channel, err = r.Conn.Channel()
	if err != nil {
		return err
	}

	_, err = r.Channel.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *RabbitMQ) Pusblish(msg string) error {
	err := r.Channel.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)

	if err != nil {
		return err
	}

	//fmt.Println("Publiching message")
	return nil
}

func (r *RabbitMQ) Consume() error {
	msgs, err := r.Channel.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	for msg := range msgs {
		fmt.Printf("Received Message: %s \n", msg.Body)
	}

	return nil
}
