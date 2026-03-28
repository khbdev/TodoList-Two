package domain

import "context"

type BrokerType string





type SendEmailUsecase interface {
	SendToRabbitMQ(ctx context.Context, name, email string) error
	SendToKafka(ctx context.Context, userID int64, taskName string) error
}