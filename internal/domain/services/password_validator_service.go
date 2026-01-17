package domain_services

import (
	"unicode"

	domain_errors "github.com/PratesJr/training-track-api/internal/domain/errors"
)

func ValidatePassword(password string) error {
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
		return domain_errors.ErrPasswordMissingUppercase
	}
	if !hasLower {
		return domain_errors.ErrPasswordMissingLowercase
	}
	if !hasSpecial {
		return domain_errors.ErrPasswordMissingSpecialChar
	}

	return nil
}
