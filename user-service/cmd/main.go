package main

import (
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"

	"user-service/inrernal/cache"
	"user-service/inrernal/config"
	"user-service/inrernal/handler"
	"user-service/inrernal/usecase"

	"user-service/inrernal/repository/postgres"

	"user-service/pkg"

	userpb "github.com/khbdev/todolist-proto/proto/user"
	loginpb "github.com/khbdev/todolist-proto/proto/login"
)

func main() {

	pkg.LoadEnv()


	db, err := config.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}


	redisClient, err := config.NewRedisClient()
	if err != nil {
		log.Fatal(err)
	}

	
	
	repo := postgres.NewUserRepo(db)
	userCache := cache.NewUserCache(redisClient)
	srv := usecase.NewUserService(repo, userCache, 5*time.Minute)

	grpcHandler := handler.NewUserGRPCHandler(srv)
	loginHandler := handler.NewLoginGRPCHandler(srv)



	grpcServer := grpc.NewServer()


	userpb.RegisterUserServiceServer(grpcServer, grpcHandler)
	loginpb.RegisterUserServiceServer(grpcServer, loginHandler)


	addr := os.Getenv("PORT")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("gRPC server running on %s", addr)


	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
