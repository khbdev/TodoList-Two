package main

import (
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"

	"notes-service/internal/cache"
	"notes-service/internal/config"
	"notes-service/internal/handler"
	"notes-service/internal/repository/postgres"
	"notes-service/internal/usecase"
	"notes-service/pkg"

	notespb "github.com/khbdev/todolist-proto/proto/notes"
)

func main() {

	//  Load .env
	pkg.LoadEnv()

	//  Postgres
	db, err := config.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	//  Redis
	redisClient, err := config.NewRedisClient()
	if err != nil {
		log.Fatal(err)
	}

	// Repository
	repo := postgres.NewTodoRepo(db)

	// Cache
	todoCache := cache.NewTodoCache(redisClient)

	// Usecase (service)
	srv := usecase.NewTodoService(repo, todoCache, 5*time.Minute)

	//  gRPC handler
	grpcHandler := handler.NewNotesHandler(srv)

	//  gRPC server
	grpcServer := grpc.NewServer()

	// Register proto service
	notespb.RegisterNotesServiceServer(grpcServer, grpcHandler)

	//  Listen
	addr := os.Getenv("PORT")
	if addr == "" {
		addr = ":50052"
	}

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Notes gRPC server running on %s", addr)

	//  Serve
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}