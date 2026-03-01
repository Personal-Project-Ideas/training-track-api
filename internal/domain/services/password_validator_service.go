// Package domainservices provides domain service implementations.
package domainservices

import (
	"context"
	"unicode"

	"github.com/PratesJr/training-track-api/internal/application/exceptions"
	"github.com/PratesJr/training-track-api/internal/common/ports"
	ports2 "github.com/PratesJr/training-track-api/internal/domain/ports"
)

type validatePass struct {
	logger ports.Logger
}

func (v validatePass) ValidatePassWord(ctx context.Context, password string) exceptions.ErrorType {
	v.logger.Info(ctx, "password_validator_service.ValidatePassWord.call")
	if len(password) < 8 {
		return exceptions.UnprocessableEntityException(ctx, "password too short", []exceptions.ErrorDetails{exceptions.NewErrorDetail("password", "min 8")})
	}

	var hasUpper, hasLower, hasSpecial bool

	for _, r := range password {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsPunct(r) || unicode.IsSymbol(r):
			hasSpecial = true
		}
	}

	if !hasUpper {
		v.logger.Error(ctx, "password_validator_service.ValidatePassWord.error")
		return exceptions.UnprocessableEntityException(ctx, "missing uppercase", []exceptions.ErrorDetails{exceptions.NewErrorDetail("password", "must contain uppercase")})
	}
	if !hasLower {
		v.logger.Error(ctx, "password_validator_service.ValidatePassWord.error")
		return exceptions.UnprocessableEntityException(ctx, "missing lowercase", []exceptions.ErrorDetails{exceptions.NewErrorDetail("password", "must contain lowercase")})
	}
	if !hasSpecial {
		v.logger.Error(ctx, "password_validator_service.ValidatePassWord.error")
		return exceptions.UnprocessableEntityException(ctx, "missing special char", []exceptions.ErrorDetails{exceptions.NewErrorDetail("password", "must contain special char")})
	}

	v.logger.Info(ctx, "password_validator_service.ValidatePassWord.end")
	return nil
}

// ValidatePassWordService creates a new password validator service.
func ValidatePassWordService(logger *ports.Logger) ports2.ValidatePwd {
	return &validatePass{logger: *logger}

}
