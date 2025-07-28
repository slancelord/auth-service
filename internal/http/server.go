package http

import (
	"log"
	"net/http"

	"auth-service/internal/config"
	"auth-service/internal/http/router"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Type "Bearer <token>" for authorization.

func Serve() {
	mux := http.NewServeMux()
	port := config.GetConfig().AppPort

	router.InitAuthRoutes(mux)
	router.InitSwaggerRoutes(mux)

	log.Printf("[INFO] Server started at http://localhost:%s\n", port)

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatalf("[ERROR] Failed to start server: %v", err)
	}
}
