package server

import (
	"net/http"

	"auth-service/internal/config"
	"auth-service/internal/server/routes"
)

func Serve() {
	var mux = http.NewServeMux()
	var port = config.GetConfig().AppPort

	routes.InitAuthRoutes(mux)

	http.ListenAndServe(":"+port, mux)
}
