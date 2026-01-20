package infraestructure_composer

import (
	"github.com/PratesJr/training-track-api/internal/domain/ports"
	"github.com/PratesJr/training-track-api/internal/infra/persistences"
	"github.com/PratesJr/training-track-api/internal/pkg/database"
	"gorm.io/gorm"
)

type InfraContainer struct {
	UserRepo ports.UserRepository
}

func Compose() *InfraContainer {

	var db *gorm.DB
	var userRepo ports.UserRepository

	db = database.DbConnection()
	userRepo = persistences.UserPersistence(db)

	infraContainer := &InfraContainer{
		UserRepo: userRepo,
	}

	return infraContainer
}
