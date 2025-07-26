package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort   string
	JWTSecret string
}

var config *Config

func getEnv(key string, defaultValue string) string {
	value, exist := os.LookupEnv(key)
	if exist {
		return value
	}

	return defaultValue
}

func GetConfig() *Config {
	if config == nil {
		log.Fatal("[ERROR] Config is not initialized")
	}

	return config
}

func InitConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("[WARN] Missing .env file")
	}

	config = &Config{
		AppPort:   getEnv("APP_PORT", "8080"),
		JWTSecret: getEnv("JWT_SECRET", "defaultsecret"),
	}
}
