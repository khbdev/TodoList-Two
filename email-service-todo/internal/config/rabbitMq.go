package config

import (
	"log"
	"os"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Conn    *amqp091.Connection
	Channel *amqp091.Channel
}

func NewRabbitMQ() *RabbitMQ {
	conn, err := amqp091.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatal("rabbitmq connection error:", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("rabbitmq channel error:", err)
	}

	return &RabbitMQ{
		Conn:    conn,
		Channel: ch,
	}
}