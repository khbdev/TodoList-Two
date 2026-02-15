package domain

import "context"

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	
	GetByID(ctx context.Context, id int) (*User, error)
	GetAll(ctx context.Context) ([]User, error)

	Update(ctx context.Context, user *User) error
	
	Delete(ctx context.Context, id int) error
}
