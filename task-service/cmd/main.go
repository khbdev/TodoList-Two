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

	// 2) kafka topic create (1 marta)
	if err := config.CreateTopic(); err != nil {
		log.Fatal(err)
	}

	// 3) kafka producer (yopib qo‘yma! defer bilan yopiladi)
	prod := event.NewProducer()
	defer func() {
		if err := prod.Close(); err != nil {
			log.Println("producer close error:", err)
		}
	}()

	// 4) db
	db, err := config.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	// 5) redis + cache
	rdb, err := config.NewRedisClient()
	if err != nil {
		log.Fatal(err)
	}
	c := cache.NewReminderCache(rdb)

	// 6) repo + usecase + handler
	repo := postgres.NewTaskRepo(db)
	srv := usecase.NewReminderService(repo, c, 5*time.Minute)
	hand := handler.NewTaskHandler(srv)

	// 7) reminder worker (background)
	ctx := context.Background()
	w := worker.NewReminderWorker(db, prod)
	go w.Run(ctx)

	// 8) grpc server
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