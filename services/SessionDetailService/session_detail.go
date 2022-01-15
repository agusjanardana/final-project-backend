package SessionDetailService

import (
	"context"
)

type SessionDetail interface {
	CitizenChooseSession(ctx context.Context, citizenId, sessionId int) ([]SessionDetailDo, error)
	GetDetailBySessionId(ctx context.Context, sessionId int)([]SessionDetailDo, error)
	GetDetailById(ctx context.Context, id int)(SessionDetailDo, error)
	GetDetailByFamilyId(ctx context.Context, fmid int)([]SessionDetailDo, error)
}
