package domain_services

import (
	"context"
	"unicode"

	"github.com/PratesJr/training-track-api/internal/common/ports"
	domain_errors "github.com/PratesJr/training-track-api/internal/domain/errors"
	ports2 "github.com/PratesJr/training-track-api/internal/domain/ports"
)

type validatePass struct {
	logger ports.Logger
}

func (v validatePass) ValidatePassWord(ctx context.Context, password string) error {
	v.logger.Info(ctx, "password_validator_service.ValidatePassWord.call")
	if len(password) < 8 {
		return domain_errors.ErrPasswordTooShort
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
		return domain_errors.ErrPasswordMissingUppercase
	}
	if !hasLower {
		v.logger.Error(ctx, "password_validator_service.ValidatePassWord.error")
		return domain_errors.ErrPasswordMissingLowercase
	}
	if !hasSpecial {
		v.logger.Error(ctx, "password_validator_service.ValidatePassWord.error")
		return domain_errors.ErrPasswordMissingSpecialChar
	}

	v.logger.Info(ctx, "password_validator_service.ValidatePassWord.end")

	return nil
}

func ValidatePassWordService(logger *ports.Logger) ports2.ValidatePwd {
	return &validatePass{logger: *logger}

}
