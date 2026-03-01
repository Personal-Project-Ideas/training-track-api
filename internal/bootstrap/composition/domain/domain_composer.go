// Package domain_composer provides composition for domain layer dependencies.
package domain_composer

import (
	"github.com/PratesJr/training-track-api/internal/common/ports"
	ports2 "github.com/PratesJr/training-track-api/internal/domain/ports"
	domain_services "github.com/PratesJr/training-track-api/internal/domain/services"
)

// DomainContainer holds domain layer dependencies.
type DomainContainer struct {
	ValidatePass ports2.ValidatePwd
}

// Compose initializes and returns the DomainContainer.
func Compose(
	logger ports.Logger,
) *DomainContainer {
	var pwdValidator = domain_services.ValidatePassWordService(&logger)

	return &DomainContainer{
		ValidatePass: pwdValidator,
	}
}
