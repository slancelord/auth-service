package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	JWTSecret        string
	AccessExpireMin  int
	RefreshExpireDay int

	WebhookUrl string
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

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("[WARN] Missing .env file: %v", err)
	}

	config = &Config{
		AppPort: getEnv("APP_PORT", "8080"),

		DBHost:     getEnv("POSTGRES_HOST", "localhost"),
		DBPort:     getEnv("POSTGRES_PORT", "5432"),
		DBUser:     getEnv("POSTGRES_USER", "user"),
		DBPassword: getEnv("POSTGRES_PASSWORD", "password"),
		DBName:     getEnv("POSTGRES_DB", "db"),

		AccessExpireMin:  getEnvAsInt("ACCESS_TOKEN_EXPIRE_MINUTES", 5),
		RefreshExpireDay: getEnvAsInt("REFRESH_TOKEN_EXPIRE_DAYS", 7),
		JWTSecret:        getEnv("JWT_SECRET", "defaultsecret"),

		WebhookUrl: getEnv("WEBHOOK_URL", "http://localhost:8080/webhook"),
	}
}
