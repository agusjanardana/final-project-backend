package web

import (
	"time"
)

//berisi struct req / res controllers dan konversi ke domain sepertinya..

//response
type ResponseRegister struct {
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	NIK             string    `json:"nik"`
	HandphoneNumber string    `json:"handphone_number"`
	Birthday        time.Time `json:"birthday"`
	Age             int       `json:"age"`
	Gender          string    `json:"gender"`
}

type RequestRegister struct {
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	Birthday        time.Time `json:"birthday"`
	Name            string    `json:"name"`
	NIK             string    `json:"nik"`
	HandphoneNumber string    `json:"handphone_number"`
	Age             int       `json:"age"`
	Gender          string    `json:"gender"`
}

type CitizenLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CitizenUpdateRequest struct {
	Birthday time.Time `json:"birthday"`
	Address  string    `json:"address"`
}

type CitizenUpdateResponse struct {
	Address  string
	Birthday time.Time
}

type RespondFind struct {
	Id              int            `json:"id"`
	Name            string         `json:"name"`
	Email           string         `json:"email"`
	NIK             string         `json:"nik"`
	HandphoneNumber string         `json:"handphone_number"`
	Age             int            `json:"age"`
	Gender          string         `json:"gender"`
	Birthday        time.Time      `json:"birthday"`
	Address         string         `json:"address"`
	FamilyMember    []FamilyMember `json:"citizen_family"`
}

type FamilyMember struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	StatusVaccines string `json:"status_vaccines"`
}
