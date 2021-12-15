package CitizenService

import (
	"context"
	"vaccine-app-be/drivers/repository/CitizenRepository"
)

type CitizenServiceImpl struct{
	CitizenRepository CitizenRepository.CitizenRepository
}

func NewCitizenService(CitizenRepository CitizenRepository.CitizenRepository) CitizenService {
	return &CitizenServiceImpl{
		CitizenRepository:  CitizenRepository,
	}
}

func (service *CitizenServiceImpl) Register(ctx context.Context, citizen Citizen) (Citizen, error) {
	panic("implement me")
}
