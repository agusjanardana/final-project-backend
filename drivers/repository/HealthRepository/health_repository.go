package HealthRepository

import (
	"context"
	"vaccine-app-be/drivers/records"
)

type HealthRepository interface {
	Register(ctx context.Context, healthF records.HealthFacilitator) (records.HealthFacilitator, error)
	FindByEmail(ctx context.Context, email string) (records.HealthFacilitator, error)
	GetAllHealthFacilitator(ctx context.Context) ([]records.HealthFacilitator, error)
	FindById(ctx context.Context, hfId int)(records.HealthFacilitator, error)
}
