package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"emai-service/internal/client"
	"emai-service/internal/config"
	"emai-service/internal/consumer"
	"emai-service/internal/service"
	"emai-service/internal/usecase"
	"emai-service/pkg/env"
)

func main() {
	// .env fayldan konfiguratsiya o'qish
	env.Load()

	// RabbitMQ konfiguratsiyasi
	rabbit := config.NewRabbitMQ()
	defer rabbit.Conn.Close()
	defer rabbit.Channel.Close()

	// Kafka konfiguratsiyasi
	kafka := config.NewKafka()
	defer kafka.Close()

	// User service client
	clientUser := client.NewUserClient()

	// Email service
	email := service.NewEmailService()

	// Usecase
	srv := usecase.NewEmailUsecase(email, clientUser)

	// Consumerlar
	consumerKafka := consumer.NewKafkaConsumer(kafka.Reader, srv)
	consumerRabbitMq := consumer.NewRabbitMQConsumer(rabbit, srv)

	
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()


	go func() {
		if err := consumerKafka.Consume(ctx); err != nil {
			log.Printf("Kafka consumer error: %v", err)
		}
	}()
	go func() {
		if err := consumerRabbitMq.Consume(ctx); err != nil {
			log.Printf("RabbitMQ consumer error: %v", err)
		}
	}()


	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	sig := <-sigCh
	log.Printf("Received signal %v, shutting down gracefully...", sig)


	cancel()
	log.Println("Service stopped")
}