// Package ports defines the interfaces for the external services and repositories, this file provides the interface of create user use case
package ports

import (
	"context"

	"github.com/PratesJr/training-track-api/internal/application/exceptions"
	"github.com/PratesJr/training-track-api/internal/domain/models"
)

// CreateUserUseCase defines the interface for the create user use case.
type CreateUserUseCase interface {
	Execute(ctx context.Context, input models.User) (exceptions.ErrorType, *models.User)
}
