package client

import (
	"auth-service/internal/connect"
	"context"
	"fmt"
	"time"

	userpb "github.com/khbdev/todolist-proto/proto/user"
	"google.golang.org/grpc"
)

type UserClient struct {
	conn   *grpc.ClientConn
	client userpb.UserServiceClient
}

func NewUserClient() (*UserClient, error) {
	conn, err := connect.ConnectService("user-service")
	if err != nil {
		return nil, fmt.Errorf("user-service ga ulanish xatosi: %v", err)
	}

	return &UserClient{
		conn:   conn,
		client: userpb.NewUserServiceClient(conn),
	}, nil
}

func (u *UserClient) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	res, err := u.client.CreateUser(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("CreateUser RPC xato: %v", err)
	}

	return res, nil
}

func (u *UserClient) Close() {
	if u.conn != nil {
		u.conn.Close()
	}
}