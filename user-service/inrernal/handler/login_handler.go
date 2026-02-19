package handler

import (
	"context"

	"user-service/inrernal/domain"

	loginpb "github.com/khbdev/todolist-proto/proto/login"
)

type LoginGRPCHandler struct {
	loginpb.UnimplementedUserServiceServer
	uc domain.UserUsecase
}

func NewLoginGRPCHandler(uc domain.UserUsecase) *LoginGRPCHandler {
	return &LoginGRPCHandler{uc: uc}
}


func (h *LoginGRPCHandler) Login(ctx context.Context, req *loginpb.LoginRequest) (*loginpb.LoginResponse, error) {
	user, ok, err := h.uc.Login(ctx, req.Email, req.Password)
	if err != nil {

		return nil, err
	}

	if user == nil {
		return &loginpb.LoginResponse{Success: false}, nil
	}

	return &loginpb.LoginResponse{
		Success: ok,
		User: &loginpb.User{
			UserId: int64(user.ID),
			Name:   user.Name,
			Email:  user.Email,
		},
	}, nil
}
