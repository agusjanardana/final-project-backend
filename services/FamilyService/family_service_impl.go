package FamilyService

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"log"
	"vaccine-app-be/drivers/records"
	"vaccine-app-be/drivers/repository/FamilyRepository"
)

type FamilyServiceImpl struct {
	FamilyRepository FamilyRepository.FamilyRepository
}

func NewFamilyService(FamilyRepository FamilyRepository.FamilyRepository) FamilyService {
	return &FamilyServiceImpl{
		FamilyRepository: FamilyRepository,
	}
}

func (service *FamilyServiceImpl) Create(ctx context.Context, id int, family FamilyMember) (FamilyMember, error) {
	err := family.Validation()
	if err != nil {
		return FamilyMember{}, err
	}

	if len(string(rune(id))) == 0 {
		return FamilyMember{}, errors.New("user id cannot be blank")
	}
	//set id
	family.CitizenId = id

	entityFamily := new(records.FamilyMember)
	copier.Copy(entityFamily, &family)
	create, err := service.FamilyRepository.Create(ctx, *entityFamily)
	if err != nil {
		return FamilyMember{}, err
	}
	respond := FamilyMember{}
	copier.Copy(&respond, &create)

	return respond, nil
}

func (service *FamilyServiceImpl) GetFamilyById(ctx context.Context, id int) (FamilyMember, error) {
	byId, err := service.FamilyRepository.GetFamilyById(ctx, id)
	if err != nil {
		return FamilyMember{}, err
	}
	respond := FamilyMember{}
	copier.Copy(&respond, &byId)

	return respond, nil

}

func (service *FamilyServiceImpl) GetCitizenOwnFamily(ctx context.Context, citizenId int) ([]FamilyMember, error) {
	if len(string(rune(citizenId))) == 0 {
		return nil, errors.New("citizen id cannot be blank")
	}

	family, err := service.FamilyRepository.GetCitizenOwnFamily(ctx, citizenId)
	if err != nil {
		return nil, err
	}
	var response []FamilyMember
	copier.Copy(&response, &family)
	return response, nil
}

func (service *FamilyServiceImpl) Update(ctx context.Context, id int, family FamilyMember) (FamilyMember, error) {
	if len(string(rune(id))) == 0 {
		return family, errors.New("id cannot be blank")
	}

	_, err := service.FamilyRepository.GetFamilyById(ctx, id)
	if err != nil {
		return FamilyMember{}, err
	}

	request := new(records.FamilyMember)
	copier.Copy(request, &family)
	log.Println(request)
	update, err := service.FamilyRepository.Update(ctx, id, *request)
	if err != nil {
		return FamilyMember{}, err
	}

	respond := FamilyMember{}
	copier.Copy(&respond, &update)

	return respond, nil
}

func (service *FamilyServiceImpl) Delete(ctx context.Context, id int, citizenId int) (string, error) {
	if len(string(rune(id))) == 0 || len(string(rune(citizenId))) == 0 {
		return "", errors.New("field cannot be blank")
	}

	byId, err := service.FamilyRepository.GetFamilyById(ctx, id)
	if err != nil {
		return "", err
	}

	if byId.CitizenId == citizenId {
		_ , err := service.FamilyRepository.Delete(ctx, id,citizenId)
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("this family is not belongs to you")
	}
	return "success delete data", nil
}
