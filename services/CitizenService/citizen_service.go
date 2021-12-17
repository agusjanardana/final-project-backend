package CitizenService

import "context"

type CitizenService interface {
	Register(ctx context.Context, citizen Citizen)(Citizen, error)
	Login(ctx context.Context, email, password string) (string, error)
}