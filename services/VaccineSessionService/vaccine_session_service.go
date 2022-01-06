package VaccineSessionService

import "context"

type VaccineSessionService interface {
	CreateSession(ctx context.Context, domain VaccineSession) (VaccineSession, error)
	GetSessionById(ctx context.Context, id int) (VaccineSession, error)
	GetSessionOwnedByHf(ctx context.Context, hfid int) ([]VaccineSession, error)
	DeleteSession(ctx context.Context, hfid, id int) (string, error)
	UpdateSession(ctx context.Context, hfid, id int, domain VaccineSession) (VaccineSession, error)
	GetAllVaccineSession(ctx context.Context)([]VaccineSession, error)
}
