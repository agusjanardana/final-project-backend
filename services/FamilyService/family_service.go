package FamilyService

import (
	"context"
)
type FamilyService interface {
	Create(ctx context.Context, id int, family FamilyMember) (FamilyMember, error)
	GetFamilyById(ctx context.Context, id int) (FamilyMember, error)
	GetCitizenOwnFamily(ctx context.Context, citizenId int) ([]FamilyMember, error)
	Update(ctx context.Context, citizenId, id int, family FamilyMember) (FamilyMember, error)
	Delete(ctx context.Context, id, citizenId int) (string, error)
	HfUpdateStatusFamily(ctx context.Context, fmid int, domain FamilyMember) (FamilyMember, error)
}
