// Package exceptions provides all application exceptions
package exceptions

import (
	"context"
	"net/http"
)

// BadRequestException returns a Bad Request error.
func BadRequestException(ctx context.Context, errMessage string, details ...ErrorDetails) ErrorType {
	ex := errorType{
		id:          GetRequestID(ctx),
		statusCode:  http.StatusBadRequest,
		code:        "400_BAD_REQUEST",
		details:     details,
		description: "Bad Request.",
		message:     errMessage,
	}
	return &ex
}

// UnauthorizedException returns an Unauthorized error.
func UnauthorizedException(ctx context.Context, errMessage string) ErrorType {
	ex := errorType{
		id:          GetRequestID(ctx),
		statusCode:  http.StatusUnauthorized,
		code:        "401_UNAUTHORIZED",
		description: "Unauthorized.",
		message:     errMessage,
	}
	return &ex
}

// UnprocessableEntityException returns an Unprocessable Entity error.
func UnprocessableEntityException(ctx context.Context, errMessage string, details []ErrorDetails) ErrorType {
	ex := errorType{
		id:          GetRequestID(ctx),
		statusCode:  http.StatusUnprocessableEntity,
		code:        "422_UNPROCESSABLE_ENTITY",
		details:     details,
		description: "Unprocessable entity.",
		message:     errMessage,
	}
	return &ex
}

// NotFoundException returns a Not Found error.
func NotFoundException(ctx context.Context, errMessage string) ErrorType {
	ex := errorType{
		id:          GetRequestID(ctx),
		statusCode:  http.StatusNotFound,
		code:        "404_NOT_FOUND",
		description: "Data Not Found.",
		message:     errMessage,
	}
	return &ex
}

// InternalServerErrorException returns an Internal Server Error.
func InternalServerErrorException(ctx context.Context, errMessage string) ErrorType {
	ex := errorType{
		id:          GetRequestID(ctx),
		statusCode:  http.StatusInternalServerError,
		code:        "500_INTERNAL_ERROR",
		description: "Something went wrong.",
		message:     errMessage,
	}
	return &ex
}
