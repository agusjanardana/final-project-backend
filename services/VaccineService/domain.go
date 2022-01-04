package VaccineService

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Vaccine struct {
	Id                  int
	HealthFacilitatorId int
	Name                string
	Stock               int
}

type VaccineUpdate struct {
	Name  string
	Stock int
}

func (vc Vaccine) ValidateRequest() error {
	return validation.ValidateStruct(&vc,
		//		validasi request email
		validation.Field(&vc.HealthFacilitatorId, validation.Required),
		//		validasi Password
		validation.Field(&vc.Name, validation.Required),
		validation.Field(&vc.Stock, validation.Required),
	)
}
