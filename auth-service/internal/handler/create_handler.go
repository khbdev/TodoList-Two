package handler

import (
	"auth-service/internal/usecase"
	"context"

	createpb "github.com/khbdev/todolist-proto/proto/create"
)

type AuthHandler struct {
	createpb.UnimplementedUserServiceServer
	authUsecase *usecase.AuthUsecase
}

func NewAuthHandler(authUsecase *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase: authUsecase,
	}
}

func (h *AuthHandler) CreateUser(ctx context.Context, req *createpb.CreateUserRequest) (*createpb.CreateUserResponse, error) {
	_, err := h.authUsecase.Register(ctx, req.GetName(), req.GetEmail(), req.GetPassword())
	if err != nil {
		return &createpb.CreateUserResponse{
			Message: "xatolik yuz berdi",
		}, nil
	}

	return &createpb.CreateUserResponse{
		Message: "user muvaffaqiyatli yaratildi",
	}, nil
}