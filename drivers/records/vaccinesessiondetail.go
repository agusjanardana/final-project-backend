package records

import "gorm.io/gorm"

type VaccineSessionDetail struct {
	gorm.Model
	Id             int `gorm:"not null, primaryKey, autoIncrement"`
	SessionId      int
	FamilyMemberId int
}
