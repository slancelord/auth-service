package routers

import (
	"net/http"

	"auth-service/internal/handlers"
)

func InitAuthRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /auth/user", handlers.UserHandler)

	mux.HandleFunc("POST /auth/token", handlers.TokenHandler)
}
