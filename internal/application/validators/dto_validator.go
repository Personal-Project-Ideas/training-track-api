package validators

import (
	"context"
	"errors"

	"github.com/PratesJr/training-track-api/internal/application/exceptions"
	"github.com/go-playground/validator/v10"
)

func ValidateDTO(dto interface{}, ctx context.Context) exceptions.ErrorType {
	validate := validator.New()
	err := validate.Struct(dto)

	if err == nil {
		return nil
	}

	var validationErr validator.ValidationErrors
	errors.As(err, &validationErr)

	var details []exceptions.ErrorDetails

	for _, e := range validationErr {
		switch e.Tag() {
		case "required":
			details = append(details,
				exceptions.NewErrorDetail(e.Field(), "is required"),
			)
		case "min":
			details = append(details,
				exceptions.NewErrorDetail(e.Field(), "min "+e.Param()),
			)
		case "email":
			details = append(details,
				exceptions.NewErrorDetail(e.Field(), "invalid email"),
			)
		}
	}

	return exceptions.UnprocessableEntityException(ctx, "invalid payload", details)
}
