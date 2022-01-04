package records

import "gorm.io/gorm"

type Vaccine struct {
	gorm.Model
	Id                  int    `gorm:"primary_key;autoIncrement;not null"`
	HealthFacilitatorId int    `gorm:"not null"`
	Name                string `gorm:"not null"`
	Stock               int    `gorm:"not null"`
	VaccineSession      []VaccineSession `gorm:"foreignKey:VaccineId;constraint:OnDelete:CASCADE;"`
}
