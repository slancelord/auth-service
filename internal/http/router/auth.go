package router

import (
	"net/http"

	"auth-service/internal/handler"
	"auth-service/internal/http/middleware"
)

func InitAuthRoutes(mux *http.ServeMux) {
	mux.Handle("GET /api/auth/user", middleware.Auth(http.HandlerFunc(handler.UserHandler)))

	mux.HandleFunc("POST /api/auth/token", handler.TokenHandler)
	mux.HandleFunc("POST /api/auth/refresh", handler.RefreshTokenHandler)
	mux.Handle("POST /api/auth/logout", middleware.Auth(http.HandlerFunc(handler.LogoutHandler)))
}
