package web

type SessionDetailDo struct {
	Id             int `json:"id"`
	SessionId      int `json:"session_id"`
	FamilyMemberId int `json:"family_member_id"`
}
