// Package routes provides the main router for the application.
package routes

import (
	"github.com/PratesJr/training-track-api/internal/application/middlewares"
	"github.com/PratesJr/training-track-api/internal/application/ports"
	"github.com/gofiber/fiber/v2"
)

// SetupRouter initializes the main router for the application.
func SetupRouter(app *fiber.App, userHandler ports.UserHandler, globalMws []middlewares.Middleware, userMws []middlewares.Middleware) {
	var routes fiber.Router
	var globalFiberMws []fiber.Handler
	for _, mw := range globalMws {
		globalFiberMws = append(globalFiberMws, mw.Handle)
	}
	if len(globalMws) > 0 {
		routes = app.Group("/", globalFiberMws...)
	}

	//Convert user specific middlewares
	var userFiberMws []fiber.Handler
	if len(userMws) > 0 {
		for _, mw := range userMws {
			userFiberMws = append(userFiberMws, mw.Handle)
		}

	}

	SetupUserRoutes(routes, userHandler, userFiberMws)
}
