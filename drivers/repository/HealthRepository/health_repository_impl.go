package HealthRepository

import (
	"context"
	"vaccine-app-be/app/config/mysql"
	"vaccine-app-be/drivers/records"
)

type HealthRepositoryImpl struct{
	client mysql.Client
}

func NewHealthRepository(client mysql.Client) HealthRepository {
	return &HealthRepositoryImpl{client}
}

func (repository *HealthRepositoryImpl) Register(ctx context.Context, healthF records.HealthFacilitator) (records.HealthFacilitator, error) {
	err := repository.client.Conn().WithContext(ctx).Create(&healthF).Debug().Error
	if err != nil{
		return healthF, err
	}
	return healthF, nil
}

func (repository *HealthRepositoryImpl) FindByEmail(ctx context.Context, email string) (records.HealthFacilitator, error) {
	healthF := records.HealthFacilitator{}
	err := repository.client.Conn().WithContext(ctx).Where("email = ?", email).Debug().First(&healthF).Error
	if err != nil{
		return healthF, err
	}
	return healthF, nil
}

