package VaccineService

import (
	"context"
	"github.com/jinzhu/copier"
	"vaccine-app-be/drivers/records"
	"vaccine-app-be/drivers/repository/VaccineRepository"
)

type VaccineServiceImpl struct {
	vaccineRepository VaccineRepository.VaccineRepository
}

func NewVaccineRepository(vaccineRepository VaccineRepository.VaccineRepository) VaccineService {
	return &VaccineServiceImpl{vaccineRepository: vaccineRepository}
}

func (service *VaccineServiceImpl) Create(ctx context.Context, vaccine Vaccine) (Vaccine, error) {
	err := vaccine.ValidateRequest()
	if err != nil {
		return Vaccine{}, err
	}
	entityVaccine := new(records.Vaccine)
	copier.Copy(entityVaccine, &vaccine)
	create, err := service.vaccineRepository.Create(ctx, *entityVaccine)
	if err != nil {
		return Vaccine{}, err
	}

	response := Vaccine{}
	copier.Copy(&response, &create)

	return response, nil
}

func (service *VaccineServiceImpl) Update(ctx context.Context, hfid, vaccineId int, vaccine Vaccine) (Vaccine, error) {
	_, err := service.FindVaccineById(ctx, vaccineId)
	if err != nil {
		return Vaccine{}, err
	}
	entityVaccine := new(records.Vaccine)
	copier.Copy(&entityVaccine, &vaccine)
	update, err := service.vaccineRepository.Update(ctx, hfid, vaccineId, *entityVaccine)
	if err != nil {
		return Vaccine{}, err
	}
	response := Vaccine{}
	copier.Copy(&response, &update)
	return response, nil
}

func (service *VaccineServiceImpl) Delete(ctx context.Context, hfid, vaccineId int) (string, error) {
	_, err := service.FindVaccineById(ctx, vaccineId)
	if err != nil {
		return "", err
	}

	_, err = service.vaccineRepository.Delete(ctx, hfid, vaccineId)
	if err != nil {
		return "", err
	}
	return "success delete data", nil
}

func (service *VaccineServiceImpl) FindVaccineById(ctx context.Context, vaccineId int) (Vaccine, error) {
	id, err := service.vaccineRepository.FindVaccineById(ctx, vaccineId)
	if err != nil {
		return Vaccine{}, err
	}
	response := Vaccine{}
	copier.Copy(&response, &id)
	return response, nil
}

func (service *VaccineServiceImpl) FindVaccineOwnedByHF(ctx context.Context, hfid int) ([]Vaccine, error) {
	hf, err := service.vaccineRepository.FindVaccineOwnedByHF(ctx, hfid)
	if err != nil {
		return nil, err
	}
	var response []Vaccine
	copier.Copy(&response, &hf)

	return response, nil
}
