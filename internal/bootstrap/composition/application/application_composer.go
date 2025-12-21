package application_composer

import (
	"github.com/PratesJr/training-track-api/internal/application/middlewares"
	"github.com/PratesJr/training-track-api/internal/application/ports"
	ports2 "github.com/PratesJr/training-track-api/internal/domain/ports"
)

type ApplicationContainer struct {
	UserHandler ports.UserHandler
	createUser  ports2.CreateUserUseCase
	Middlewares []middlewares.Middleware
}

func Compose() *ApplicationContainer {
	applicationContainer := &ApplicationContainer{}

	return applicationContainer
}
