package CitizenRepository

import (
	"context"
	"errors"
	"time"
	"vaccine-app-be/app/config/mysql"
	"vaccine-app-be/drivers/records"
)

type CitizenRepositoryImpl struct {
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

func (repository *CitizenRepositoryImpl) FindByEmail(ctx context.Context, email string) (records.Citizen, error) {
	citizen := records.Citizen{}
	err := repository.client.Conn().WithContext(ctx).Where("email = ?", email).First(&citizen).Error
	if err != nil {
		return citizen, err
	}
	return citizen, nil
}

func (repository *CitizenRepositoryImpl) Update(ctx context.Context, userId int, birthDay time.Time, address string) (records.Citizen, error) {
	citizen := records.Citizen{}
	err := citizen.Birthday.Scan(birthDay)
	if err != nil {
		return records.Citizen{}, err
	}
	citizen.Address = address
	err = repository.client.Conn().WithContext(ctx).Where("id = ?", userId).Updates(&citizen).Error
	if err != nil {
		return records.Citizen{}, err
	}
	return citizen, nil
}

func (repository *CitizenRepositoryImpl) FindById(ctx context.Context, citizenId int) (records.Citizen, error) {
	citizen := records.Citizen{}
	data := repository.client.Conn().WithContext(ctx).Where("id = ?", citizenId).Debug().Find(&citizen)
	if data.RowsAffected == 0 {
		return citizen, errors.New("data not found")
	}
	return citizen, nil
}
