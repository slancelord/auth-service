package servers

import (
	"log"
	"net/http"

	"auth-service/internal/config"
	"auth-service/internal/routers"
)

func Serve() {
	mux := http.NewServeMux()
	port := config.GetConfig().AppPort

	routers.InitAuthRoutes(mux)
	routers.InitSwaggerHandler(mux)

	log.Printf("[INFO] Server started at http://localhost:%s\n", port)

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatalf("[ERROR] Failed to start server: %v", err)
	}
}
