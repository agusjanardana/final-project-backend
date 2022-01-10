package FamilyRepository

import (
	"context"
	"errors"
	"vaccine-app-be/app/config/mysql"
	"vaccine-app-be/drivers/records"
)

type FamilyRepositoryImpl struct {
	client mysql.Client
}

func NewFamilyRepository(client mysql.Client) FamilyRepository {
	return &FamilyRepositoryImpl{client}
}

func (repository *FamilyRepositoryImpl) Create(ctx context.Context, family records.FamilyMember) (records.FamilyMember, error) {
	err := repository.client.Conn().Debug().WithContext(ctx).Create(&family).Error
	if err != nil {
		return records.FamilyMember{}, err
	}
	return family, nil
}

func (repository *FamilyRepositoryImpl) GetFamilyById(ctx context.Context, id int) (records.FamilyMember, error) {
	var familyMember records.FamilyMember
	data := repository.client.Conn().Debug().WithContext(ctx).Where("id = ?", id).Find(&familyMember)
	if data.RowsAffected == 0 {
		return familyMember, errors.New("data not found")
	}
	return familyMember, nil
}

func (repository *FamilyRepositoryImpl) GetCitizenOwnFamily(ctx context.Context, citizenId int) ([]records.FamilyMember, error) {
	var familyRecord []records.FamilyMember
	data := repository.client.Conn().Debug().WithContext(ctx).Preload("VaccineSessionDetail").Where("citizen_id = ?", citizenId).Find(&familyRecord)
	if data.RowsAffected == 0 {
		return familyRecord, errors.New("data not found")
	}
	return familyRecord, nil
}

func (repository *FamilyRepositoryImpl) Update(ctx context.Context, id int, family records.FamilyMember) (records.FamilyMember, error) {
	err := repository.client.Conn().Debug().WithContext(ctx).Where("id = ?", id).Updates(&family).Error
	if err != nil {
		return family, err
	}
	return family, nil
}

func (repository *FamilyRepositoryImpl) Delete(ctx context.Context, id int, citizenId int) (records.FamilyMember, error) {
	var familyMember records.FamilyMember

	err := repository.client.Conn().Debug().WithContext(ctx).Where("id = ? AND citizen_id = ? ", id, citizenId).Delete(&familyMember).Error
	if err != nil {
		return familyMember, err
	}

	return familyMember, nil
}
