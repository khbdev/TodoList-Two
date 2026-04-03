package config

import (
	"os"

	"github.com/segmentio/kafka-go"
)

type Kafka struct {
	Reader *kafka.Reader
}

func NewKafka() *Kafka {
	return &Kafka{
		Reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{os.Getenv("KAFKA_URL")},
			Topic:   os.Getenv("TOPIC"),
			GroupID: os.Getenv("GROUP_NAME"),
		}),
	}
}

func (k *Kafka) Close() {
	k.Reader.Close()
}