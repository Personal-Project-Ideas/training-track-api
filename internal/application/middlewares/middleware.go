// Package middlewares provides the middleware interfaces for the application.
package middlewares

import "github.com/gofiber/fiber/v2"

// Middleware defines the signature for Fiber middlewares
type Middleware interface {
	Handle(*fiber.Ctx) error
}
