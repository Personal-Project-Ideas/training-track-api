package ports

import (
	"context"

	"github.com/PratesJr/training-track-api/internal/application/exceptions"
)

// ValidatePwd defines the interface for password validation.
type ValidatePwd interface {
	ValidatePassWord(ctx context.Context, pwd string) exceptions.ErrorType
}
