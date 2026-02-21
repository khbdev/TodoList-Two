package config

import (
	"fmt"
	"net"
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
	bootstrap := env("KAFKA_BOOTSTRAP", "")
	topic := env("KAFKA_TOPIC", "")

	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}


	conn, err := dialer.Dial("tcp", bootstrap)
	if err != nil {
		return fmt.Errorf("bootstrap dial error: %w", err)
	}
	defer conn.Close()

	
	controller, err := conn.Controller()
	if err != nil {
		return fmt.Errorf("controller error: %w", err)
	}

	ctrlAddr := net.JoinHostPort(controller.Host, fmt.Sprintf("%d", controller.Port))


	adminConn, err := dialer.Dial("tcp", ctrlAddr)
	if err != nil {
		return fmt.Errorf("controller dial error: %w", err)
	}
	defer adminConn.Close()


	err = adminConn.CreateTopics(kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     3,
		ReplicationFactor: 1,
	})

	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "already exists") {
			return nil
		}
		return err
	}

	return nil
}