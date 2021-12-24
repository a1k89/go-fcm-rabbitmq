package main

import (
	"github.com/joho/godotenv"
	"log"
	"notify/consumer"
)

func main()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Please provide .env file")
	}

	consumer.StartConsumeFromMainChannel()
}