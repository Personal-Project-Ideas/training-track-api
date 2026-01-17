// Package middlewares   provides the context middleware for the application.
package middlewares

import (
	"context"

	"github.com/PratesJr/training-track-api/internal/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type contextMiddleware struct{}

func (m *contextMiddleware) Handle(c *fiber.Ctx) error {
	reqID := c.Get(string(config.RequestIDKey))
	if reqID == "" {
		reqID = uuid.NewString()
	}

	ctx := context.WithValue(c.UserContext(), config.RequestIDKey, reqID)
	c.SetUserContext(ctx)
	c.Set("X-Request-ID", reqID)

	return c.Next()
}

// NewContextMiddleware  creates a new instance of ContextMiddleware.
func NewContextMiddleware() Middleware {
	return &contextMiddleware{}
}
