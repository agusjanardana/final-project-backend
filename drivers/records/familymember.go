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
	Nik       string `gorm:"not null"`
	Gender    string `gorm:"not null"`
	Age       int    `gorm:"not null"`
	Handphone string `gorm:"not null"`
	CitizenId int
}
