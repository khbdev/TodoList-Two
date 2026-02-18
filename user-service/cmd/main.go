package main

import (
	"log"
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
 
   srv := usecase.NewUserService(repo)

   hand := handler.NewUserGRPCHandler(srv)


  cache := cache.NewUserCache(redisClient)

  _ = cache
}
