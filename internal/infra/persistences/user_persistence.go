package persistences

import (
	"context"
	"github.com/PratesJr/training-track-api/internal/application/mappers"
	"github.com/PratesJr/training-track-api/internal/domain/models"
	"github.com/PratesJr/training-track-api/internal/domain/ports"
	"gorm.io/gorm"
)

type userPersistence struct {
	db *gorm.DB
}

func (u *userPersistence) Create(ctx context.Context, user models.User) (models.User, error) {
	entity := mappers.MapUserModelToEntity(user)
	result := u.db.Create(&entity)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return mappers.MapUserEntityToModel(entity), nil
}

func (u *userPersistence) Update(ctx context.Context, user models.User) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userPersistence) FindByID(ctx context.Context, id string) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userPersistence) FindByEmail(ctx context.Context, email string) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userPersistence) List(ctx context.Context) ([]models.User, error) {
	//TODO implement me
	panic("implement me")
}

// UserPersistence creates a new instance of user persistence
func UserPersistence(db *gorm.DB) ports.UserRepository {
	return &userPersistence{db: db}
}
