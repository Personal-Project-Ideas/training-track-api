package domain_composer

import (
	"github.com/PratesJr/training-track-api/internal/common/ports"
	ports2 "github.com/PratesJr/training-track-api/internal/domain/ports"
	domain_services "github.com/PratesJr/training-track-api/internal/domain/services"
)

type DomainContainer struct {
	ValidatePass ports2.ValidatePwd
}

func Compose(
	logger ports.Logger,
) *DomainContainer {
	var pwdValidator ports2.ValidatePwd

	pwdValidator = domain_services.ValidatePassWordService(&logger)
	return &DomainContainer{
		ValidatePass: pwdValidator,
	}
}
