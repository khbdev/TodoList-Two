package client

import (
	"auth-service/internal/connect"
	"context"
	"fmt"
	"time"

	loginpb "github.com/khbdev/todolist-proto/proto/login"
	"google.golang.org/grpc"
)

type LoginClient struct {
	conn   *grpc.ClientConn
	client loginpb.UserServiceClient
}

func NewLoginClient() (*LoginClient, error) {
	conn, err := connect.ConnectService("user-service")
	if err != nil {
		return nil, fmt.Errorf("user-service login proto ga ulanish xatosi: %v", err)
	}

	return &LoginClient{
		conn:   conn,
		client: loginpb.NewUserServiceClient(conn),
	}, nil
}

func (l *LoginClient) Login(ctx context.Context, req *loginpb.LoginRequest) (*loginpb.LoginResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	res, err := l.client.Login(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("Login RPC xato: %v", err)
	}

	return res, nil
}

func (l *LoginClient) Close() {
	if l.conn != nil {
		l.conn.Close()
	}
}