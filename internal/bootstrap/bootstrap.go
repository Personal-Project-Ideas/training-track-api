// Package bootstrap is responsible for initializing and configuring the application components.
package bootstrap

import (
	"github.com/PratesJr/training-track-api/internal/application/middlewares"
	"github.com/PratesJr/training-track-api/internal/application/routes"
	application_composer "github.com/PratesJr/training-track-api/internal/bootstrap/composition/application"
	common_composer "github.com/PratesJr/training-track-api/internal/bootstrap/composition/common"
	domain_composer "github.com/PratesJr/training-track-api/internal/bootstrap/composition/domain"
	infraestructure_composer "github.com/PratesJr/training-track-api/internal/bootstrap/composition/infraestructure"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AppContainer struct {
	App                 *fiber.App
	InfraComposer       *infraestructure_composer.InfraContainer
	ApplicationComposer *application_composer.ApplicationContainer
	DomainComposer      *domain_composer.DomainContainer
}

func Bootstrap(testDB ...*gorm.DB) *AppContainer {
	var globalMiddlewares []middlewares.Middleware
	var userMiddlewares []middlewares.Middleware

	app := fiber.New()

	commons := common_composer.Compose()

	infra := infraestructure_composer.Compose(commons.Logger, testDB...)

	domain := domain_composer.Compose(commons.Logger)

	application := application_composer.Compose(infra.UserRepo, &commons.Logger, domain.ValidatePass)

	globalMiddlewares = append(globalMiddlewares, middlewares.NewContextMiddleware())

	routes.SetupRouter(
		app,
		application.UserHandler,
		globalMiddlewares,
		userMiddlewares,
	)

	return &AppContainer{
		App:                 app,
		InfraComposer:       infra,
		ApplicationComposer: application,
		DomainComposer:      domain,
	}
}
