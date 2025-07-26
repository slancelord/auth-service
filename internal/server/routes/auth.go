package routes

import (
	"net/http"

	"auth-service/internal/handler"
)

func InitAuthRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /auth/token", handler.TokenHandler)
}
