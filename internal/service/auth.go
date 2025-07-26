package service

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"auth-service/internal/config"
)

func GenerateToken(guid string) (string, string, error) {
	secretKey := config.GetConfig().JWTSecret
	bytes := make([]byte, 32)

	now := time.Now()
	exp := now.Add(15 * time.Minute)

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub": guid,
		"iat": now,
		"exp": exp,
	})

	accessString, err := accessToken.SignedString(secretKey)
	if err != nil {
		return "", "", err
	}

	_, err = rand.Read(bytes)
	if err != nil {
		return "", "", err
	}

	refreshString := base64.StdEncoding.EncodeToString(bytes)

	return accessString, refreshString, nil
}
