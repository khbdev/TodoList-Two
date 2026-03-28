package event

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func env(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}


func NewProducer() *Producer {
	bootstrap := env("KAFKA_BOOTSTRAP", "localhost:9092")
	topic := env("KAFKA_TOPIC", "task")

	writer := &kafka.Writer{
		Addr:         kafka.TCP(bootstrap),
		Topic:        topic,
		Balancer:     &kafka.Hash{},
		RequiredAcks: kafka.RequireOne,
		Async:        false,
	}

	return &Producer{
		writer: writer,
	}
}

type TaskEvent struct {
	UserID int64  `json:"user_id"`
	Task   string `json:"task"`
}

func (p *Producer) PublishTask(ctx context.Context, userID int64, task string) error {
	if userID <= 0 {
		return fmt.Errorf("invalid user_id")
	}
	if task == "" {
		return fmt.Errorf("task is empty")
	}

	payload, err := json.Marshal(TaskEvent{
		UserID: userID,
		Task:   task,
	})
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}

	msg := kafka.Message{
		Key:   []byte(strconv.FormatInt(userID, 10)),
		Value: payload,
		Time:  time.Now(),
	}

	if err := p.writer.WriteMessages(ctx, msg); err != nil {
		return fmt.Errorf("kafka write error: %w", err)
	}

	return nil
}

func (p *Producer) Close() error {
	return p.writer.Close()
}