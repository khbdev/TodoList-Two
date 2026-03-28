package main

import (
	"auth-service/internal/client"
	"auth-service/internal/config"
	"auth-service/internal/handler"
	"auth-service/internal/producer"
	"auth-service/internal/usecase"
	"auth-service/pkg/env"
	"log"
	"net"
	"os"

	authpb "github.com/khbdev/todolist-proto/proto/auth"
	createpb "github.com/khbdev/todolist-proto/proto/create"
	"google.golang.org/grpc"
)

func main() {
	env.LoadEnv()

	port := os.Getenv("PORT")
	if port == "" {
		port = "50051"
	}

	rabbitMq := config.NewRabbitMq()

	userServiceClient, err := client.NewUserClient()
	if err != nil {
		log.Fatal("user client error: ", err)
	}

	loginProtoClient, err := client.NewLoginClient()
	if err != nil {
		log.Fatal("login client error: ", err)
	}

	producer := producer.NewProducer(rabbitMq)

	createUserUsecase := usecase.NewAuthUsecase(userServiceClient, *producer)
	loginUserUsecase := usecase.NewLoginUsecase(loginProtoClient)

	createHandler := handler.NewAuthHandler(createUserUsecase)
	loginHandler := handler.NewLoginHandler(loginUserUsecase)

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("listen error: ", err)
	}

	grpcServer := grpc.NewServer()

	createpb.RegisterUserServiceServer(grpcServer, createHandler)
	authpb.RegisterAuthServiceServer(grpcServer, loginHandler)

	log.Println("gRPC server running on port:", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("grpc serve error: ", err)
	}
}