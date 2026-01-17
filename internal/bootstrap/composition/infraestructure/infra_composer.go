package infraestructure_composer

import (
	"github.com/PratesJr/training-track-api/internal/domain/ports"
	"gorm.io/gorm"
)

type InfraContainer struct {
	UserRepo ports.UserRepository
	Db       *gorm.DB
}

func Compose() *InfraContainer {

	infraContainer := &InfraContainer{}

	return infraContainer
}
