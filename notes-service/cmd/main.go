package main

import (
	// "log"
	"notes-service/internal/config"
	"notes-service/pkg"
)



func main(){
	// load env
	pkg.LoadEnv()

	// sql connection

	// db, err := 
	// if err != nil {
	// 	log.Fatal(err)
	// }

	config.NewPostgresDB()
	config.NewRedisClient()


 

}