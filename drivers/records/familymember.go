package records

import (
	"gorm.io/gorm"
	"time"
)

type FamilyMember struct {
	gorm.Model
	Id        int    `gorm:"not null, primary_key, autoIncrement"`
	Name      string `gorm:"not null"`
	Birthday  time.Time
	Nik       int    `gorm:"not null"`
	Gender    string `gorm:"not null"`
	Age       int    `gorm:"not null"`
	Handphone int    `gorm:"not null"`
	CitizenId int
}
