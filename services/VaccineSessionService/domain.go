package VaccineSessionService

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"time"
)

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

func (vs VaccineSession) ValidateRequest() error {
	return validation.ValidateStruct(&vs,
		validation.Field(&vs.StartDate, validation.Required),
		validation.Field(&vs.EndDate, validation.Required),
		validation.Field(&vs.Quota, validation.Required),
		validation.Field(&vs.SessionType, validation.Required),
	)
}

type UniqueSession struct {
	Id int
}
