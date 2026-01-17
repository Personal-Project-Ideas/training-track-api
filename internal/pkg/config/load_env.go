package config

import (
	"os"

	"github.com/PratesJr/training-track-api/internal/helpers"
)

// Env holds all application configuration.
type Env struct {
	AppPort *string
	DbURL   *string
	DbUser  *string
	DbPass  *string
	DbName  *string
	DbPort  *string
	Env     *string
}

// LoadConfig reads env vars and returns a Config instance.
func LoadConfig() *Env {
	return &Env{
		AppPort: helpers.ToPointer(getEnv("APP_PORT", "8080")),
		DbURL:   helpers.ToPointer(getEnv("DATABASE_URL", "")),
		DbUser:  helpers.ToPointer(getEnv("DATABASE_USER", "")),
		DbPass:  helpers.ToPointer(getEnv("DATABASE_PASSWD", "")),
		DbName:  helpers.ToPointer(getEnv("DATABASE_NAME", "")),
		DbPort:  helpers.ToPointer(getEnv("DATABASE_PORT", "")),
		Env:     helpers.ToPointer(getEnv("ENV", "")),
	}
}

// getEnv loads env var os default value
func getEnv(key string, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
