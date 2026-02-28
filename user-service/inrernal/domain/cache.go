package domain

import (
	"context"
	"time"
	"user-service/inrernal/repository/model"
)

type UserCache interface {
	GetByID(ctx context.Context, id int) (*model.User, error)
	SetByID(ctx context.Context, user *model.User, ttl time.Duration) error
	DeleteByID(ctx context.Context, id int) error
}
