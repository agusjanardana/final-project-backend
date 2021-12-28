package FamilyRepository

import (
	"context"
	"vaccine-app-be/drivers/records"
)

type FamilyRepository interface {
	Create(ctx context.Context, family records.FamilyMember) (records.FamilyMember, error)
	GetFamilyById(ctx context.Context, id int) (records.FamilyMember, error)
	GetCitizenOwnFamily(ctx context.Context, citizenId int) ([]records.FamilyMember, error)
	Update(ctx context.Context, id int, family records.FamilyMember) (records.FamilyMember, error)
	Delete(ctx context.Context, id , citizenId int) (records.FamilyMember, error)
}
