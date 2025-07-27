package main

import (
	"log"
	"task_management/data"
	"task_management/router"

	"github.com/joho/godotenv"
)

func main() {
	// Load env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	data.InitDB()
	router.Routes()

}
