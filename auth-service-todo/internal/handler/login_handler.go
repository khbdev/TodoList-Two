package handler

import (
	"auth-service/internal/usecase"
	"context"

	authpb "github.com/khbdev/todolist-proto/proto/auth"
)

type LoginHandler struct {
	authpb.UnimplementedAuthServiceServer
	loginUsecase *usecase.LoginUsecase
}

func NewLoginHandler(loginUsecase *usecase.LoginUsecase) *LoginHandler {
	return &LoginHandler{
		loginUsecase: loginUsecase,
	}
}

func (h *LoginHandler) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	_, token, err := h.loginUsecase.Login(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, err
	}

	return &authpb.LoginResponse{
		AccessToken: token,
	}, nil
}