package main

import (
	"auth-service/internal/client"
	"auth-service/internal/config"
	"auth-service/pkg/env"
	"log"
)



func main(){
	env.LoadEnv()
	 
	rabbitMq := config.NewRabbitMq()

	_ = rabbitMq

	userServiceClient, err := client.NewUserClient()
	if err != nil {
		log.Fatal("user client error: ", err)
	}
	
	loginProtoClient, err := client.NewLoginClient()
		if err != nil {
		log.Fatal("user client error: ", err)
	}
	

	_ = userServiceClient
	_ = loginProtoClient


	select {}
	
}