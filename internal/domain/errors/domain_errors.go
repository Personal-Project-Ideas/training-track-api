package domain_errors

import "errors"

var (
	ErrPasswordTooShort           = errors.New("password must be at least 8 characters long")
	ErrPasswordMissingUppercase   = errors.New("password must contain at least one uppercase letter")
	ErrPasswordMissingLowercase   = errors.New("password must contain at least one lowercase letter")
	ErrPasswordMissingSpecialChar = errors.New("password must contain at least one special character")
)
