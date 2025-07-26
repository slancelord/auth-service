package config

import (
	"os"
)

type Config struct {
	AppPort string
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		config = &Config{
			AppPort: getEnv("APP_PORT", "8080"),
		}
	}

	return config
}

func getEnv(key string, defaultValue string) string {
	var value, exist = os.LookupEnv(key)
	if exist {
		return value
	}

	return defaultValue
}
