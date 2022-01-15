package SessionDetailService

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
	"vaccine-app-be/drivers/records"
	familyMock "vaccine-app-be/drivers/repository/FamilyRepository/mocks"
	detailMock "vaccine-app-be/drivers/repository/VaccineSessionDetailRepository/mocks"
	sessionMock "vaccine-app-be/drivers/repository/VaccineSessionRepository/mocks"
)

var (
	sessionDetail  detailMock.VaccineSessionDetail
	family         familyMock.FamilyRepository
	vaccineSession sessionMock.VaccineSessionRepository
	detailService  SessionDetail
)

func setup() SessionDetail {
	detailService := NewSessionDetail(&sessionDetail, &family, &vaccineSession)

	return detailService
}

func TestCitizenChooseSession(t *testing.T) {
	SessionService := setup()
	t.Run("test case 1, citizen success choose session", func(t *testing.T) {
		citizenId := 1
		sessionId := 1

		//expected
		expectedReturnFamilyOwnedByCitizen := []records.FamilyMember{
			{
				Id:        1,
				Name:      "Agus",
				Birthday:  time.Now(),
				Nik:       "2313131",
				Gender:    "Male",
				Age:       15,
				Handphone: "087762827361",
				CitizenId: 1,
			},
			{
				Id:        2,
				Name:      "Agus2",
				Birthday:  time.Now(),
				Nik:       "2313131",
				Gender:    "Male",
				Age:       15,
				Handphone: "087762827361",
				CitizenId: 1,
			},
		}

		expectedReturnSessionFindById := records.VaccineSession{
			Id:                   1,
			StartDate:            time.Now(),
			EndDate:              time.Now(),
			Quota:                200,
			SessionType:          "SESI 1",
			VaccineId:            1,
			HealthFacilitatorId:  1,
			Status:               "AVAILABLE",
			VaccineSessionDetail: []records.VaccineSessionDetail{},
		}
		//{Id: 1, SessionId: 1, FamilyMemberId: 1}, {Id: 2, SessionId: 1, FamilyMemberId: 2}

		expectedReturnSessionDetail := records.VaccineSessionDetail{
			Id:             1,
			SessionId:      1,
			FamilyMemberId: 1,
		}

		var expectedReturnFindDetailBySessionId []records.VaccineSessionDetail

		family.On("GetCitizenOwnFamily", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturnFamilyOwnedByCitizen, nil).Once()
		sessionDetail.On("GetDetailBySessionId", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturnFindDetailBySessionId, nil).Once()
		vaccineSession.On("FindById", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturnSessionFindById, nil).Once()
		sessionDetail.On("Create", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(expectedReturnSessionDetail, nil).Twice()
		vaccineSession.On("Update", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.Anything).Return(records.VaccineSession{}, nil).Once()

		session, err := SessionService.CitizenChooseSession(context.Background(), citizenId, sessionId)
		assert.Nil(t, err)
		assert.Equal(t, session[0].Id, expectedReturnSessionDetail.Id)

	})
}

func TestGetDetailBySessionId(t *testing.T) {
	SessionService := setup()
	t.Run("test case 1, success find detail by session", func(t *testing.T) {
		sessionId := 1
		expectedReturn := []records.VaccineSessionDetail{
			{
				Id:             1,
				SessionId:      1,
				FamilyMemberId: 1,
			},
			{
				Id:             2,
				SessionId:      1,
				FamilyMemberId: 2,
			},
		}

		sessionDetail.On("GetDetailBySessionId", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturn, nil).Once()
		dataService, err := SessionService.GetDetailBySessionId(context.Background(), sessionId)

		assert.Nil(t, err)
		assert.Equal(t, dataService[0].Id, expectedReturn[0].Id)
	})

	t.Run("test case 2, failed session id with that not found", func(t *testing.T) {
		sessionId := 1

		err2 := errors.New("data not found")
		sessionDetail.On("GetDetailBySessionId", mock.Anything, mock.AnythingOfType("int")).Return([]records.VaccineSessionDetail{}, err2).Once()

		data, err := SessionService.GetDetailBySessionId(context.Background(), sessionId)
		assert.Empty(t, data)
		assert.Equal(t, err, err2)
	})
}

func TestGetDetailById(t *testing.T) {
	SessionService := setup()
	t.Run("test case 1, success find detail by id", func(t *testing.T) {
		detailId := 1
		expectedReturn := records.VaccineSessionDetail{
			Id:             1,
			SessionId:      1,
			FamilyMemberId: 1,
		}

		sessionDetail.On("GetDetailById", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturn, nil).Once()

		data, err := SessionService.GetDetailById(context.Background(), detailId)
		assert.Nil(t, err)
		assert.Equal(t, data.Id, expectedReturn.Id)
	})

	t.Run("test case 2, failed detail by id not found", func(t *testing.T) {
		detailId := 1

		err2 := errors.New("data not found")

		sessionDetail.On("GetDetailById", mock.Anything, mock.AnythingOfType("int")).Return(records.VaccineSessionDetail{}, err2).Once()

		data, err := SessionService.GetDetailById(context.Background(), detailId)
		assert.Empty(t, data)
		assert.Equal(t, err, err2)
	})
}

func TestGetDetailByFamilyId(t *testing.T) {
	SessionService := setup()
	t.Run("test case 1, success find detail by id", func(t *testing.T) {
		familyId := 1
		expectedReturn := []records.VaccineSessionDetail{
			{
				Id:             1,
				SessionId:      1,
				FamilyMemberId: 1,
			},
			{
				Id:             2,
				SessionId:      2,
				FamilyMemberId: 1,
			},
		}

		sessionDetail.On("GetDetailByFamilyId", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturn, nil).Once()

		data, err := SessionService.GetDetailByFamilyId(context.Background(), familyId)
		assert.Equal(t, data[0].Id, expectedReturn[0].Id)
		assert.Nil(t, err)
	})

	t.Run("test case 2, failed find by family id not found", func(t *testing.T) {
		familyId := 1

		err2 := errors.New("data not found")

		sessionDetail.On("GetDetailByFamilyId", mock.Anything, mock.AnythingOfType("int")).Return([]records.VaccineSessionDetail{}, err2).Once()

		data, err := SessionService.GetDetailByFamilyId(context.Background(), familyId)
		assert.Empty(t, data)
		assert.Equal(t, err, err2)
	})
}
