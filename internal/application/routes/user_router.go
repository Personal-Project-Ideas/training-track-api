// Package routes  define specific routes for the user handler
package routes

import (
	"github.com/PratesJr/training-track-api/internal/application/ports"
	"github.com/gofiber/fiber/v2"
)

// SetupUserRoutes sets up the user-related routes for the application.
func SetupUserRoutes(router fiber.Router, handler ports.UserHandler, middlewares []fiber.Handler) {

	userGroup := router.Group("/users")
	userGroup.Get("/:id", handler.GetByID)
	userGroup.Post("/", handler.Create)

	if len(middlewares) > 0 {
		userGroup.Use(middlewares)
	}
}
