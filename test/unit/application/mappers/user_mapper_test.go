// Package mappers_test contains tests for all mappers in the mappers package.

package mappers_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/PratesJr/training-track-api/internal/application/dtos"
	"github.com/PratesJr/training-track-api/internal/application/mappers"
	"github.com/PratesJr/training-track-api/internal/domain/models"
	"github.com/PratesJr/training-track-api/internal/infra/entities"
)

func TestMapCreateInputToUser(t *testing.T) {
	t.Run("it should map create input dto to user model", func(t *testing.T) {
		dto := dtos.UserInputDTO{
			FullName: "John Doe",
			Age:      30,
			Email:    "john@test.com",
			Password: "password123",
		}

		user := mappers.MapCreateInputToUser(dto)

		assert.Equal(t, dto.FullName, user.FullName)
		assert.Equal(t, dto.Age, user.Age)
		assert.Equal(t, dto.Email, user.Email)
		assert.NotNil(t, user.Password)
		assert.Equal(t, dto.Password, *user.Password)
	})
}

func TestMapUpdateInputToUser(t *testing.T) {
	t.Run("it should update only provided fields", func(t *testing.T) {
		fullName := "Updated Name"
		age := 40

		existing := models.User{
			FullName: "Old Name",
			Age:      20,
			Email:    "old@test.com",
		}

		dto := dtos.UserUpdateInputDTO{
			FullName: &fullName,
			Age:      &age,
		}

		user := mappers.MapUpdateInputToUser(dto, existing)

		assert.Equal(t, fullName, user.FullName)
		assert.Equal(t, age, user.Age)
		assert.Equal(t, "old@test.com", user.Email)
	})

	t.Run("it should not change fields when dto is empty", func(t *testing.T) {
		existing := models.User{
			FullName: "Same Name",
			Age:      25,
			Email:    "same@test.com",
		}

		dto := dtos.UserUpdateInputDTO{}

		user := mappers.MapUpdateInputToUser(dto, existing)

		assert.Equal(t, existing.FullName, user.FullName)
		assert.Equal(t, existing.Age, user.Age)
		assert.Equal(t, existing.Email, user.Email)
	})
}

func TestMapUserToOutputDTO(t *testing.T) {
	t.Run("it should map user model to output dto", func(t *testing.T) {
		id := uuid.New()
		createdAt := time.Now()
		updatedAt := time.Now().Add(time.Hour)

		user := models.User{
			ID:        &id,
			FullName:  "John Doe",
			Age:       30,
			Email:     "john@test.com",
			CreatedAt: &createdAt,
			UpdatedAt: &updatedAt,
		}

		dto := mappers.MapUserToOutputDTO(user)

		assert.Equal(t, id.String(), dto.ID)
		assert.Equal(t, user.FullName, dto.FullName)
		assert.Equal(t, user.Age, dto.Age)
		assert.Equal(t, user.Email, dto.Email)
		assert.Equal(t, createdAt.String(), dto.CreatedAt)
		assert.Equal(t, updatedAt.String(), dto.UpdatedAt)
	})
}

func TestMapUserEntityToModel(t *testing.T) {
	t.Run("it should map user entity to domain model", func(t *testing.T) {
		id := uuid.New()
		createdAt := time.Now()
		updatedAt := time.Now().Add(time.Hour)

		entity := entities.UserEntity{
			ID:        id,
			FullName:  "John Doe",
			Age:       30,
			Email:     "john@test.com",
			Password:  "hashed-password",
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}

		user := mappers.MapUserEntityToModel(entity)

		assert.Equal(t, id, *user.ID)
		assert.Equal(t, entity.FullName, user.FullName)
		assert.Equal(t, entity.Age, user.Age)
		assert.Equal(t, entity.Email, user.Email)
		assert.NotNil(t, user.Password)
		assert.Equal(t, entity.Password, *user.Password)
		assert.Equal(t, createdAt, *user.CreatedAt)
		assert.Equal(t, updatedAt, *user.UpdatedAt)
	})
}

func TestMapUserModelToEntity(t *testing.T) {
	t.Run("it should map user model to entity", func(t *testing.T) {
		password := "hashed-password"

		user := models.User{
			FullName: "John Doe",
			Age:      30,
			Email:    "john@test.com",
			Password: &password,
		}

		entity := mappers.MapUserModelToEntity(user)

		assert.Equal(t, user.FullName, entity.FullName)
		assert.Equal(t, user.Age, entity.Age)
		assert.Equal(t, user.Email, entity.Email)
		assert.Equal(t, password, entity.Password)
	})
}
