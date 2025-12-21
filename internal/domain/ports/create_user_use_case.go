// Package ports defines the interfaces for the external services and repositories, this file provides the interface of create user use case
package ports

import "github.com/PratesJr/training-track-api/internal/domain/models"

type CreateUserUseCase interface {
	Execute(input models.User) (*models.User, error)
}
