// Package bootstrap is responsible for initializing and configuring the application components.
package bootstrap

import (
	application_composer "github.com/PratesJr/training-track-api/internal/bootstrap/composition/application"
	domain_composer "github.com/PratesJr/training-track-api/internal/bootstrap/composition/domain"
	infraestructure_composer "github.com/PratesJr/training-track-api/internal/bootstrap/composition/infraestructure"
	"github.com/gofiber/fiber/v2"
)

type AppContainer struct {
	App                 *fiber.App
	InfraComposer       *infraestructure_composer.InfraContainer
	ApplicationComposer *application_composer.ApplicationContainer
	DomainComposer      *domain_composer.DomainContainer
}

func Bootstrap() *AppContainer {
	app := fiber.New()

	infra := infraestructure_composer.Compose()

	application := application_composer.Compose(infra.UserRepo)

	domain := domain_composer.Compose()

	return &AppContainer{
		App:                 app,
		InfraComposer:       infra,
		ApplicationComposer: application,
		DomainComposer:      domain,
	}
}
