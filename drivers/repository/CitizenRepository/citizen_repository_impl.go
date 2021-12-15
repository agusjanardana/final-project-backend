package CitizenRepository

import (
	"context"
	"vaccine-app-be/app/config/mysql"
	"vaccine-app-be/drivers/records"
)

type CitizenRepositoryImpl struct{
	client mysql.Client
}

func NewCitizenRepository(client mysql.Client) CitizenRepository {
	return &CitizenRepositoryImpl{client}

}

func (repository *CitizenRepositoryImpl) Register(ctx context.Context, citizens records.Citizen) (records.Citizen, error) {
	err := repository.client.Conn().WithContext(ctx).Create(&citizens).Error
	if err != nil {
		return citizens, err
	}
	return citizens, nil
}
