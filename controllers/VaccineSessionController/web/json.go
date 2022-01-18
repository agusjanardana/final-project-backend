package web

import "time"

type VaccineSessionCreateRequest struct {
	StartDate           time.Time `json:"start_date"`
	EndDate             time.Time `json:"end_date"`
	Quota               int       `json:"quota"`
	SessionType         string    `json:"session_type"`
	VaccineId           int       `json:"vaccine_id"`
	HealthFacilitatorId int       `json:"health_facilitator_id"`
	Status              string    `json:"status"`
}

type VaccineSessionCreateResponse struct {
	Id                  int       `json:"id"`
	StartDate           time.Time `json:"start_date"`
	EndDate             time.Time `json:"end_date"`
	Quota               int       `json:"quota"`
	SessionType         string    `json:"session_type"`
	VaccineId           int       `json:"vaccine_id "`
	HealthFacilitatorId int       `json:"health_facilitator_id"`
	Status              string    `json:"status"`
}

type VaccineSessionCreateResponseOwnedCitizen struct {
	Id                  int                       `json:"id"`
	StartDate           time.Time                 `json:"start_date"`
	EndDate             time.Time                 `json:"end_date"`
	Quota               int                       `json:"quota"`
	SessionType         string                    `json:"session_type"`
	VaccineId           int                       `json:"vaccine_id "`
	HealthFacilitatorId int                       `json:"health_facilitator_id"`
	Status              string                    `json:"status"`
	Vaccine             Vaccine                   `json:"vaccine"`
	HealthFacilitator   HealthFacilitatorsCitizen `json:"health_facilitator"`
}

type Vaccine struct {
	Name string `json:"name"`
}

type HealthFacilitatorsCitizen struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type VaccineSessionFindByIdResponse struct {
	Id                   int                    `json:"id"`
	StartDate            time.Time              `json:"start_date"`
	EndDate              time.Time              `json:"end_date"`
	Quota                int                    `json:"quota"`
	SessionType          string                 `json:"session_type"`
	VaccineId            int                    `json:"vaccine_id "`
	HealthFacilitatorId  int                    `json:"health_facilitator_id"`
	Status               string                 `json:"status"`
	VaccineSessionDetail []VaccineSessionDetail `json:"vaccine_session_detail"`
}

type VaccineSessionDetail struct {
	Id             int `json:"id"`
	SessionId      int `json:"session_id"`
	FamilyMemberId int `json:"family_member_id"`
}
