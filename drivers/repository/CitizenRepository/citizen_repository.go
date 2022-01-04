package CitizenRepository

import (
	"context"
	"time"
	"vaccine-app-be/drivers/records"
)

type CitizenRepository interface {
	Register(ctx context.Context, citizens records.Citizen) (records.Citizen, error)
	FindByEmail(ctx context.Context, email string) (records.Citizen, error)
	Update(ctx context.Context, userId int, birthDay time.Time, address string) (records.Citizen, error)
}
