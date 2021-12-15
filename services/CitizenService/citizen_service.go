package CitizenService

import "context"

type CitizenService interface {
	Register(ctx context.Context, citizen Citizen)(Citizen, error)
}