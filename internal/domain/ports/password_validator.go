package ports

import "context"

type ValidatePwd interface {
	ValidatePassWord(ctx context.Context, pwd string) error
}
