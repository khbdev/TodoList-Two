package usecase

import (
	"auth-service/internal/client"
	"context"
	"fmt"

	userpb "github.com/khbdev/todolist-proto/proto/user"
)

type AuthUsecase struct {
	userClient *client.UserClient
}


func NewAuthUsecase(userClient *client.UserClient) *AuthUsecase {
	return &AuthUsecase{
		userClient: userClient,
	}
}


func (u *AuthUsecase) Register(ctx context.Context, name, email, password string) (*userpb.CreateUserResponse, error) {


	req := &userpb.CreateUserRequest{
		Name: name,
		Email:    email,
		Password: password,
	}


	res, err := u.userClient.CreateUser(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("usecase register error: %v", err)
	}

	return res, nil
}