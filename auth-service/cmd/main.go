package main

import (
	"auth-service/internal/config"
	"auth-service/pkg/env"
)



func main(){
	env.LoadEnv()
	 
	rabbitMq := config.NewRabbitMq()

	_ = rabbitMq
	
	
}