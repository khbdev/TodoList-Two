package main

import "emai-service/internal/config"


func main(){

	rabbit := config.NewRabbitMQ()
	defer rabbit.Conn.Close()
	defer rabbit.Channel.Close()

	kafka := config.NewKafka()

	defer kafka.Close()

}