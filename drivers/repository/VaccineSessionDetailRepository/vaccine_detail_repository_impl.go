package VaccineSessionDetailRepository

import (
	"context"
	"errors"
	"vaccine-app-be/app/config/mysql"
	"vaccine-app-be/drivers/records"
)

type VaccineSessionDetailImpl struct {
	client mysql.Client
}

func NewSessionDetail(client mysql.Client) VaccineSessionDetail {
	return &VaccineSessionDetailImpl{client: client}
}

func (repository *VaccineSessionDetailImpl) Create(ctx context.Context, sessionId, fmId int) (records.VaccineSessionDetail, error) {
	var record records.VaccineSessionDetail
	record.SessionId = sessionId
	record.FamilyMemberId = fmId
	err := repository.client.Conn().WithContext(ctx).Create(&record).Error
	if err != nil {
		return records.VaccineSessionDetail{}, err
	}
	return record, nil
}

func (repository *VaccineSessionDetailImpl) GetDetailBySessionId(ctx context.Context, sessionId int) ([]records.VaccineSessionDetail, error) {
	var record []records.VaccineSessionDetail
	data := repository.client.Conn().WithContext(ctx).Where("session_id = ?", sessionId).Find(&record)
	if data.RowsAffected == 0 {
		return []records.VaccineSessionDetail{}, errors.New("data not found")
	}
	return record, nil
}

func (repository *VaccineSessionDetailImpl) GetDetailById(ctx context.Context, id int) (records.VaccineSessionDetail, error) {
	var record records.VaccineSessionDetail
	data := repository.client.Conn().WithContext(ctx).Where("id = ?", id).Find(&record)
	if data.RowsAffected == 0 {
		return records.VaccineSessionDetail{}, errors.New("data not found")
	}
	return record, nil
}

func (repository *VaccineSessionDetailImpl) GetDetailByFamilyId(ctx context.Context, fmid int) ([]records.VaccineSessionDetail, error) {
	var record []records.VaccineSessionDetail
	data := repository.client.Conn().WithContext(ctx).Where("family_member_id = ?", fmid).Find(&record)
	if data.RowsAffected == 0 {
		return []records.VaccineSessionDetail{}, errors.New("data not found")
	}
	return record, nil
}
