package main

import (
	"log"
	"time"
	"user-service/inrernal/cache"
	"user-service/inrernal/config"
	"user-service/inrernal/handler"
	"user-service/inrernal/repository/postgres"
	"user-service/inrernal/usecase"

	"user-service/pkg"
)

func main() {
	// env
	pkg.LoadEnv()

	// postgres
	db, err := config.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	
	// redis
	redisClient, err := config.NewRedisClient()
	if err != nil {
		log.Fatal(err)
	}

	


  repo :=  postgres.NewUserRepo(db)
  cache := cache.NewUserCache(redisClient)
   srv := usecase.NewUserService(repo, cache, 5*time.Minute)

   hand := handler.NewUserGRPCHandler(srv)

   

 


}
