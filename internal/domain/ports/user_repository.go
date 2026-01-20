// Package ports defines the interfaces for the external services and repositories, this file provides the interface of user repository
package ports

import (
	"context"
	"github.com/PratesJr/training-track-api/internal/domain/models"
)

type UserRepository interface {
	Create(ctx context.Context, user models.User) (models.User, error)
	Update(ctx context.Context, user models.User) (models.User, error)
	FindByID(ctx context.Context, id string) (models.User, error)
	FindByEmail(ctx context.Context, email string) (models.User, error)
	List(ctx context.Context) ([]models.User, error)
}
