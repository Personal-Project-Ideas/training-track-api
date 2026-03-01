// Package ports provides interfaces for common services like logging.
package ports

import "context"

// Logger provides logging capabilities for the application.
type Logger interface {
	Info(ctx context.Context, msg string, attrs ...any)
	Warn(ctx context.Context, msg string, attrs ...any)
	Error(ctx context.Context, msg string, attrs ...any)
}
