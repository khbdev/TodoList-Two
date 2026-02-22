package main

import (
	"log"
	"task-service/internal/config"
	"task-service/internal/repository/postgres"
	"task-service/pkg"
)






func main(){
	pkg.LoadEnv()

	db, err := config.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}
	// config.NewPostgresDB()
	config.NewRedisClient()
	config.CreateTopic()
     repo := postgres.NewTaskRepo(db)

	 _ = repo
}
