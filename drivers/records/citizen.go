package records

import (
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type Citizen struct {
	gorm.Model
	Id              int    `gorm:"primary_key, autoIncrement,not null"`
	Name            string `gorm:"not null"`
	Email           string `gorm:"not null"'`
	Password        string `gorm:"not null"`
	NIK             string `gorm:"not null"`
	Address         string `gorm:"not null"`
	HandphoneNumber int
	VaccinePass     string
	StatusVaccines  string `gorm:"type:enum('BELUM VAKSIN', 'VAKSIN DOSIS 1', 'VAKSIN DOSIS 2');default:'BELUM VAKSIN'"`
	Birthday        null.Time
	FamilyMember    []FamilyMember `gorm:"foreignKey:CitizenId"`
}
