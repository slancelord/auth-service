package handlers

import (
	"encoding/json"
	"net/http"

	"auth-service/internal/services"
)

// TokenHandler godoc
// @Summary Получить токены
// @Description Генерирует access и refresh токены по GUID
// @Tags auth
// @Accept json
// @Produce json
// @Param guid query string true "GUID пользователя"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "Missing guid"
// @Failure 500 {string} string "Failed to generate tokens"
// @Router /auth/token [post]
func TokenHandler(w http.ResponseWriter, r *http.Request) {
	guid := r.URL.Query().Get("guid")
	if guid == "" {
		http.Error(w, "Missing guid", http.StatusBadRequest)
	}

	accessToken, refreshToken, err := services.GenerateToken(guid)
	if err != nil {
		http.Error(w, "Failed to generate tokens", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"access":  accessToken,
		"refresh": refreshToken,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test Message"))
}
