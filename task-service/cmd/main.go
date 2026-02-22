package main

import (
	"log"
	"os"
	"task-service/internal/cache"
	"task-service/internal/config"
	"task-service/internal/handler"
	"task-service/internal/repository/postgres"
	"task-service/internal/usecase"
	"task-service/pkg"
	"time"

	taskpb "github.com/khbdev/todolist-proto/proto/task"
	"google.golang.org/grpc"
)






func main(){
	pkg.LoadEnv()

		// config.CreateTopic()
	db, err := config.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}
	rdb, err := config.NewRedisClient()
	if err != nil {
		log.Fatal(err)
	}
	 
	cache := cache.NewReminderCache(rdb)
	

     repo := postgres.NewTaskRepo(db)
	 srv := usecase.NewReminderService(repo, cache, 5*time.Minute)
	 hand := handler.NewTaskHandler(srv)

	 grpcServer := grpc.NewServer()

	 taskpb.RegisterTaskServiceServer(grpcServer, hand)

	 	addr := os.Getenv("PORT")
	if addr == "" {
		addr = ":50053"
	}



	
}
