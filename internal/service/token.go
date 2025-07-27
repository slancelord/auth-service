package service

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"auth-service/internal/config"
	"auth-service/internal/model"
	"auth-service/internal/repository"
)

func GenerateToken(guid string) (string, string, string, error) {
	accessExpireMin := time.Duration(config.GetConfig().AccessExpireMin)
	secretKey := config.GetConfig().JWTSecret
	bytes := make([]byte, 32)
	tokenPairId := uuid.New().String()

	now := time.Now()
	exp := now.Add(accessExpireMin * time.Minute)

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub": guid,
		"iat": now.Unix(),
		"exp": exp.Unix(),
		"jti": tokenPairId,
	})

	accessString, err := accessToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", "", "", err
	}

	_, err = rand.Read(bytes)
	if err != nil {
		return "", "", "", err
	}

	refreshString := base64.StdEncoding.EncodeToString(bytes)

	return accessString, refreshString, tokenPairId, nil
}

func ValidateToken(tokenString string) (*model.Session, error) {
	secretKey := config.GetConfig().JWTSecret
	claims := jwt.MapClaims{}
	sessionRepo := repository.NewRepository[model.Session]()

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	guid, ok := claims["sub"].(string)
	if !ok {
		return nil, errors.New("invalid token")
	}

	currentSession, err := sessionRepo.FirstByField("user_id", guid)
	if err != nil {
		return nil, errors.New("session not found")
	}

	if exp, ok := claims["exp"].(float64); ok {
		if int64(exp) < time.Now().Unix() {
			return currentSession, errors.New("token expired")
		}
	}

	return currentSession, nil
}

func ValidateTokenSkipExp(tokenString string) (*model.Session, error) {
	secretKey := config.GetConfig().JWTSecret
	claims := jwt.MapClaims{}
	sessionRepo := repository.NewRepository[model.Session]()

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	}, jwt.WithoutClaimsValidation())

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	guid, ok := claims["sub"].(string)
	if !ok {
		return nil, errors.New("invalid token")
	}

	currentSession, err := sessionRepo.FirstByField("user_id", guid)
	if err != nil {
		return nil, errors.New("session not found")
	}

	return currentSession, nil
}

func ValidatePairToken(accessString string, refreshString string, r *http.Request) (*model.Session, error) {
	secretKey := config.GetConfig().JWTSecret
	claims := jwt.MapClaims{}
	sessionRepo := repository.NewRepository[model.Session]()

	token, err := jwt.ParseWithClaims(accessString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	}, jwt.WithoutClaimsValidation())

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	id, ok := claims["jti"].(string)
	if !ok {
		return nil, errors.New("invalid token")
	}

	currentSession, err := sessionRepo.FirstByField("token_pair_id", id)
	if err != nil {
		return nil, errors.New("session not found")
	}

	if id != currentSession.TokenPairId {
		return currentSession, errors.New("token pair mismatch")
	}

	if currentSession.UserAgent != r.UserAgent() {
		return currentSession, errors.New("user-Agent mismatch")
	}

	if HashToken(refreshString) != currentSession.RefreshToken {
		return currentSession, errors.New("invalid refresh token")
	}

	if time.Now().After(currentSession.ExpiresAt) {
		return currentSession, errors.New("refresh token expired")
	}

	return currentSession, nil
}

func HashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}
