package config

import (
	"os"

	"github.com/PratesJr/training-track-api/internal/helpers"
	"github.com/joho/godotenv"
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
	err := godotenv.Load(".env")
	if err != nil {
		return nil
	}
	return &Env{
		AppPort: helpers.ToPointer(os.Getenv("APP_PORT")),
		DbURL:   helpers.ToPointer(os.Getenv("DATABASE_HOST")),
		DbUser:  helpers.ToPointer(os.Getenv("DATABASE_USER")),
		DbPass:  helpers.ToPointer(os.Getenv("DATABASE_PASSWD")),
		DbName:  helpers.ToPointer(os.Getenv("DATABASE_NAME")),
		DbPort:  helpers.ToPointer(os.Getenv("DATABASE_PORT")),
		Env:     helpers.ToPointer(os.Getenv("ENV")),
	}
}
