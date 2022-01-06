package web

import "time"

type HealthFaRegisterRequest struct {
	Name      string `json:"facilitator_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Address   string `json:"address"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
	Type      string `json:"type"`
}

type HealthFaRegisterResponse struct {
	Name      string `json:"facilitator_name"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
	Type      string `json:"type"`
}

type HealthFaLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

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
