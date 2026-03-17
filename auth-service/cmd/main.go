package main

import (
	"auth-service/internal/client"
	"auth-service/internal/config"
	"auth-service/internal/usecase"
	"auth-service/pkg/env"
	"context"
	"fmt"
	"log"
	"time"
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
	

	
	_ = loginProtoClient

	createUserUsecase := usecase.NewAuthUsecase(userServiceClient)
ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
defer cancel()
user, err :=	createUserUsecase.Register(ctx, "Azizbek", "khbcoderssssss@gmail.com", "salom123")
if err != nil {
	log.Fatal(err)
}
fmt.Println(user)

	select {}
	
}