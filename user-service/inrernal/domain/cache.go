package domain

import (
	"context"
	"time"
)

type UserCache interface {
	GetByID(ctx context.Context, id int) (*User, error)
	SetByID(ctx context.Context, user *User, ttl time.Duration) error
	DeleteByID(ctx context.Context, id int) error
}
