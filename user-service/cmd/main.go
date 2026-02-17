package main

import (
	"log"
	"user-service/inrernal/cache"
	"user-service/inrernal/config"
	"user-service/inrernal/repository/postgres"

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

	_ = db
	// redis
	redisClient, err := config.NewRedisClient()
	if err != nil {
		log.Fatal(err)
	}

	_ = redisClient


  repo :=  postgres.NewUserRepo(db)
 
  _ = repo


  cache := cache.NewUserCache(redisClient)

  _ = cache
}
