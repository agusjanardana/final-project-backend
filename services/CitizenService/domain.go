package CitizenService

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"gopkg.in/guregu/null.v4"
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
	Birthday        null.Time
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
