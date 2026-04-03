// internal/usecase/email_usecase.go
package usecase

import (
	"context"
	"emai-service/internal/client"
	"emai-service/internal/domain"
	"fmt"
	"time"
)

type emailUsecase struct {
	userClient  *client.UserClient
	emailSender domain.EmailSender
}


func NewEmailUsecase(sender domain.EmailSender, uc *client.UserClient) domain.SendEmailUsecase {
	return &emailUsecase{
		userClient:  uc,
		emailSender: sender,
	}
}


func (u *emailUsecase) SendToRabbitMQ(ctx context.Context, name, email string) error {
	data := domain.EmailData{
		Name:    name,
		Email:   email,
		Subject: "Ro'yhatdan o'tish",
		Body:    "Siz Todolist dan ro'yhatdan o'tdingiz",
	}

	return u.emailSender.Send(ctx, data)
}

// SendToKafka: userID orqali user ma’lumotini olib, taskName bilan yuboradi
func (u *emailUsecase) SendToKafka(ctx context.Context, userID int64, taskName string) error {
	user, err := u.userClient.GetUser(userID)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	data := domain.EmailData{
		Name:    user.Name,
		Email:   user.Email,
		Subject: fmt.Sprintf("Task vaqti: %s", time.Now().Format("2006-01-02 15:04")),
		Body:    taskName,
	}

	if err := u.emailSender.Send(ctx, data); err != nil {
		return fmt.Errorf("failed to send to kafka: %w", err)
	}

	return nil
}