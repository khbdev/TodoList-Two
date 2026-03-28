package config

import (
	"fmt"
	"log"

	"os"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

func env(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}

func CreateTopic() error {
	bootstrap := env("KAFKA_BOOTSTRAP", "localhost:9092")
	topic := env("KAFKA_TOPIC", "task")

	log.Println("Kafka: connecting to", bootstrap, "topic:", topic)

	dialer := &kafka.Dialer{Timeout: 10 * time.Second, DualStack: true}

	conn, err := dialer.Dial("tcp", bootstrap)
	if err != nil {
		return fmt.Errorf("bootstrap dial error: %w", err)
	}
	defer conn.Close()


	err = conn.CreateTopics(kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     3,
		ReplicationFactor: 1,
	})
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "already exists") {
			log.Println("Kafka: topic already exists:", topic)
			return nil
		}
		return fmt.Errorf("create topics error: %w", err)
	}

	log.Println("Kafka: topic created:", topic)
	return nil
}