package records

import (
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type Citizen struct {
	gorm.Model
	Id              int    `gorm:"primary_key;autoIncrement;not null;"`
	Name            string `gorm:"not null"`
	Email           string `gorm:"not null"'`
	Password        string `gorm:"not null"`
	NIK             string `gorm:"not null"`
	Address         string `gorm:"not null"`
	HandphoneNumber string
	Gender          string
	Age             int
	VaccinePass     string
	Birthday        null.Time
	FamilyMember    []FamilyMember `gorm:"foreignKey:CitizenId"`
	Role            string         `gorm:"default:'USER'"`
}
