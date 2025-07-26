package server

import (
	"log"
	"net/http"

	"auth-service/internal/config"
	"auth-service/internal/server/routes"
)

func Serve() {
	mux := http.NewServeMux()
	port := config.GetConfig().AppPort

	routes.InitAuthRoutes(mux)

	log.Printf("[INFO] Server started at http://localhost:%s\n", port)

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatalf("[ERROR] Failed to start server: %v", err)
	}
}
