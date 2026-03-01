// Package commoncomposer provides composition for common layer dependencies.
package commoncomposer

import (
	"log/slog"
	"os"

	"github.com/PratesJr/training-track-api/internal/common/ports"
	pkglogger "github.com/PratesJr/training-track-api/internal/pkg/logger"
)

// CommonContainer holds common layer dependencies.
type CommonContainer struct {
	Logger ports.Logger
}

// Compose initializes and returns the CommonContainer.
func Compose() *CommonContainer {
	base := slog.New(
		slog.NewJSONHandler(os.Stdout, nil),
	)

	logger := pkglogger.New(base)

	return &CommonContainer{
		Logger: logger,
	}
}
