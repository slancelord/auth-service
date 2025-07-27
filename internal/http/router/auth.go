package router

import (
	"net/http"

	"auth-service/internal/handler"
	"auth-service/internal/http/middleware"
)

func InitAuthRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /auth/user", handler.UserHandler)

	mux.HandleFunc("POST /auth/token", handler.TokenHandler)
}
