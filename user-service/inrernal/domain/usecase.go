package domain

import (
	"context"
	"user-service/inrernal/repository/model"
)

type UserUsecase interface {
	Create(ctx context.Context, name string, email string, password string) (*model.User, error)
	Update(ctx context.Context, id int, name string, email string) (*model.User, error)
	Delete(ctx context.Context, id int) error

	GetByID(ctx context.Context, id int) (*model.User, error)
	GetAll(ctx context.Context) ([]model.User, error)

	Login(ctx context.Context, email string, password string) (*model.User, bool, error)
}
