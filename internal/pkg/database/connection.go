// Package database provides database connection utilities for the application.
package database

import (
	"fmt"

	config "github.com/PratesJr/training-track-api/internal/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DbConnection returns a new database connection using GORM.
func DbConnection() *gorm.DB {
	properties := config.LoadConfig()

	if properties == nil {
		panic("Failed to load database config")
	}

	dns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		*properties.DbURL,
		*properties.DbUser,
		*properties.DbPass,
		*properties.DbName,
		*properties.DbPort,
	)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		panic("Failed to connect on database")
	}

	return db
}
