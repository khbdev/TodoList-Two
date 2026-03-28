package main

import (
	"emai-service/internal/client"
	"emai-service/internal/config"
	"emai-service/internal/service"
	"emai-service/internal/usecase"
	"emai-service/pkg/env"
)


func main(){
	env.Load()

	rabbit := config.NewRabbitMQ()
	defer rabbit.Conn.Close()
	defer rabbit.Channel.Close()

	kafka := config.NewKafka()

	defer kafka.Close()
 
	clientUser := client.NewUserClient()
	_ = clientUser

	email := service.NewEmailService()

	_ = email

	srv := usecase.NewEmailUsecase(email, clientUser)

	

}