package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)



var (
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
)

func PostgressConnection() *sql.DB {
	connstr := fmt.Sprintf("host=%s port=% user=%s password=%s dbname=%s sslmode=disable")

	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("PostgresSql connection SuccessFull")
}