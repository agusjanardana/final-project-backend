package VaccineRepository

import (
	"context"
	"errors"
	"vaccine-app-be/app/config/mysql"
	"vaccine-app-be/drivers/records"
)

type VaccineRepositoryImpl struct {
	client mysql.Client
}

func NewVaccineRepository(client mysql.Client) VaccineRepository {
	return &VaccineRepositoryImpl{client}
}

func (repository *VaccineRepositoryImpl) Create(ctx context.Context, vaccine records.Vaccine) (records.Vaccine, error) {
	err := repository.client.Conn().WithContext(ctx).Create(&vaccine).Error
	if err != nil {
		return records.Vaccine{}, err
	}
	return vaccine, nil
}

func (repository *VaccineRepositoryImpl) Update(ctx context.Context, hfid, vaccineId int, vaccine records.Vaccine) (records.Vaccine, error) {
	err := repository.client.Conn().WithContext(ctx).Where("id = ? AND health_facilitator_id = ?", vaccineId, hfid).Updates(&vaccine).Error
	if err != nil {
		return records.Vaccine{}, err
	}
	return vaccine, nil
}

func (repository *VaccineRepositoryImpl) Delete(ctx context.Context, hfid, vaccineId int) (records.Vaccine, error) {
	var vaccine records.Vaccine
	err := repository.client.Conn().WithContext(ctx).Where("id = ? AND health_facilitator_id = ?", vaccineId, hfid).Delete(&vaccine).Error
	if err != nil {
		return records.Vaccine{}, err
	}
	return vaccine, nil
}

func (repository *VaccineRepositoryImpl) FindVaccineById(ctx context.Context, vaccineId int) (records.Vaccine, error) {
	var vaccine records.Vaccine
	data := repository.client.Conn().WithContext(ctx).Where("id = ?", vaccineId).Find(&vaccine)
	if data.RowsAffected == 0 {
		return vaccine, errors.New("data not found")
	}
	return vaccine, nil
}

func (repository *VaccineRepositoryImpl) FindVaccineOwnedByHF(ctx context.Context, hfid int) ([]records.Vaccine, error) {
	var vaccine []records.Vaccine
	data := repository.client.Conn().WithContext(ctx).Where("health_facilitator_id = ?", hfid).Find(&vaccine)
	if data.RowsAffected == 0 {
		return vaccine, errors.New("data not found")
	}
	return vaccine, nil
}
