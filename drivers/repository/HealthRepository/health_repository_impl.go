package HealthRepository

import (
	"context"
	"errors"
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
	err := repository.client.Conn().WithContext(ctx).Preload("Vaccine.VaccineSession").Preload("VaccineSession.VaccineSessionDetail").Preload("VaccineSession").Find(&record).Debug().Error
	if err != nil {
		return record, err
	}
	return record, nil
}

func (repository *HealthRepositoryImpl) FindById(ctx context.Context, hfId int) (records.HealthFacilitator, error) {
	healthF := records.HealthFacilitator{}

	data := repository.client.Conn().WithContext(ctx).Preload("Vaccine.VaccineSession").Preload("VaccineSession.VaccineSessionDetail").Preload("VaccineSession").Where("id = ?", hfId).Find(&healthF)
	if data.RowsAffected == 0 {
		return records.HealthFacilitator{}, errors.New("data not found")
	}
	return healthF, nil
}

func (repository *HealthRepositoryImpl) Update(ctx context.Context, hfId int, record records.HealthFacilitator) (records.HealthFacilitator, error) {
	err := repository.client.Conn().WithContext(ctx).Where("id = ?", hfId).Updates(&record).Error
	if err != nil {
		return records.HealthFacilitator{}, err
	}
	return record, nil
}
