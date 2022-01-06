package records

import (
	"gorm.io/gorm"
	"time"
)

type VaccineSession struct {
	gorm.Model
	Id                   int `gorm:"not null;primaryKey;autoIncrement"`
	StartDate            time.Time
	EndDate              time.Time
	Quota                int
	SessionType          string
	VaccineId            int
	HealthFacilitatorId  int
	Status               string                 `gorm:"default:'AVAILABLE'"`
	VaccineSessionDetail []VaccineSessionDetail `gorm:"foreignKey:SessionId;constraint:OnDelete:CASCADE;"`
}
