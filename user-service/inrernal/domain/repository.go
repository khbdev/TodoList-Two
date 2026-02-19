package domain

import (
	"context"
	"user-service/inrernal/repository/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	
	GetByID(ctx context.Context, id uint) (*model.User, error)
	GetAll(ctx context.Context) ([]model.User, error)

	Update(ctx context.Context, user *model.User) error
	
	Delete(ctx context.Context, id uint) error

}



type AuthUsecase interface {
	Login(ctx context.Context, email, password string) (*User, bool, error)
}
