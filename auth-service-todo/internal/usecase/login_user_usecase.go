package usecase

import (
	"auth-service/internal/client"
	"auth-service/pkg/jwt"
	"context"
	"fmt"

	loginpb "github.com/khbdev/todolist-proto/proto/login"
)

type LoginUsecase struct {
	loginClient *client.LoginClient
}

func NewLoginUsecase(loginClient *client.LoginClient) *LoginUsecase {
	return &LoginUsecase{
		loginClient: loginClient,
	}
}

func (l LoginUsecase) Login(ctx context.Context, email, password string) (*loginpb.LoginResponse, string, error) {
	login := &loginpb.LoginRequest{
		Email:    email,
		Password: password,
	}

	res, err := l.loginClient.Login(ctx, login)
	if err != nil {
		return nil, "", fmt.Errorf("usecase login error: %v", err)
	}
token, err := jwt.GenerateAccesJwtAcccToken(uint(res.User.UserId))
if err != nil {
	return nil, "", fmt.Errorf("token generation error: %v", err)
}
	return res, token, nil
}