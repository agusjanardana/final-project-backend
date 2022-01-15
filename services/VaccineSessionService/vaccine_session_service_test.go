package VaccineSessionService

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
	"vaccine-app-be/drivers/records"
	CitizenMock "vaccine-app-be/drivers/repository/CitizenRepository/mocks"
	FamilyMock "vaccine-app-be/drivers/repository/FamilyRepository/mocks"
	VaccineMock "vaccine-app-be/drivers/repository/VaccineRepository/mocks"
	SessionMock "vaccine-app-be/drivers/repository/VaccineSessionRepository/mocks"
)

var (
	vaccineSessionRepository SessionMock.VaccineSessionRepository
	vaccineRepository        VaccineMock.VaccineRepository
	familyRepository         FamilyMock.FamilyRepository
	citizenRepository        CitizenMock.CitizenRepository
)

func setup() VaccineSessionService {
	vaccineSession := NewSessionService(&vaccineSessionRepository, &vaccineRepository, &familyRepository, &citizenRepository)
	return vaccineSession
}

func TestCreateSession(t *testing.T) {
	VaccineService := setup()
	t.Run("test case 1, success create session", func(t *testing.T) {
		domain := VaccineSession{
			StartDate:           time.Now(),
			EndDate:             time.Now().AddDate(2024, 9, 12),
			Quota:               200,
			SessionType:         "SESI 1",
			VaccineId:           1,
			HealthFacilitatorId: 1,
			Status:              "AVAILABLE",
		}

		expectedReturnFind := records.Vaccine{
			Id:                  1,
			HealthFacilitatorId: 1,
			Name:                "Sinovac",
			Stock:               200,
		}

		expectedReturnCreate := records.VaccineSession{
			StartDate:           time.Now(),
			EndDate:             time.Now().AddDate(2024, 9, 12),
			Quota:               200,
			SessionType:         "SESI 1",
			VaccineId:           1,
			HealthFacilitatorId: 1,
			Status:              "AVAILABLE",
		}

		vaccineRepository.On("FindVaccineById", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturnFind, nil).Once()
		vaccineSessionRepository.On("Create", mock.Anything, mock.Anything).Return(expectedReturnCreate, nil).Once()
		vaccineRepository.On("Update", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.Anything).Return(records.Vaccine{}, nil).Once()

		session, err := VaccineService.CreateSession(context.Background(), domain)
		assert.Nil(t, err)
		assert.Equal(t, session.Id, expectedReturnCreate.Id)
	})

	t.Run("test case 2, validate failed", func(t *testing.T) {
		domain := VaccineSession{
			StartDate:           time.Now(),
			Quota:               200,
			SessionType:         "SESI 1",
			VaccineId:           1,
			HealthFacilitatorId: 1,
			Status:              "AVAILABLE",
		}

		session, err := VaccineService.CreateSession(context.Background(), domain)
		assert.NotEmpty(t, err)
		assert.Empty(t, session)
	})

	t.Run("test case 3, find vaccine by id not found", func(t *testing.T) {
		domain := VaccineSession{
			StartDate:           time.Now(),
			EndDate:             time.Now().AddDate(2024, 9, 12),
			Quota:               200,
			SessionType:         "SESI 1",
			VaccineId:           1,
			HealthFacilitatorId: 1,
			Status:              "AVAILABLE",
		}
		err2 := errors.New("data not found")
		vaccineRepository.On("FindVaccineById", mock.Anything, mock.AnythingOfType("int")).Return(records.Vaccine{}, err2).Once()
		session, err := VaccineService.CreateSession(context.Background(), domain)
		assert.Equal(t, err, err2)
		assert.Empty(t, session)
	})

	t.Run("test case 4, vaccine stock not enough", func(t *testing.T) {
		domain := VaccineSession{
			StartDate:           time.Now(),
			EndDate:             time.Now().AddDate(2024, 9, 12),
			Quota:               200,
			SessionType:         "SESI 1",
			VaccineId:           1,
			HealthFacilitatorId: 1,
			Status:              "AVAILABLE",
		}

		expectedReturnFind := records.Vaccine{
			Id:                  1,
			HealthFacilitatorId: 1,
			Name:                "Sinovac",
			Stock:               50,
		}

		vaccineRepository.On("FindVaccineById", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturnFind, nil).Once()

		session, err := VaccineService.CreateSession(context.Background(), domain)
		assert.Equal(t, err, errors.New("vaccine stocks could not meet demand"))
		assert.Empty(t, session)
	})
}

