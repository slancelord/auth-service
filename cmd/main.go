package main

import (
	"log"

	"auth-service/internal/server"

	"github.com/joho/godotenv"
)

func init() {
	// load env file
	if err := godotenv.Load(); err != nil {
		log.Println("[WARN] .env not found")
	}
}

func main() {
	server.Serve()
}
