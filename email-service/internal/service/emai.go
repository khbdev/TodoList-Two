package service

import (
	"context"
	"emai-service/internal/domain"
	"fmt"
	"net/smtp"
	"os"
)

// EmailData – yuboriladigan email ma’lumotlari
type EmailData struct {
	Name    string
	Email   string
	Subject string
	Body    string
}

// EmailService – domain.EmailSender implementatsiyasi
type EmailService struct {
	smtpHost string
	smtpPort string
	username string
	password string
}

// NewEmailService – .env dan ma’lumotlarni oladi
func NewEmailService() domain.EmailSender {
	return &EmailService{
		smtpHost: os.Getenv("SMTP_HOST"),
		smtpPort: os.Getenv("SMTP_PORT"),
		username: os.Getenv("SMTP_USER"),
		password: os.Getenv("SMTP_PASS"),
	}
}

func (s *EmailService) Send(ctx context.Context, data any) error {
	var d EmailData
	switch v := data.(type) {
	case EmailData:
		d = v
	case *EmailData:
		d = *v
	default:
		return fmt.Errorf("invalid data type: %T", data)
	}

	auth := smtp.PlainAuth("", s.username, s.password, s.smtpHost)
	to := []string{d.Email}
	msg := []byte(
		"To: " + d.Email + "\r\n" +
			"Subject: " + d.Subject + "\r\n" +
			"\r\n" +
			d.Body + "\r\n",
	)

	addr := s.smtpHost + ":" + s.smtpPort
	err := smtp.SendMail(addr, auth, s.username, to, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	fmt.Printf("Email sent to %s successfully\n", d.Email)
	return nil
}