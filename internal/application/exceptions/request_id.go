// Package exceptions provides all application exceptions
package exceptions

import (
	"context"

	"github.com/PratesJr/training-track-api/internal/pkg/config"
)

func getRequestId(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	if id, ok := ctx.Value(config.RequestIDKey).(string); ok {
		return id
	}

	return ""
}
