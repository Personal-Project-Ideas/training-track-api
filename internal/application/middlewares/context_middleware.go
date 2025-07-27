// Package middlewares   provides the context middleware for the application.
package middlewares

import (
	"context"
	"github.com/PratesJr/training-track-api/internal/application/configuration"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type contextMiddleware struct{}

// Handler NewContextMiddleware creates a new instance of ContextMiddleware.
func (m *contextMiddleware) Handler(c *fiber.Ctx) error {
	reqID := c.Get(string(configuration.RequestIDKey))
	if reqID == "" {
		reqID = uuid.NewString()
	}

	ctx := context.WithValue(context.Background(), configuration.RequestIDKey, reqID)
	c.SetUserContext(ctx)
	c.Set("X-Request-ID", reqID)

	return c.Next()
}
