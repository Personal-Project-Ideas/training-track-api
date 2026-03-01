// Package applicationcomposer provides composition for application layer dependencies.
package applicationcomposer

import (
	"github.com/PratesJr/training-track-api/internal/application/handlers"
	"github.com/PratesJr/training-track-api/internal/application/middlewares"
	"github.com/PratesJr/training-track-api/internal/application/ports"
	"github.com/PratesJr/training-track-api/internal/application/usecases"
	ports3 "github.com/PratesJr/training-track-api/internal/common/ports"
	ports2 "github.com/PratesJr/training-track-api/internal/domain/ports"
)

// ApplicationContainer holds application layer dependencies.
type ApplicationContainer struct {
	UserHandler ports.UserHandler
	createUser  *ports2.CreateUserUseCase
	Middlewares []middlewares.Middleware
}

// Compose initializes and returns the ApplicationContainer.
func Compose(
	userRepository ports2.UserRepository,
	logger *ports3.Logger,
	pwdValidator ports2.ValidatePwd,
) *ApplicationContainer {
	createUserUseCase := usecases.CreateUserUseCase(&userRepository, logger, &pwdValidator)

	userHandler := handlers.UserHandler(&createUserUseCase, logger)

	return &ApplicationContainer{
		createUser:  &createUserUseCase,
		UserHandler: userHandler,
	}
}
