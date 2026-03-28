package client

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/khbdev/todolist-proto/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient struct {
	client userpb.UserServiceClient
}


func NewUserClient() *UserClient {
	url := os.Getenv("USER_SERVICE_URL")
	if url == "" {
		url = "localhost:50051"
	}

	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(fmt.Sprintf("failed to connect: %v", err))
	}

	client := userpb.NewUserServiceClient(conn)
	log.Println("User service connection established")
	return &UserClient{client: client}
}


func (c *UserClient) GetUser(id int64) (*userpb.User, error) {
	var (
		res *userpb.GetUserResponse
		err error
	)

	maxRetries := 3
	for attempt := 1; attempt <= maxRetries; attempt++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		req := &userpb.GetUserRequest{Id: id}

		res, err = c.client.GetUser(ctx, req)
		cancel() 

		if err == nil {
			return res.User, nil
		}

		log.Printf("Attempt %d/%d failed: %v", attempt, maxRetries, err)
		time.Sleep(time.Second * time.Duration(attempt))
	}

	return nil, fmt.Errorf("all %d attempts failed: last error: %w", maxRetries, err)
}