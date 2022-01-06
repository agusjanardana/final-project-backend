package HealthRepository

import (
	"context"
	"gorm.io/gorm/clause"
	"vaccine-app-be/app/config/mysql"
	"vaccine-app-be/drivers/records"
)

type HealthRepositoryImpl struct {
	client mysql.Client
}

func NewHealthRepository(client mysql.Client) HealthRepository {
	return &HealthRepositoryImpl{client}
}

func (repository *HealthRepositoryImpl) Register(ctx context.Context, healthF records.HealthFacilitator) (records.HealthFacilitator, error) {
	err := repository.client.Conn().WithContext(ctx).Create(&healthF).Debug().Error
	if err != nil {
		return healthF, err
	}
	return healthF, nil
}

func (repository *HealthRepositoryImpl) FindByEmail(ctx context.Context, email string) (records.HealthFacilitator, error) {
	healthF := records.HealthFacilitator{}
	err := repository.client.Conn().WithContext(ctx).Where("email = ?", email).Debug().First(&healthF).Error
	if err != nil {
		return healthF, err
	}
	return healthF, nil
}

func (repository *HealthRepositoryImpl) GetAllHealthFacilitator(ctx context.Context) ([]records.HealthFacilitator, error) {
	var record []records.HealthFacilitator
	err := repository.client.Conn().WithContext(ctx).Preload(clause.Associations).Find(&record).Debug().Error
	if err != nil {
		return record, err
	}
	return record, nil
}
