package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	taskpb "github.com/khbdev/todolist-proto/proto/task"
	"google.golang.org/grpc"

	"task-service/internal/cache"
	"task-service/internal/config"
	"task-service/internal/event"
	"task-service/internal/handler"
	"task-service/internal/repository/postgres"
	"task-service/internal/usecase"
	"task-service/internal/worker"
	"task-service/pkg"
)

func main() {
	// 1) env
	pkg.LoadEnv()

	
	if err := config.CreateTopic(); err != nil {
		log.Fatal(err)
	}

	
	prod := event.NewProducer()
	defer func() {
		if err := prod.Close(); err != nil {
			log.Println("producer close error:", err)
		}
	}()


	db, err := config.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	
	rdb, err := config.NewRedisClient()
	if err != nil {
		log.Fatal(err)
	}
	c := cache.NewReminderCache(rdb)

	repo := postgres.NewTaskRepo(db)
	srv := usecase.NewReminderService(repo, c, 5*time.Minute)
	hand := handler.NewTaskHandler(srv)

	
	ctx := context.Background()
	w := worker.NewReminderWorker(db, prod)
	go w.Run(ctx)


	grpcServer := grpc.NewServer()
	taskpb.RegisterTaskServiceServer(grpcServer, hand)

	addr := os.Getenv("PORT")
	if addr == "" {
		addr = ":50053"
	}

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Task gRPC server running on %s", addr)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}