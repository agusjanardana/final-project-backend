package VaccineSessionRepository

import (
	"context"
	"vaccine-app-be/drivers/records"
)

type VaccineSessionRepository interface {
	Create(ctx context.Context, record records.VaccineSession) (records.VaccineSession, error)
	Update(ctx context.Context, id, hfid int, record records.VaccineSession) (records.VaccineSession, error)
	Delete(ctx context.Context, id, hfid int) (records.VaccineSession, error)
	FindById(ctx context.Context, id int) (records.VaccineSession, error)
	FindSessionOwnedByHf(ctx context.Context, hfid int) ([]records.VaccineSession, error)
	GetAllVaccineSession(ctx context.Context) ([]records.VaccineSession, error)
	//GetExpiredSession(ctx context.Context) ([]records.VaccineSession, error)
	//UpdateSessionStatus(ctx context.Context, hfid, id int, status string) error
}
