// Package ports on application provides ports for handlers and other application components to external world communication.
package ports

import "github.com/gofiber/fiber/v2"

// UserHandler defines the interface for user-related operations.
type UserHandler interface {
	Create(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
}
