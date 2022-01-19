package CitizenService

import (
	"context"
	"time"
)

type CitizenService interface {
	Register(ctx context.Context, citizen Citizen) (Citizen, error)
	Login(ctx context.Context, email, password string) (string, error)
	Update(ctx context.Context, userId int, birthDay time.Time, address string) (Citizen, error)
	CitizenFindById(ctx context.Context, userId int) (Citizen, error)
	GetCitizenRelationWithHealthFacilitators(ctx context.Context, hfId int)([]Citizen, error)
}
