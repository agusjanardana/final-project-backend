package SessionDetailService

import "context"

type SessionDetail interface {
	CitizenChooseSession(ctx context.Context, citizenId, sessionId int) ([]SessionDetailDo, error)
}
