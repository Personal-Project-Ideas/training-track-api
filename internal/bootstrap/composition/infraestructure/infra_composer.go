package infraestructure_composer

import (
	ports "github.com/PratesJr/training-track-api/internal/common/ports"
	ports2 "github.com/PratesJr/training-track-api/internal/domain/ports"
	"github.com/PratesJr/training-track-api/internal/infra/persistences"
	"github.com/PratesJr/training-track-api/internal/pkg/database"
	"gorm.io/gorm"
)

type InfraContainer struct {
	UserRepo ports2.UserRepository
}

func Compose(logger ports.Logger, testDB ...*gorm.DB) *InfraContainer {
	var db *gorm.DB
	if len(testDB) > 0 && testDB[0] != nil {
		db = testDB[0]
	} else {
		db = database.DbConnection()
	}
	userRepo := persistences.UserPersistence(db, logger)
	infraContainer := &InfraContainer{
		UserRepo: userRepo,
	}
	return infraContainer
}
