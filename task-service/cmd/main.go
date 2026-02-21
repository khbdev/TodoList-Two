package main

import (

	"task-service/internal/config"
	"task-service/pkg"
)






func main(){
	pkg.LoadEnv()

	// db, err := config.NewPostgresDB()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	config.NewPostgresDB()
	config.NewRedisClient()
	config.CreateTopic()

}
