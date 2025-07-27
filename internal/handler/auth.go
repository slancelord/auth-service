package handler

import (
	"encoding/json"
	"net/http"

	"auth-service/internal/service"
)

// TokenHandler godoc
// @Summary Get tokens
// @Description Generates access and refresh tokens using a user GUID
// @Tags auth
// @Accept json
// @Produce json
// @Param User-Agent header string true "User Agent" Enums(Mozilla/5.0, Opera/9.80)
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "Missing GUID"
// @Failure 500 {string} string "Failed to generate tokens"
// @Router /api/auth/token [post]
func TokenHandler(w http.ResponseWriter, r *http.Request) {
	refreshExpireDay := time.Duration(config.GetConfig().RefreshExpireDay)
	sessionRepo := repository.NewRepository[model.Session]()

	guid := r.URL.Query().Get("guid")
	if guid == "" {
		http.Error(w, "Missing guid", http.StatusBadRequest)
	}

	accessToken, refreshToken, tokenPairId, err := service.GenerateToken(guid)
	if err != nil {
		http.Error(w, "Failed to generate tokens", http.StatusInternalServerError)
		return
	}

	session := model.Session{
		UserID:       guid,
		RefreshToken: service.HashToken(refreshToken),
		TokenPairId:  tokenPairId,
		UserAgent:    r.UserAgent(),
		IPAddress:    r.RemoteAddr,
		CreatedAt:    time.Now(),
		ExpiresAt:    time.Now().Add(refreshExpireDay * 24 * time.Hour),
	}

	err = sessionRepo.Create(&session)
	if err != nil {
		http.Error(w, "Failed to save session", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"access":  accessToken,
		"refresh": refreshToken,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// UserHandler godoc
// @Summary Get current user ID
// @Description Returns the GUID of the currently authenticated user using a valid access token.
// @Tags auth
// @Accept json
// @Produce plain
// @Security ApiKeyAuth
// @Success 200 {string} string "User GUID"
// @Failure 401 {string} string "Unauthorized: Missing or invalid token"
// @Failure 500 {string} string "Session not found"
// @Router /api/auth/user [get]
func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test Message"))
}
