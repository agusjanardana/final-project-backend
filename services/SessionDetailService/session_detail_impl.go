package SessionDetailService

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"vaccine-app-be/drivers/records"
	"vaccine-app-be/drivers/repository/FamilyRepository"
	"vaccine-app-be/drivers/repository/VaccineSessionDetailRepository"
	"vaccine-app-be/drivers/repository/VaccineSessionRepository"
)

type SessionDetailImpl struct {
	sessionDetail  VaccineSessionDetailRepository.VaccineSessionDetail
	family         FamilyRepository.FamilyRepository
	vaccineSession VaccineSessionRepository.VaccineSessionRepository
}

func NewSessionDetail(sessionDetail VaccineSessionDetailRepository.VaccineSessionDetail, family FamilyRepository.FamilyRepository, vaccineSession VaccineSessionRepository.VaccineSessionRepository) SessionDetail {
	return &SessionDetailImpl{sessionDetail: sessionDetail, family: family, vaccineSession: vaccineSession}
}
func (service *SessionDetailImpl) CitizenChooseSession(ctx context.Context, citizenId, sessionId int) ([]SessionDetailDo, error) {
	citizenFamily, err := service.family.GetCitizenOwnFamily(ctx, citizenId)
	if err != nil {
		return nil, err
	}

	//dapatkan jumlah keluarga
	countFamily := len(citizenFamily)

	dataSession, err := service.vaccineSession.FindById(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	if string(rune(dataSession.Id)) == "" {
		return nil, errors.New("session doesn't exist")
	}

	if dataSession.Quota == 0 {
		return nil, errors.New("session quota is full")
	}

	detailData, _ := service.sessionDetail.GetDetailBySessionId(ctx, sessionId)
	countSessionData := len(detailData)

	if countFamily+countSessionData > dataSession.Quota {
		return []SessionDetailDo{}, errors.New("the quota in this session cannot accommodate your family members")
	} else {
		data := make([]SessionDetailDo, countFamily)
		for i, v := range citizenFamily {
			create, err := service.sessionDetail.Create(ctx, sessionId, v.Id)
			if err != nil {
				return []SessionDetailDo{}, err
			}
			data[i].Id = create.Id
			data[i].SessionId = create.SessionId
			data[i].FamilyMemberId = create.FamilyMemberId
		}

		//kurangi jumlah quota di DB untuk data realtime.
		entitySession := records.VaccineSession{}
		entitySession.Quota = dataSession.Quota - countFamily
		_, err := service.vaccineSession.Update(ctx, dataSession.Id, dataSession.HealthFacilitatorId, entitySession)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
}

func (service *SessionDetailImpl) GetDetailBySessionId(ctx context.Context, sessionId int) ([]SessionDetailDo, error) {
	data, err := service.sessionDetail.GetDetailBySessionId(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	var response []SessionDetailDo
	copier.Copy(&response, &data)

	return response, nil
}

func (service *SessionDetailImpl) GetDetailById(ctx context.Context, id int) (SessionDetailDo, error) {
	data, err := service.sessionDetail.GetDetailById(ctx, id)
	if err != nil {
		return SessionDetailDo{}, err

	}
	response := SessionDetailDo{}
	copier.Copy(&response, &data)

	return response, nil
}

func (service *SessionDetailImpl) GetDetailByFamilyId(ctx context.Context, fmid int) ([]SessionDetailDo, error) {
	data, err := service.sessionDetail.GetDetailByFamilyId(ctx, fmid)
	if err != nil {
		return nil, err
	}

	var response []SessionDetailDo
	copier.Copy(&response, &data)
	return response, nil

}
