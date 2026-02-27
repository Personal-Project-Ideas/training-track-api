package persistences

import (
	"context"

	"github.com/PratesJr/training-track-api/internal/application/mappers"
	ports2 "github.com/PratesJr/training-track-api/internal/common/ports"
	"github.com/PratesJr/training-track-api/internal/domain/models"
	"github.com/PratesJr/training-track-api/internal/domain/ports"
	"gorm.io/gorm"
)

type userPersistence struct {
	db     *gorm.DB
	logger ports2.Logger
}

func (u *userPersistence) Create(ctx context.Context, user models.User) (models.User, error) {
	u.logger.Info(ctx, "user_persistence.Create.call")

	entity := mappers.MapUserModelToEntity(user)
	result := u.db.Create(&entity)
	if result.Error != nil {
		u.logger.Error(ctx, "user_persistence.Create.error", result.Error)
		return models.User{}, result.Error
	}

	u.logger.Info(ctx, "user_persistence.Create.end")

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
func UserPersistence(db *gorm.DB, logger ports2.Logger) ports.UserRepository {
	return &userPersistence{db: db, logger: logger}
}
