package application_composer

import (
	"github.com/PratesJr/training-track-api/internal/application/handlers"
	"github.com/PratesJr/training-track-api/internal/application/middlewares"
	"github.com/PratesJr/training-track-api/internal/application/ports"
	"github.com/PratesJr/training-track-api/internal/application/usecases"
	ports3 "github.com/PratesJr/training-track-api/internal/common/ports"
	ports2 "github.com/PratesJr/training-track-api/internal/domain/ports"
)

type ApplicationContainer struct {
	UserHandler ports.UserHandler
	createUser  *ports2.CreateUserUseCase
	Middlewares []middlewares.Middleware
}

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
