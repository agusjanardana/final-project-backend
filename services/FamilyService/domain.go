package FamilyService

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"time"
)

type FamilyMember struct {
	Id        int
	Name      string
	Birthday  time.Time
	Nik       string
	Gender    string
	Age       int
	Handphone string
	CitizenId int
}

func (fm FamilyMember) Validation() error {
	return validation.ValidateStruct(&fm,
		validation.Field(&fm.Name, validation.Required),
		validation.Field(&fm.Nik, validation.Required),
		validation.Field(&fm.Age, validation.Required),
	)
}
