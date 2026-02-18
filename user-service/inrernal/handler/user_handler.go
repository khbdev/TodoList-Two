package handler

import (
	"context"

	"user-service/inrernal/domain"
	"user-service/pkg"

	userpb "github.com/khbdev/todolist-proto/proto/user"
)

type UserGRPCHandler struct {
	userpb.UnimplementedUserServiceServer
	uc domain.UserUsecase
}

func NewUserGRPCHandler(uc domain.UserUsecase) *UserGRPCHandler {
	return &UserGRPCHandler{uc: uc}
}

// -------------------- CreateUser --------------------
func (h *UserGRPCHandler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	user, err := h.uc.Create(ctx, req.Name, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	resp := pkg.ToUserResponse(user)

	return &userpb.CreateUserResponse{
		User: &userpb.User{
			Id:        int64(resp.ID),
			Name:      resp.Name,
			Email:     resp.Email,
			CreatedAt: resp.CreatedAt.String(),
			UpdatedAt: resp.UpdatedAt.String(),
		},
	}, nil
}

// -------------------- GetUser --------------------
func (h *UserGRPCHandler) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	user, err := h.uc.GetByID(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	resp := pkg.ToUserResponse(user)

	return &userpb.GetUserResponse{
		User: &userpb.User{
			Id:        int64(resp.ID),
			Name:      resp.Name,
			Email:     resp.Email,
			CreatedAt: resp.CreatedAt.String(),
			UpdatedAt: resp.UpdatedAt.String(),
		},
	}, nil
}

// -------------------- UpdateUser --------------------
func (h *UserGRPCHandler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	// proto’da password bor, lekin sening usecase update password qabul qilmaydi -> ignore
	user, err := h.uc.Update(ctx, int(req.Id), req.Name, req.Email)
	if err != nil {
		return nil, err
	}

	resp := pkg.ToUserResponse(user)

	return &userpb.UpdateUserResponse{
		User: &userpb.User{
			Id:        int64(resp.ID),
			Name:      resp.Name,
			Email:     resp.Email,
			CreatedAt: resp.CreatedAt.String(),
			UpdatedAt: resp.UpdatedAt.String(),
		},
	}, nil
}

// -------------------- DeleteUser --------------------
func (h *UserGRPCHandler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	if err := h.uc.Delete(ctx, int(req.Id)); err != nil {
		return nil, err
	}
	return &userpb.DeleteUserResponse{Message: "deleted"}, nil
}

// -------------------- ListUsers --------------------
func (h *UserGRPCHandler) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	// limit/offset hozir usecase’da yo’q -> ignore
	users, err := h.uc.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	respList := pkg.ToUserResponseList(users)

	pbUsers := make([]*userpb.User, 0, len(respList))
	for _, r := range respList {
		pbUsers = append(pbUsers, &userpb.User{
			Id:        int64(r.ID),
			Name:      r.Name,
			Email:     r.Email,
			CreatedAt: r.CreatedAt.String(),
			UpdatedAt: r.UpdatedAt.String(),
		})
	}

	return &userpb.ListUsersResponse{
		Users: pbUsers,
	}, nil
}
