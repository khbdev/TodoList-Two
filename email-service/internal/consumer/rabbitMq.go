package consumer

import (
	"context"
	"emai-service/internal/config"
	"emai-service/internal/domain"
	"encoding/json"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQConsumer struct {
	channel *amqp091.Channel
	usecase domain.SendEmailUsecase
}

type RabbitMQAuthMessage struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewRabbitMQConsumer(rabbit *config.RabbitMQ, usecase domain.SendEmailUsecase) *RabbitMQConsumer {
	return &RabbitMQConsumer{
		channel: rabbit.Channel,
		usecase: usecase,
	}
}

func (c *RabbitMQConsumer) Consume(ctx context.Context) error {
	queueName := "email_auth_queue"

	msgs, err := c.channel.Consume(
		queueName,
		"",
		false, 
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	log.Println("rabbitmq consumer started...")

	for {
		select {
		case <-ctx.Done():
			log.Println("rabbitmq consumer stopped")
			return nil

		case msg, ok := <-msgs:
			if !ok {
				log.Println("rabbitmq channel closed")
				return nil
			}

			var data RabbitMQAuthMessage
			if err := json.Unmarshal(msg.Body, &data); err != nil {
				log.Println("json unmarshal error:", err)
				_ = msg.Nack(false, false)
				continue
			}

			if err := c.usecase.Execute(ctx, "rabbitmq", data); err != nil {
				log.Println("usecase error:", err)
				_ = msg.Nack(false, true)
				continue
			}

			if err := msg.Ack(false); err != nil {
				log.Println("ack error:", err)
			}
		}
	}
}