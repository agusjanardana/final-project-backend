package VaccineSessionRepository

import (
	"context"
	"errors"
	"vaccine-app-be/app/config/mysql"
	"vaccine-app-be/drivers/records"
)

type VaccineSessionRepositoryImpl struct {
	client mysql.Client
}

func NewVaccineSessionRepository(client mysql.Client) VaccineSessionRepository {
	return &VaccineSessionRepositoryImpl{client: client}
}

func (repository *VaccineSessionRepositoryImpl) Create(ctx context.Context, record records.VaccineSession) (records.VaccineSession, error) {
	err := repository.client.Conn().WithContext(ctx).Create(&record).Error
	if err != nil {
		return records.VaccineSession{}, err
	}
	return record, nil
}

func (repository *VaccineSessionRepositoryImpl) Update(ctx context.Context, id, hfid int, record records.VaccineSession) (records.VaccineSession, error) {
	err := repository.client.Conn().WithContext(ctx).Where("id = ? AND health_facilitator_id = ?", id, hfid).Updates(&record).Error
	if err != nil {
		return records.VaccineSession{}, err
	}
	return record, nil
}

func (repository *VaccineSessionRepositoryImpl) Delete(ctx context.Context, id, hfid int) (records.VaccineSession, error) {
	var record records.VaccineSession
	err := repository.client.Conn().WithContext(ctx).Where("id = ? AND health_facilitator_id = ?", id, hfid).Delete(&record).Error
	if err != nil {
		return records.VaccineSession{}, err
	}
	return record, nil
}

func (repository *VaccineSessionRepositoryImpl) FindById(ctx context.Context, id int) (records.VaccineSession, error) {
	var record records.VaccineSession
	data := repository.client.Conn().WithContext(ctx).Preload("VaccineSessionDetail").Where("id = ?", id).Find(&record)
	if data.RowsAffected == 0 {
		return records.VaccineSession{}, errors.New("data not found")
	}
	return record, nil
}

func (repository *VaccineSessionRepositoryImpl) FindSessionOwnedByHf(ctx context.Context, hfid int) ([]records.VaccineSession, error) {
	var record []records.VaccineSession
	data := repository.client.Conn().WithContext(ctx).Where("health_facilitator_id = ?", hfid).Find(&record)
	if data.RowsAffected == 0 {
		return []records.VaccineSession{}, errors.New("data not found")
	}
	return record, nil
}

func (repository *VaccineSessionRepositoryImpl) GetAllVaccineSession(ctx context.Context) ([]records.VaccineSession, error) {
	var record []records.VaccineSession
	data := repository.client.Conn().WithContext(ctx).Find(&record)
	if data.RowsAffected == 0 {
		return []records.VaccineSession{}, errors.New("data not found")
	}
	return record, nil
}
