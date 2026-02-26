package domain_composer

import (
	"github.com/PratesJr/training-track-api/internal/common/ports"
	domain_services "github.com/PratesJr/training-track-api/internal/domain/services"
)

type DomainContainer struct {
	domain_services.ValidatePass
}

func Compose(
	logger ports.Logger,
) *DomainContainer {
	var pwdValidator *domain_services.ValidatePass

	pwdValidator = domain_services.ValidatePassWordService(&logger)
	return &DomainContainer{
		ValidatePass: *pwdValidator,
	}
}
