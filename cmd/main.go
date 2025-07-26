package main

import (
	"auth-service/internal/config"
	"auth-service/internal/server"
)

func main() {
	config.InitConfig()
	server.Serve()
}
