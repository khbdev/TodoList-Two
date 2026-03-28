package domain

import "context"


type EmailSender interface {
	Send(ctx context.Context, data EmailData) error
}