func TestGetSessionById(t *testing.T) {
	VaccineService := setup()
	t.Run("test case 1, success get session by Id", func(t *testing.T) {
		expectedReturnFind := records.VaccineSession{
			Id:                  1,
			StartDate:           time.Now(),
			EndDate:             time.Now().AddDate(2024, 9, 12),
			Quota:               200,
			SessionType:         "SESI 1",
			VaccineId:           1,
			HealthFacilitatorId: 1,
			Status:              "AVAILABLE",
		}

		vaccineSessionRepository.On("FindById", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturnFind, nil).Once()

		dataService, err := VaccineService.GetSessionById(context.Background(), 1)
		assert.Equal(t, dataService.Id, expectedReturnFind.Id)
		assert.Nil(t, err)
	})

	t.Run("test case 2, vaccine session not found", func(t *testing.T) {
		err2 := errors.New("data not found")
		vaccineSessionRepository.On("FindById", mock.Anything, mock.AnythingOfType("int")).Return(records.VaccineSession{}, err2).Once()
		dataService, err := VaccineService.GetSessionById(context.Background(), 1)
		assert.Equal(t, err, err2)
		assert.Empty(t, dataService)
	})
}

func TestGetSessionOwnedByHf(t *testing.T) {
	VaccineService := setup()
	t.Run("test case 1, success get session owned by hf", func(t *testing.T) {
		expectedReturnFind := []records.VaccineSession{
			{
				Id:                  1,
				StartDate:           time.Now(),
				EndDate:             time.Now().AddDate(2024, 9, 12),
				Quota:               200,
				SessionType:         "SESI 1",
				VaccineId:           1,
				HealthFacilitatorId: 1,
				Status:              "AVAILABLE",
			},
			{
				Id:                  2,
				StartDate:           time.Now(),
				EndDate:             time.Now().AddDate(2024, 9, 12),
				Quota:               200,
				SessionType:         "SESI 2",
				VaccineId:           1,
				HealthFacilitatorId: 1,
				Status:              "AVAILABLE",
			},
		}

		vaccineSessionRepository.On("FindSessionOwnedByHf", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturnFind, nil).Once()
		dataService, err := VaccineService.GetSessionOwnedByHf(context.Background(), 1)
		assert.Nil(t, err)
		assert.Equal(t, dataService[0].Id, expectedReturnFind[0].Id)
	})

	t.Run("test case 2, data not found", func(t *testing.T) {
		err2 := errors.New("data not found")
		vaccineSessionRepository.On("FindSessionOwnedByHf", mock.Anything, mock.AnythingOfType("int")).Return([]records.VaccineSession{}, err2).Once()
		dataService, err := VaccineService.GetSessionOwnedByHf(context.Background(), 1)
		assert.Equal(t, err, err2)
		assert.Empty(t, dataService)
	})
}

