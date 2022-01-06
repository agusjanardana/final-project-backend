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
