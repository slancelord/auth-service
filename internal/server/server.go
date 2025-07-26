package server

import (
	"net/http"

	"auth-service/internal/server/routes"
)

func Serve() {
	var mux = http.NewServeMux()

	routes.InitAuthRoutes(mux)

	http.ListenAndServe("localhost:8080", mux)
}
