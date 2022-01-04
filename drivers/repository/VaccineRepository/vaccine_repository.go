package VaccineRepository

import (
	"context"
	"vaccine-app-be/drivers/records"
)

type VaccineRepository interface {
	Create(ctx context.Context, vaccine records.Vaccine) (records.Vaccine, error)
	Update(ctx context.Context, hfid, vaccineId int, vaccine records.Vaccine) (records.Vaccine, error)
	Delete(ctx context.Context, hfid, vaccineId int) (records.Vaccine, error)
	FindVaccineById(ctx context.Context, vaccineId int) (records.Vaccine, error)
	FindVaccineOwnedByHF(ctx context.Context, hfid int) ([]records.Vaccine, error)
}
