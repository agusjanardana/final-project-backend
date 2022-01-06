package HfService

import "context"

type HealthService interface {
	Register(ctx context.Context, healthF HealthFacilitator) (HealthFacilitator, error)
	Login(ctx context.Context, email, password string) (string, error)
	GetAllHealthFacilitator(ctx context.Context) ([]HealthFacilitator, error)
}
