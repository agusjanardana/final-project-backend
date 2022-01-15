package web

import (
	"time"
)

type FamilyMemberRequest struct {
	Name      string    `json:"name"`
	Birthday  time.Time `json:"birthday"`
	Nik       string    `json:"nik"`
	Gender    string    `json:"gender"`
	Age       int       `json:"age"`
	Handphone string    `json:"handphone"`
}

type FamilyMemberResponse struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	Birthday       time.Time `json:"birthday"`
	StatusVaccines string    `json:"status_vaccines"`
	Nik            string    `json:"nik"`
	Gender         string    `json:"gender"`
	Age            int       `json:"age"`
	Handphone      string    `json:"handphone"`
}
