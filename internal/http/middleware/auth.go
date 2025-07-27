package middleware

import (
	"auth-service/internal/model"
	"auth-service/internal/service"
	"context"
	"net/http"
	"strings"
)

type contextKey string

const SessionContextKey = contextKey("session")

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		currentSession, err := service.ValidateToken(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		context := context.WithValue(r.Context(), SessionContextKey, currentSession)
		next.ServeHTTP(w, r.WithContext(context))
	})
}

func Session(ctx context.Context) (*model.Session, bool) {
	session, ok := ctx.Value(SessionContextKey).(*model.Session)
	return session, ok
}
