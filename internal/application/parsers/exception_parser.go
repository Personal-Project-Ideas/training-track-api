package parsers

import (
	"context"
	"time"

	"github.com/PratesJr/training-track-api/internal/application/enums"
	exceptions2 "github.com/PratesJr/training-track-api/internal/application/exceptions"
	"github.com/PratesJr/training-track-api/internal/pkg/config"
)

// ParseHttpError parses an error and returns an HttpException.
func ParseHttpError(
	ctx context.Context,
	customErr exceptions2.ErrorType,
	err error,
) exceptions2.HttpException {

	var response exceptions2.HttpException

	response.Datetime = time.Now().Format(time.RFC3339)

	if reqID, ok := ctx.Value(config.RequestIDKey).(string); ok {
		response.Id = reqID
	} else if customErr != nil {
		if idGetter, ok := any(customErr).(interface{ Id() string }); ok {
			response.Id = idGetter.Id()
		}
	}

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
