package parsers

import (
	"context"
	"time"

	"github.com/PratesJr/training-track-api/internal/application/enums"
	exceptions2 "github.com/PratesJr/training-track-api/internal/application/exceptions"
)

func ParseHttpError(
	customErr exceptions2.ErrorType,
	ctx context.Context,
	err error,
) exceptions2.HttpException {

	var response exceptions2.HttpException

	response.Datetime = time.Now().Format(time.RFC3339)

	if customErr == nil && err != nil {
		response.StatusCode = enums.StatusCode.InternalServerError
		response.Description = "Unexpected error."
		return response
	}

	response.StatusCode = customErr.StatusCode()
	response.Description = customErr.Description()
	response.Details = customErr.Details()

	return response
}
