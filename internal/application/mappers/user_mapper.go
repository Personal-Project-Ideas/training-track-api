package mappers

import (
	"github.com/PratesJr/training-track-api/internal/application/dtos"
	"github.com/PratesJr/training-track-api/internal/domain/models"
)

// MapCreateInputToUser converts a UserInputDTO to a User model.
func MapCreateInputToUser(dto dtos.UserInputDTO) *models.User {
	return &models.User{
		FullName: dto.FullName,
		Age:      dto.Age,
		Email:    dto.Email,
		Password: dto.Password,
	}
}

// MapUpdateInputToUser converts a UserUpdateInputDTO to a User model.
func MapUpdateInputToUser(dto dtos.UserUpdateInputDTO, user models.User) *models.User {
	if dto.FullName != nil {
		user.FullName = *dto.FullName
	}
	if dto.Age != nil {
		user.Age = *dto.Age
	}
	if dto.Email != nil {
		user.Email = *dto.Email
	}
	if dto.Password != nil {
		user.Password = *dto.Password
	}
	return &user
}

// MapUserToOutputDTO converts a User model to a UserOutputDTO.
func MapUserToOutputDTO(user models.User) *dtos.UserOutputDTO {
	return &dtos.UserOutputDTO{
		ID:        user.ID.String(),
		FullName:  user.FullName,
		Age:       user.Age,
		Email:     user.Email,
		CreatedAt: *user.CreatedAt,
		UpdatedAt: *user.UpdatedAt,
	}
}
