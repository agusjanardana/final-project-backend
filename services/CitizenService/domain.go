package CitizenService

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"time"
	"vaccine-app-be/drivers/records"
)

type Citizen struct {
	Id              int
	Name            string
	Email           string
	Password        string
	NIK             string
	Address         string
	HandphoneNumber string
	VaccinePass     string
	StatusVaccines  string
	Age             int
	Gender          string
	Birthday        time.Time
}

func (c *Citizen) ToRecordFamily() records.FamilyMember {
	recordFamily := records.FamilyMember{
		Name:                 c.Name,
		Birthday:             c.Birthday,
		Nik:                  c.NIK,
		Gender:               c.Gender,
		Age:                  c.Age,
		Handphone:            c.HandphoneNumber,
		CitizenId:            c.Id,
	}
	return recordFamily
}

func (citizen Citizen) Validate() error {
	return validation.ValidateStruct(&citizen,
		//validasi Nama tidak boleh kosong
		validation.Field(&citizen.Name, validation.Required),
		//validasi Email no kosong, and regex
		validation.Field(&citizen.Email, validation.Required, is.Email),
		//validasi NIK tidak boleh kosong
		validation.Field(&citizen.NIK, validation.Required),
		//validation password no kosong
		validation.Field(&citizen.Password, validation.Required),
	)
}
