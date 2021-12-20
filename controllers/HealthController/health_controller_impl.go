package HealthController

import (
	"context"
	"vaccine-app-be/drivers/records"
	"vaccine-app-be/services/HfService"
)

type HealthFacilitatorImpl struct{
	healthService HfService.HealthService
}

func (controller *HealthFacilitatorImpl) Register(ctx context.Context, healthF records.HealthFacilitator) (records.HealthFacilitator, error) {
	panic("implement me")
}

func (controller *HealthFacilitatorImpl) FindByEmail(ctx context.Context, email string) (records.HealthFacilitator, error) {
	panic("implement me")
}
