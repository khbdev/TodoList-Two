package config

import (
	"fmt"
	"log"
	"os"

	"github.com/rabbitmq/amqp091-go"
)




type RabbitMQConnection struct {
	Conn  *amqp091.Connection
	Channel *amqp091.Channel

}



func NewRabbitMq() *RabbitMQConnection {
	url := os.Getenv("AMQP_URL")

	conn, err := amqp091.Dial(url)
	if err != nil {
		log.Fatal("Failed to Connection")
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to Channel")
	}
	defer ch.Close()

	r := &RabbitMQConnection{
		Conn: conn,
		Channel: ch,
	}
	
fmt.Println("RabbitMQ Connection SuccessFull")
if err := r.RabbitMqSetup(); err != nil {
		log.Fatal("RabbitMQ Setup nil")
	}
	return  r
	
}



func (r *RabbitMQConnection ) RabbitMqSetup() error {
	queue_name := "email_auth_queue"
	exchange_name := "email_auth_exchange"
	rauting_key := "email_key"

	err := r.Channel.ExchangeDeclare(
			exchange_name,
		"direct",     
		true, false, false, false, nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	_, err = r.Channel.QueueDeclare(
		queue_name,
		true, false, false, false, nil,
	)

	if err != nil {
		log.Fatal(err)
	}
	err = r.Channel.QueueBind(
		queue_name,
		rauting_key,
		exchange_name,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("RabbiMQ Setup SuccessFull ")
	return err
}