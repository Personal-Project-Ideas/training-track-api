package handlers

import (
	"github.com/PratesJr/training-track-api/internal/application/dtos"
	"github.com/PratesJr/training-track-api/internal/application/exceptions"
	"github.com/PratesJr/training-track-api/internal/application/mappers"
	"github.com/PratesJr/training-track-api/internal/application/parsers"
	"github.com/PratesJr/training-track-api/internal/application/ports"
	"github.com/PratesJr/training-track-api/internal/application/validators"
	ports3 "github.com/PratesJr/training-track-api/internal/common/ports"
	ports2 "github.com/PratesJr/training-track-api/internal/domain/ports"
	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	createUser ports2.CreateUserUseCase
	logger     ports3.Logger
}

func (u userHandler) Create(c *fiber.Ctx) error {
	ctx := c.Context()

	u.logger.Info(ctx, "user_handler.Create.call")

	var body dtos.UserInputDTO
	if err := c.BodyParser(&body); err != nil {
		u.logger.Error(ctx, "user_handler.Create.error", err)
		ex := exceptions.BadRequestException(
			ctx,
			"invalid request body",
		)

		return c.Status(ex.StatusCode()).JSON(ex.ToMap())
	}

	if err := validators.ValidateDTO(ctx, body); err != nil {
		u.logger.Error(ctx, "user_handler.Create.error", err)

		ex := parsers.ParseHttpError(ctx, err, nil)
		if ex.StatusCode == 0 {
			return c.Status(422).JSON(ex)
		}
		return c.Status(ex.StatusCode).JSON(ex)
	}

	data := mappers.MapCreateInputToUser(body)

	executeErr, response := u.createUser.Execute(ctx, *data)

	if executeErr != nil {
		u.logger.Error(ctx, "user_handler.Create.error", executeErr)

		exception := parsers.ParseHttpError(nil, ctx, executeErr)

		return c.Status(exception.StatusCode).JSON(exception)
	}

	u.logger.Info(ctx, "user_handler.Create.result")

	return c.Status(201).JSON(response)
}

func (u userHandler) GetByID(c *fiber.Ctx) error {
	return c.Status(501).SendString("Not Implemented.")
}

// UserHandler creates a new instance of the user handler.
func UserHandler(createUser *ports2.CreateUserUseCase, logger *ports3.Logger) ports.UserHandler {
	return &userHandler{
		createUser: *createUser,
		logger:     *logger,
	}
}
