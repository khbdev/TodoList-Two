package domain

import (
	"context"
	"user-service/internal/repositroy/model"
)


type UserRepository interface {
	Create( ctx context.Context, user model.User) (error)
	GetAll(ctx context.Context) ([]model.User, error)
	GetByID(ctx context.Context, user_id int) (model.User, error)
	Update(ctx context.Context, user model.User) (model.User, error)
	Delete(ctx context.Context, user_id int) (bool, error)
}