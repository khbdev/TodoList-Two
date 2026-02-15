package main

import (
	"log"
	"user-service/inrernal/config"
	"user-service/pkg"
)


func main(){
	// env 
	pkg.LoadEnv()

	// postgres
	db, err := config.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	_ = db

}