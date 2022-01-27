package web

type VaccineCreateRequest struct {
	HealthFacilitatorId int    `json:"health_facilitator_id"`
	Name                string `json:"name"`
	Stock               int    `json:"stock"`
}

type VaccineCreateResponse struct {
	Id                  int    `json:"id"`
	HealthFacilitatorId int    `json:"health_facilitator_id"`
	Name                string `json:"name"`
	Stock               int    `json:"stock"`
}

type VaccineUpdateRequest struct {
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

type VaccineUpdateResponse struct {
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

type VaccineFindByIdResponse struct {
	Id                  int    `json:"id"`
	HealthFacilitatorId int    `json:"health_facilitator_id"`
	Name                string `json:"name"`
	Stock               int    `json:"stock"`
}
