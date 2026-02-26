package logger

import (
	"context"
	"log/slog"

	"github.com/PratesJr/training-track-api/internal/common/ports"
)

type SlogLogger struct {
	log *slog.Logger
}

func New(base *slog.Logger) ports.Logger {
	return &SlogLogger{log: base}
}

func (l *SlogLogger) Info(ctx context.Context, msg string, attrs ...any) {
	l.log.InfoContext(ctx, msg, attrs...)
}

func (l *SlogLogger) Warn(ctx context.Context, msg string, attrs ...any) {
	l.log.WarnContext(ctx, msg, attrs...)
}

func (l *SlogLogger) Error(ctx context.Context, msg string, attrs ...any) {
	l.log.ErrorContext(ctx, msg, attrs...)
}
