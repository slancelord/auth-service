package main

import (
	"auth-service/internal/config"
	"auth-service/internal/database"
	"auth-service/internal/server"
)

func main() {
	config.Init()

	database.Init()
	defer database.GetDB().Close()

	server.Serve()
}
