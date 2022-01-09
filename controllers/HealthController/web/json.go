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
	Name           string           `json:"name"`
	Address        string           `json:"address"`
	Longitude      string           `json:"longitude"`
	Latitude       string           `json:"latitude"`
	Type           string           `json:"type"`
	Vaccine        []Vaccine        `json:"vaccine"`
	VaccineSession []VaccineSession `json:"vaccine_session"`
}

type Vaccine struct {
	Id                  int    `json:"id"`
	HealthFacilitatorId int    `json:"health_facilitator_id"`
	Name                string `json:"name"`
	Stock               int    `json:"stock"`
}

type VaccineSession struct {
	Id                   int                    `json:"id"`
	StartDate            time.Time              `json:"start_date"`
	EndDate              time.Time              `json:"end_date"`
	Quota                int                    `json:"quota"`
	SessionType          string                 `json:"session_type"`
	VaccineId            int                    `json:"vaccine_id"`
	HealthFacilitatorId  int                    `json:"health_facilitator_id"`
	Status               string                 `json:"status"`
	VaccineSessionDetail []VaccineSessionDetail `json:"vaccine_session_detail"`
}

type VaccineSessionDetail struct {
	Id             int `json:"id"`
	SessionId      int `json:"session_id"`
	FamilyMemberId int `json:"family_member_id"`
}
