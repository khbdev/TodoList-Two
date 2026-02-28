package main

import (
	"log"
	"user-service/internal/config"
	"user-service/pkg"
)




func main(){
	pkg.LoadEnv()

	db, err := config.PostgressConnection()
	if err != nil {
		log.Fatal(err)
	}

	_ = db
}