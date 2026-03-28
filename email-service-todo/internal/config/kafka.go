package config

import "github.com/segmentio/kafka-go"

type Kafka struct {
	Reader *kafka.Reader
}

func NewKafka() *Kafka {
	return &Kafka{
		Reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{"localhost:9092"},
			Topic:   "task_event",
			GroupID: "email-service",
		}),
	}
}

func (k *Kafka) Close() {
	k.Reader.Close()
}