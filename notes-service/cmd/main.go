package main

import (
	// "log"
	"log"
	"notes-service/internal/cache"
	"notes-service/internal/config"
	"notes-service/internal/repository/postgres"
	"notes-service/internal/usecase"
	"notes-service/pkg"
	"time"
)



func main(){
	// load env
	pkg.LoadEnv()

	// sql connection

	db, err := config.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

   




	rdb, err := config.NewRedisClient()
if err != nil {
		log.Fatal(err)
	}


	
   cache := cache.NewTodoCache(rdb)
	repo := postgres.NewTodoRepo(db)
	srv := usecase.NewTodoService(repo, cache, 5*time.Minute)

 

}