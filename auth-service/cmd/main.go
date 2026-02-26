package main

import (
	"auth-service/internal/config"
	"auth-service/pkg/env"
	"auth-service/pkg/jwt"
	"fmt"
	"log"
)



func main(){
	env.LoadEnv()
	 
	rabbitMq := config.NewRabbitMq()

	_ = rabbitMq

	
	// jwt, err := jwt.GenerateAccesJwtAcccToken(1)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(jwt)
	
}