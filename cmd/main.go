package main

import (
	"auth-service/internal/config"
	"auth-service/internal/database"
	"auth-service/internal/http"
)

func main() {
	config.Init()

	database.Init()

	http.Serve()
}
