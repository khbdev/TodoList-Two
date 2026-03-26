package domain

import "context"

type BrokerType string





type SendEmailUsecase interface {
	Execute(ctx context.Context, brokerType string, data any) error
}
