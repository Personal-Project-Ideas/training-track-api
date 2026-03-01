// Package logger provides logging implementations for the application.
package logger

import (
	"context"
	"log/slog"
	"strings"

	"github.com/PratesJr/training-track-api/internal/common/ports"
	"github.com/PratesJr/training-track-api/internal/pkg/config"
)

// SlogLogger is a structured logger implementation using slog.
type SlogLogger struct {
	log       *slog.Logger
	sensitive map[string]struct{}
}

// New creates a new SlogLogger instance.
func New(base *slog.Logger) ports.Logger {
	return &SlogLogger{
		log: base,
		sensitive: map[string]struct{}{
			"password":      {},
			"token":         {},
			"secret":        {},
			"authorization": {},
			"api_key":       {},
		},
	}
}

func (l *SlogLogger) sanitize(attrs []any) []any {

	sanitized := make([]any, 0, len(attrs))

	for i := 0; i < len(attrs); i += 2 {

		if i+1 >= len(attrs) {
			break
		}

		key, ok := attrs[i].(string)
		if !ok {
			sanitized = append(sanitized, attrs[i], attrs[i+1])
			continue
		}

		value := attrs[i+1]

		if _, isSensitive := l.sensitive[strings.ToLower(key)]; isSensitive {
			sanitized = append(sanitized, key, "******")
			continue
		}

		sanitized = append(sanitized, key, value)
	}

	return sanitized
}

func (l *SlogLogger) enrich(ctx context.Context, attrs []any) []any {
	if requestID, ok := ctx.Value(config.RequestIDKey).(string); ok {
		attrs = append(attrs, "request_id", requestID)
	}

	return attrs
}

// Info logs an info message with optional attributes.
func (l *SlogLogger) Info(ctx context.Context, msg string, attrs ...any) {
	attrs = l.sanitize(attrs)
	attrs = l.enrich(ctx, attrs)
	l.log.InfoContext(ctx, msg, attrs...)
}

// Warn logs a warning message with optional attributes.
func (l *SlogLogger) Warn(ctx context.Context, msg string, attrs ...any) {
	attrs = l.sanitize(attrs)
	attrs = l.enrich(ctx, attrs)
	l.log.WarnContext(ctx, msg, attrs...)
}

func (l *SlogLogger) Error(ctx context.Context, msg string, attrs ...any) {
	attrs = l.sanitize(attrs)
	attrs = l.enrich(ctx, attrs)
	l.log.ErrorContext(ctx, msg, attrs...)
}
