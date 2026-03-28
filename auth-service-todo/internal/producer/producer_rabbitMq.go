package producer

import (
	"auth-service/internal/config"
	"encoding/json"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	conn    *config.RabbitMQConnection
	exchange string
	routingKey string
}


func NewProducer(conn *config.RabbitMQConnection) *Producer {
	return &Producer{
		conn:       conn,
		exchange:   "email_auth_exchange",
		routingKey: "email_key",
	}
}


func (p *Producer) Publish(message interface{}) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = p.conn.Channel.Publish(
		p.exchange,   
		p.routingKey, 
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		log.Println("Publish error:", err)
		return err
	}

	log.Println("Message published successfully")
	return nil
}