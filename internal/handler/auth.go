package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"auth-service/internal/config"
	"auth-service/internal/http/middleware"
	"auth-service/internal/model"
	"auth-service/internal/repository"
	"auth-service/internal/service"
)

type TokenRequest struct {
	AccessToken  string `json:"access"`
	RefreshToken string `json:"refresh"`
}

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
	session, ok := middleware.Session((r.Context()))
	if !ok {
		http.Error(w, "Session not found", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(session.UserID))
}

// RefreshToken godoc
// @Summary Refresh access and refresh tokens
// @Description Refreshes the token pair using a valid access token and refresh token.
// @Tags auth
// @Accept json
// @Produce json
// @Param tokenRequest body TokenRequest true "Tokens for refresh"
// @Success 200 {object} map[string]string "New access and refresh tokens"
// @Failure 400 {string} string "Invalid request body or Invalid token"
// @Failure 500 {string} string "Failed to generate new tokens or save session"
// @Router /api/auth/refresh [post]
func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	sessionRepo := repository.NewRepository[model.Session]()
	refreshExpireDay := time.Duration(config.GetConfig().RefreshExpireDay)

	var tokenReq TokenRequest
	if err := json.NewDecoder(r.Body).Decode(&tokenReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	currentSession, err := service.ValidatePairToken(tokenReq.AccessToken, tokenReq.RefreshToken, r)
	if currentSession != nil {
		defer sessionRepo.Delete(currentSession.ID)
	}
	if err != nil {
		http.Error(w, "Invalid token", http.StatusBadRequest)
		return
	}

	if currentSession.IPAddress != r.RemoteAddr {
		go service.SendWebhook(map[string]string{
			"user_id": currentSession.UserID,
			"old_ip":  currentSession.IPAddress,
			"new_ip":  r.RemoteAddr,
			"time":    time.Now().Format(time.RFC3339),
		})
	}

	newAccess, newRefresh, tokenPairId, err := service.GenerateToken(currentSession.UserID)
	if err != nil {
		http.Error(w, "Failed to generate new tokens", http.StatusInternalServerError)
		return
	}

	newSession := model.Session{
		UserID:       currentSession.UserID,
		RefreshToken: service.HashToken(newRefresh),
		TokenPairId:  tokenPairId,
		UserAgent:    r.UserAgent(),
		IPAddress:    r.RemoteAddr,
		CreatedAt:    time.Now(),
		ExpiresAt:    time.Now().Add(refreshExpireDay * 24 * time.Hour),
	}

	err = sessionRepo.Create(&newSession)
	if err != nil {
		http.Error(w, "Failed to save session", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"access":  newAccess,
		"refresh": newRefresh,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// LogoutHandler godoc
// @Summary Logout user
// @Description Logout user
// @Tags auth
// @Accept json
// @Produce plain
// @Security ApiKeyAuth
// @Success 200 {string} string "Logout successful"
// @Failure 500 {string} string "Session not found or failed to logout"
// @Router /api/auth/logout [post]
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	sessionRepo := repository.NewRepository[model.Session]()

	session, ok := middleware.Session((r.Context()))
	if !ok {
		http.Error(w, "Session not found", http.StatusInternalServerError)
		return
	}

	err := sessionRepo.Delete(session.ID)
	if err != nil {
		http.Error(w, "Failed to logout", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logout successful"))
}
