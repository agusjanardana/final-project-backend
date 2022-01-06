package VaccineSessionDetailRepository

import (
	"context"
	"vaccine-app-be/drivers/records"
)

type VaccineSessionDetail interface {
	Create(ctx context.Context, sessionId, fmId int)(records.VaccineSessionDetail, error)
	GetDetailBySessionId(ctx context.Context, sessionId int)([]records.VaccineSessionDetail, error)
	GetDetailById(ctx context.Context, id int)(records.VaccineSessionDetail, error)
	GetDetailByFamilyId(ctx context.Context, fmid int)([]records.VaccineSessionDetail, error)
}