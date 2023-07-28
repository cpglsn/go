package main

import (
	"fmt"

	"github.com/cpglsn/rabbitmq/internal/rabbitmq"
)

type App struct {
	Rmq *rabbitmq.RabbitMQ
}

func Run() error {
	fmt.Println("Go with RabbitMQ")

	rmq := rabbitmq.NewRabbitMQService()

	app := App{
		Rmq: rmq,
	}

	if err := app.Rmq.Connect(); err != nil {
		return err
	}
	defer func() {
		app.Rmq.Conn.Close()
		fmt.Println("Closing connection")
	}()

	err := app.Rmq.Pusblish("Hi")
	if err != nil {
		return err
	}

	err = app.Rmq.Consume()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Printf("Error setting up application: %v \n", err)
	}
}
