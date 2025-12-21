package config

import (
	"os"
)

// Env holds all application configuration.
type Env struct {
	AppPort string
	DbURL   string
}

// LoadConfig reads env vars and returns a Config instance.
func LoadConfig() *Env {
	return &Env{
		AppPort: getEnv("APP_PORT", "8080"),
		DbURL:   getEnv("DATABASE_URL", ""),
	}
}

// getEnv loads env var os default value
func getEnv(key string, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
