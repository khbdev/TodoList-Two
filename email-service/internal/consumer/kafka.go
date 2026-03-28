package consumer

import (
	"context"
	"emai-service/internal/domain"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	reader  *kafka.Reader
	usecase domain.SendEmailUsecase
}

type KafkaTaskMessage struct {
	UserId int64 `json:"user_id"`
	TaskName string `json:"task_name"`
}

func NewKafkaConsumer(reader *kafka.Reader, usecase domain.SendEmailUsecase) *KafkaConsumer {
	return &KafkaConsumer{
		reader:  reader,
		usecase: usecase,
	}
}

func (c *KafkaConsumer) Consume(ctx context.Context) error {
	log.Println("kafka consumer started...")

	for {
		select {
		case <-ctx.Done():
			log.Println("kafka consumer stopped")
			return nil
		default:
			msg, err := c.reader.FetchMessage(ctx)
			if err != nil {
				log.Println("kafka fetch message error:", err)
				continue
			}

			var data KafkaTaskMessage
			if err := json.Unmarshal(msg.Value, &data); err != nil {
				log.Println("kafka json unmarshal error:", err)
				continue
			}

			if err := c.usecase.SendToKafka(ctx, data.UserId, data.TaskName); err != nil {
				log.Println("kafka usecase error:", err)
				continue
			}

			if err := c.reader.CommitMessages(ctx, msg); err != nil {
				log.Println("kafka commit error:", err)
				continue
			}
		}
	}
}