package HfService

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type HealthFacilitator struct {
	Id                        int
	Name                      string
	Email                     string
	Password                  string
	FacilitatorsCertification string
	Address                   string
	Longitude                 string
	Latitude                  string
	Type                      string
}

func (hf HealthFacilitator) ValidateRequest() error {
	return validation.ValidateStruct(&hf,
		//		validasi request email
		validation.Field(&hf.Email, validation.Required, is.Email),
		//		validasi Password
		validation.Field(&hf.Password, validation.Required),
	)
}

func (hf HealthFacilitator) Validate() error {
	return validation.ValidateStruct(&hf,
		//validasi Nama tidak boleh kosong
		validation.Field(&hf.Name, validation.Required),
		//validasi Email no kosong, and regex
		validation.Field(&hf.Email, validation.Required, is.Email),
		//validasi NIK tidak boleh kosong
		validation.Field(&hf.Address, validation.Required),
		//validation password no kosong
		validation.Field(&hf.Password, validation.Required),
		//validation longitude no kosong
		validation.Field(&hf.Longitude, validation.Required),
		//validation Latitude no kosong
		validation.Field(&hf.Latitude, validation.Required),
	)
}