func TestDeleteSession(t *testing.T) {
	VaccineService := setup()
	t.Run("test case 1, success delete session", func(t *testing.T) {
		expectedReturnFind := records.VaccineSession{
			Id:                  1,
			StartDate:           time.Now(),
			EndDate:             time.Now().AddDate(2024, 9, 12),
			Quota:               200,
			SessionType:         "SESI 1",
			VaccineId:           1,
			HealthFacilitatorId: 1,
			Status:              "AVAILABLE",
		}

		vaccineSessionRepository.On("FindById", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturnFind, nil).Once()
		vaccineSessionRepository.On("Delete", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(records.VaccineSession{}, nil).Once()

		sessionService, err := VaccineService.DeleteSession(context.Background(), 1, 1)
		assert.Nil(t, err)
		assert.Equal(t, sessionService, "success delete data")
	})
	t.Run("test case 2, data not found", func(t *testing.T) {
		err2 := errors.New("data not found")
		vaccineSessionRepository.On("FindById", mock.Anything, mock.AnythingOfType("int")).Return(records.VaccineSession{}, err2).Once()

		sessionService, err := VaccineService.DeleteSession(context.Background(), 1, 1)
		assert.Empty(t, sessionService)
		assert.Equal(t, err, err2)
	})

	t.Run("test case 3, this session doesnt blongs to you", func(t *testing.T) {
		expectedReturnFind := records.VaccineSession{
			Id:                  1,
			StartDate:           time.Now(),
			EndDate:             time.Now().AddDate(2024, 9, 12),
			Quota:               200,
			SessionType:         "SESI 1",
			VaccineId:           1,
			HealthFacilitatorId: 2,
			Status:              "AVAILABLE",
		}

		vaccineSessionRepository.On("FindById", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturnFind, nil).Once()

		sessionService, err := VaccineService.DeleteSession(context.Background(), 1, 1)
		assert.Empty(t, sessionService)
		assert.Equal(t, err, errors.New("this session doesn't belongs to you"))
	})
}

func TestUpdateSession(t *testing.T) {
	VaccineService := setup()
	t.Run("test case 1, success update", func(t *testing.T) {
		domain := VaccineSession{
			StartDate:           time.Now(),
			EndDate:             time.Now().AddDate(2024, 9, 12),
			Quota:               200,
			SessionType:         "SESI 1 EDIT",
			VaccineId:           1,
			HealthFacilitatorId: 1,
			Status:              "AVAILABLE",
		}

		expectedReturnUpdate := records.VaccineSession{
			StartDate:           time.Now(),
			EndDate:             time.Now().AddDate(2024, 9, 12),
			Quota:               200,
			SessionType:         "SESI 1 EDIT",
			VaccineId:           1,
			HealthFacilitatorId: 1,
			Status:              "AVAILABLE",
		}

		expectedReturnFind := records.VaccineSession{
			Id:                  1,
			StartDate:           time.Now(),
			EndDate:             time.Now().AddDate(2024, 9, 12),
			Quota:               200,
			SessionType:         "SESI 1",
			VaccineId:           1,
			HealthFacilitatorId: 1,
			Status:              "AVAILABLE",
		}

		vaccineSessionRepository.On("FindById", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturnFind, nil).Once()
		vaccineSessionRepository.On("Update", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.Anything).Return(expectedReturnUpdate, nil).Once()

		session, err := VaccineService.UpdateSession(context.Background(), 1, 1, domain)
		assert.Nil(t, err)
		assert.Equal(t, session.SessionType, expectedReturnUpdate.SessionType)
	})

	t.Run("test case 2, validate failed", func(t *testing.T) {
		domain := VaccineSession{
			StartDate:           time.Now(),
			Quota:               200,
			SessionType:         "SESI 1",
			VaccineId:           1,
			HealthFacilitatorId: 1,
			Status:              "AVAILABLE",
		}
		session, err := VaccineService.UpdateSession(context.Background(), 1, 1, domain)
		assert.NotEmpty(t, err)
		assert.Empty(t, session)
	})

	t.Run("test case 2, data not found", func(t *testing.T) {
		domain := VaccineSession{
			StartDate:           time.Now(),
			EndDate:             time.Now().AddDate(2024, 9, 12),
			Quota:               200,
			SessionType:         "SESI 1",
			VaccineId:           1,
			HealthFacilitatorId: 1,
			Status:              "AVAILABLE",
		}
		err2 := errors.New("data not found")
		vaccineSessionRepository.On("FindById", mock.Anything, mock.AnythingOfType("int")).Return(records.VaccineSession{}, err2).Once()

		session, err := VaccineService.UpdateSession(context.Background(), 1, 1, domain)
		assert.Empty(t, session)
		assert.Equal(t, err, err2)
	})

	t.Run("test case 3, doesnt belongs to HF", func(t *testing.T) {
		domain := VaccineSession{
			StartDate:           time.Now(),
			EndDate:             time.Now().AddDate(2024, 9, 12),
			Quota:               200,
			SessionType:         "SESI 1 EDIT",
			VaccineId:           1,
			HealthFacilitatorId: 1,
			Status:              "AVAILABLE",
		}

		expectedReturnFind := records.VaccineSession{
			Id:                  1,
			StartDate:           time.Now(),
			EndDate:             time.Now().AddDate(2024, 9, 12),
			Quota:               200,
			SessionType:         "SESI 1",
			VaccineId:           1,
			HealthFacilitatorId: 2,
			Status:              "AVAILABLE",
		}

		vaccineSessionRepository.On("FindById", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturnFind, nil).Once()

		session, err := VaccineService.UpdateSession(context.Background(), 1, 1, domain)
		assert.Empty(t, session)
		assert.Equal(t, err, errors.New("cannot update this session, doesn't belong to yours"))
	})
}

func TestGetAllVaccineSession(t *testing.T) {
	VaccineService := setup()
	t.Run("test case 1, success get all vaccine Session", func(t *testing.T) {
		expectedReturn := []records.VaccineSession{
			{
				Id:                  1,
				StartDate:           time.Now(),
				EndDate:             time.Now().AddDate(2024, 9, 12),
				Quota:               200,
				SessionType:         "SESI 1",
				VaccineId:           1,
				HealthFacilitatorId: 2,
				Status:              "AVAILABLE",
			},
			{
				Id:                  2,
				StartDate:           time.Now(),
				EndDate:             time.Now().AddDate(2024, 9, 12),
				Quota:               200,
				SessionType:         "SESI 2",
				VaccineId:           1,
				HealthFacilitatorId: 2,
				Status:              "AVAILABLE",
			},
			{
				Id:                  3,
				StartDate:           time.Now(),
				EndDate:             time.Now().AddDate(2024, 9, 12),
				Quota:               200,
				SessionType:         "SESI 3",
				VaccineId:           1,
				HealthFacilitatorId: 2,
				Status:              "AVAILABLE",
			},
		}
		vaccineSessionRepository.On("GetAllVaccineSession", mock.Anything).Return(expectedReturn, nil).Once()

		session, err := VaccineService.GetAllVaccineSession(context.Background())
		assert.Nil(t, err)
		assert.NotEmpty(t, session)
	})

	t.Run("test case 2, failed not found the data", func(t *testing.T) {
		err2 := errors.New("data not found")
		vaccineSessionRepository.On("GetAllVaccineSession", mock.Anything).Return([]records.VaccineSession{}, err2).Once()
		session, err := VaccineService.GetAllVaccineSession(context.Background())
		assert.Equal(t, err, err2)
		assert.Empty(t, session)
	})
}

func TestGetCitizenAndFamilySelectedSession(t *testing.T) {
	VaccineService := setup()
	t.Run("test case 1, success get citizen and selected session", func(t *testing.T) {
		expectedReturnFindCitizen := records.Citizen{Id: 1, Name: "Agus"}
		expectedReturnCitizenOwn := []records.FamilyMember{
			{
				Id: 1,
				Name: "Agus",
				CitizenId: 1,
			},
			{
				Id: 2,
				Name: "Juna",
				CitizenId: 1,
			},
			{
				Id: 3,
				Name: "Krisna",
				CitizenId: 1,
			},
		}

		expectedReturnFind := records.VaccineSession{
			Id:                  1,
			StartDate:           time.Now(),
			EndDate:             time.Now().AddDate(2024, 9, 12),
			Quota:               200,
			SessionType:         "SESI 1",
			VaccineId:           1,
			HealthFacilitatorId: 2,
			Status:              "AVAILABLE",
		}

		citizenRepository.On("FindById", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturnFindCitizen, nil).Once()
		familyRepository.On("GetCitizenOwnFamily", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturnCitizenOwn, nil).Once()
		vaccineSessionRepository.On("FindById", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturnFind, nil).Once()

		sessionData, err := VaccineService.GetCitizenAndFamilySelectedSession(context.Background(), 1)
		assert.Nil(t, err)
		assert.NotEmpty(t, sessionData)
	})

	t.Run("test case 2, citizen not found", func(t *testing.T) {

		citizenRepository.On("FindById", mock.Anything, mock.AnythingOfType("int")).Return(records.Citizen{}, errors.New("data not found")).Once()
		sessionData, err := VaccineService.GetCitizenAndFamilySelectedSession(context.Background(), 1)
		assert.Empty(t, sessionData)
		assert.Equal(t, err, errors.New("data not found"))
	})

	t.Run("test case 3, failed family", func(t *testing.T) {
		expectedReturnFindCitizen := records.Citizen{Id: 1, Name: "Agus"}

		citizenRepository.On("FindById", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturnFindCitizen, nil).Once()
		familyRepository.On("GetCitizenOwnFamily", mock.Anything, mock.AnythingOfType("int")).Return([]records.FamilyMember{}, errors.New("data not found")).Once()

		sessionData, err := VaccineService.GetCitizenAndFamilySelectedSession(context.Background(), 1)
		assert.Empty(t, sessionData)
		assert.Equal(t, err, errors.New("data not found"))
	})
}
