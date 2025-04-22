// main.go
package main

import (
	"babybetgo/database"
	"babybetgo/handlers"
	"babybetgo/server"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitDB()
	handlers.DB = database.GetDB()
	server.ServerStart()

}
