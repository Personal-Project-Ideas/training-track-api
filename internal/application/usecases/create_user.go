// Package usecases contains all use-cases for application, this file contains create user use case
package usecases

import (
	"context"

	ports2 "github.com/PratesJr/training-track-api/internal/common/ports"
	"github.com/PratesJr/training-track-api/internal/domain/models"
	"github.com/PratesJr/training-track-api/internal/domain/ports"
	domain_services "github.com/PratesJr/training-track-api/internal/domain/services"
)

type createUserUseCase struct {
	userRepository ports.UserRepository
	logger         ports2.Logger
	validatePwd    domain_services.ValidatePass
}

func (c createUserUseCase) Execute(ctx context.Context, input models.User) (error, *models.User) {
	c.logger.Info(ctx, "create_user.Execute.call", input)

	err := c.validatePwd.ValidatePassWord(ctx, *input.Password)

	if err != nil {
		return err, nil
	}

	c.logger.Info(ctx, "create_user.Execute.end")

	//TODO implement me
	panic("implement me")
}

// CreateUserUseCase creates a new instance of create user use case
func CreateUserUseCase(userRepository *ports.UserRepository, logger *ports2.Logger, pwdValidator *domain_services.ValidatePass) ports.CreateUserUseCase {
	return &createUserUseCase{
		userRepository: *userRepository,
		logger:         *logger,
		validatePwd:    *pwdValidator,
	}
}
