package CitizenService

import (
	"gorm.io/gorm"
	"time"
)

type Citizen struct {
	gorm.Model
	Id              int    `gorm:"primary_key, autoIncrement,not null"`
	Name            string `gorm:"not null"`
	NIK             string `gorm:"not null"`
	Address         string `gorm:"not null"`
	HandphoneNumber int
	VaccinePass     string
	StatusVaccines  string `gorm:"type:enum('BELUM VAKSIN', 'VAKSIN DOSIS 1', 'VAKSIN DOSIS 2');default:'BELUM VAKSIN'"`
	Birthday        time.Time
}
