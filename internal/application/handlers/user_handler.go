// Package handlers  provides the implementation of the user handler interface.
package handlers

import (
	"github.com/PratesJr/training-track-api/internal/application/ports"
	"github.com/gofiber/fiber/v2"
)

type userHandler struct{}

func (u userHandler) Create(c *fiber.Ctx) error {
	return c.Status(501).SendString("Not Implemented")
}

func (u userHandler) GetByID(c *fiber.Ctx) error {
	return c.Status(501).SendString("Not Implemented")
}

// NewUserHandler creates a new instance of the user handler.
func NewUserHandler() ports.UserHandler {
	return &userHandler{}
}
