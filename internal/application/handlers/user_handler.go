// Package handlers  provides the implementation of the user handler interface.
package handlers

import (
	"github.com/PratesJr/training-track-api/internal/application/dtos"
	"github.com/PratesJr/training-track-api/internal/application/exceptions"
	"github.com/PratesJr/training-track-api/internal/application/parsers"
	"github.com/PratesJr/training-track-api/internal/application/ports"
	"github.com/PratesJr/training-track-api/internal/application/validators"
	ports2 "github.com/PratesJr/training-track-api/internal/domain/ports"
	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	createUser *ports2.CreateUserUseCase
}

func (u userHandler) Create(c *fiber.Ctx) error {
	ctx := c.Context()
	var body dtos.UserInputDTO
	if err := c.BodyParser(&body); err != nil {
		ex := exceptions.BadRequestException(
			ctx,
			"invalid request body",
		)

		return c.Status(ex.StatusCode()).JSON(ex.ToMap())
	}

	if err := validators.ValidateDTO(body, ctx); err != nil {

		ex := parsers.ParseHttpError(err, ctx, nil)

		return c.Status(ex.StatusCode).JSON(ex)
	}

	return c.Status(501).SendString(
		"Not Implemented.")
}

func (u userHandler) GetByID(c *fiber.Ctx) error {
	return c.Status(501).SendString("Not Implemented.")
}

// UserHandler  creates a new instance of the user handler.
func UserHandler(createUser ports2.CreateUserUseCase) ports.UserHandler {
	return &userHandler{
		createUser: &createUser,
	}
}
