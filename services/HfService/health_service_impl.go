package HfService

import (
	"context"
	"vaccine-app-be/app/middleware"
	"vaccine-app-be/drivers/repository/HealthRepository"
)

type HealthServiceImpl struct {
	HealthRepository HealthRepository.HealthRepository
	jwtAuth          *middleware.ConfigJWT
}

func NewHealthService(HealthRepository HealthRepository.HealthRepository, jwtAuth *middleware.ConfigJWT) HealthService {
	return &HealthServiceImpl{
		HealthRepository: HealthRepository,
		jwtAuth:          jwtAuth,
	}
}

func (service *HealthServiceImpl) Register(ctx context.Context, healthF HealthFacilitator) (HealthFacilitator, error) {
	panic("implement me")
}

func (service *HealthServiceImpl) Login(ctx context.Context, email, password string) (string, error) {
	panic("implement me")
}
