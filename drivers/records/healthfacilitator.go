package records

import "gorm.io/gorm"

type HealthFacilitator struct {
	gorm.Model
	Id                        int    `gorm:"primary_key;autoIncrement;not null"`
	Name                      string `gorm:"not null"`
	Email                     string `gorm:"not null"`
	Password                  string `gorm:"not null"`
	FacilitatorsCertification string
	Address                   string
	Longitude                 string    `gorm:"not null"`
	Latitude                  string    `gorm:"not null"`
	Type                      string    `gorm:"not null"`
	Role                      string    `gorm:"default:'ADMIN'"`
	Vaccine                   []Vaccine `gorm:"foreignKey:HealthFacilitatorId;"`
}
