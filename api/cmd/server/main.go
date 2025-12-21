// Package main is the entry point for the API server.
package main

import (
	"github.com/PratesJr/training-track-api/internal/bootstrap"
	"github.com/PratesJr/training-track-api/internal/pkg/config"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	app := bootstrap.Bootstrap()

	log.Println("Starting server on :3000")

	if err := app.App.Listen(cfg.AppPort); err != nil {
		log.Fatal(err)
	}

}
