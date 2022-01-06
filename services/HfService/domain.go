package HfService

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"time"
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
	Vaccine                   []Vaccine
	VaccineSession            []VaccineSession
}

type Vaccine struct {
	Id                  int
	HealthFacilitatorId int
	Name                string
	Stock               int
	VaccineSession      []VaccineSession
}

type VaccineSession struct {
	Id                   int
	StartDate            time.Time
	EndDate              time.Time
	Quota                int
	SessionType          string
	VaccineId            int
	HealthFacilitatorId  int
	Status               string
	VaccineSessionDetail []VaccineSessionDetail
}

type VaccineSessionDetail struct {
	Id             int
	SessionId      int
	FamilyMemberId int
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
