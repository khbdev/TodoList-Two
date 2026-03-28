package usecase

import (
	"auth-service/internal/client"
	"auth-service/internal/model"
	"auth-service/internal/producer"
	"context"
	"fmt"
	"log"

	userpb "github.com/khbdev/todolist-proto/proto/user"
)

type AuthUsecase struct {
	userClient *client.UserClient
	producer producer.Producer
}


func NewAuthUsecase(userClient *client.UserClient, producer producer.Producer) *AuthUsecase {
	return &AuthUsecase{
		userClient: userClient,
		producer: producer,
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
     request := model.Req{
      Name: res.User.Name,
	  Email: res.User.Email,
	 }

	if err := u.producer.Publish(request); err != nil {
		log.Fatal("Error: ", err)
	}
	return res, nil
}