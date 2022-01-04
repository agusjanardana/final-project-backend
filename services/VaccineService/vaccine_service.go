package VaccineService

import (
	"context"
)

type VaccineService interface {
	Create(ctx context.Context, vaccine Vaccine) (Vaccine, error)
	Update(ctx context.Context, hfid, vaccineId int, vaccine Vaccine) (Vaccine, error)
	Delete(ctx context.Context, hfid, vaccineId int) (string, error)
	FindVaccineById(ctx context.Context, vaccineId int) (Vaccine, error)
	FindVaccineOwnedByHF(ctx context.Context, hfid int) ([]Vaccine, error)

}