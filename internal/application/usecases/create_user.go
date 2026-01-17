// Package usecases contains all use-cases for application, this file contains create user use case
package usecases

import (
	"github.com/PratesJr/training-track-api/internal/domain/models"
	"github.com/PratesJr/training-track-api/internal/domain/ports"
	domain_services "github.com/PratesJr/training-track-api/internal/domain/services"
)

type createUserUseCase struct {
	userRepository *ports.UserRepository
}

func (c createUserUseCase) Execute(input models.User) (*models.User, error) {
	err := domain_services.ValidatePassword(*input.Password)

	if err != nil {
		return nil, err
	}
	//TODO implement me
	panic("implement me")
}

// CreateUserUseCase creates a new instance of create user use case
func CreateUserUseCase(userRepository ports.UserRepository) ports.CreateUserUseCase {
	return &createUserUseCase{
		userRepository: &userRepository,
	}
}
