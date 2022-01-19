package CitizenService

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log"
	"testing"
	"time"
	middleware2 "vaccine-app-be/app/middleware"
	"vaccine-app-be/drivers/records"
	"vaccine-app-be/drivers/repository/CitizenRepository/mocks"
	familyMock "vaccine-app-be/drivers/repository/FamilyRepository/mocks"
	detailMock "vaccine-app-be/drivers/repository/VaccineSessionDetailRepository/mocks"
	sessionMock "vaccine-app-be/drivers/repository/VaccineSessionRepository/mocks"
	"vaccine-app-be/utilities"
)

var (
	citizenRepository mocks.CitizenRepository
	familyRepository  familyMock.FamilyRepository
	sessionRepository sessionMock.VaccineSessionRepository
	detailRepository  detailMock.VaccineSessionDetail
	jwtAuth           *middleware2.ConfigJWT
	service           CitizenService
)

func setup() CitizenService {
	jwtAuth := &middleware2.ConfigJWT{
		SecretJWT: "testmock123",
		ExpiredIn: 2,
	}
	citizenService := NewCitizenService(&citizenRepository, jwtAuth, &familyRepository, &sessionRepository, &detailRepository)

	return citizenService
}

func TestRegister(t *testing.T) {
	CitizenService := setup()
	t.Run("test case 1, valid test for register", func(t *testing.T) {
		domain := Citizen{
			Name:     "agus",
			Email:    "bjanardana@yahoo.com",
			Password: "123456",
			NIK:      "1234567",
		}
		expectedReturn := records.Citizen{
			Name:  "agus",
			Email: "bjanardana@yahoo.com",
			NIK:   "1234567",
		}
		citizenRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(records.Citizen{}, nil).Once()
		citizenRepository.On("Register", mock.Anything, mock.Anything).Return(expectedReturn, nil).Once()
		familyRepository.On("Create", mock.Anything, mock.Anything).Return(records.FamilyMember{}, nil).Once()

		actualResult, err := CitizenService.Register(context.Background(), domain)
		assert.Nil(t, err)
		assert.Equal(t, expectedReturn.Name, actualResult.Name)
	})

	t.Run("test case 2, email already used", func(t *testing.T) {
		domain := Citizen{
			Name:     "agus",
			Email:    "bjanardana@yahoo.com",
			Password: "123456",
			NIK:      "1234567",
		}
		expectedReturn := records.Citizen{
			Name:  "agus",
			Email: "bjanardana@yahoo.com",
			NIK:   "1234567",
		}
		citizenRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(expectedReturn, nil).Once()
		citizenRepository.On("Register", mock.Anything, mock.Anything).Return(expectedReturn, nil).Once()

		_, err := CitizenService.Register(context.Background(), domain)
		assert.Equal(t, err, errors.New("email already used"))
		assert.Equal(t, expectedReturn.Email, domain.Email)
	})

	t.Run("test case 3, validation testing", func(t *testing.T) {
		domain := Citizen{
			Name:     "agus",
			Email:    "bjanardana",
			Password: "123456",
			NIK:      "1234567",
		}
		expectedReturn := records.Citizen{
			Name:  "agus",
			Email: "bjanardana@gmail.com",
			NIK:   "1234567",
		}
		citizenRepository.On("Register", mock.Anything, mock.Anything).Return(expectedReturn, nil).Once()

		actualResult, err := CitizenService.Register(context.Background(), domain)
		assert.NotEqualValues(t, err, expectedReturn.Email, actualResult.Email)
	})
}

func TestLogin(t *testing.T) {
	CitizenService := setup()
	t.Run("test case 1, success login", func(t *testing.T) {
		hashedPassword, _ := utilities.HashPassword("agus123456")
		log.Println("test print password", hashedPassword)
		result := Citizen{Email: "bjanardana@yahoo.com", Password: "agus123456"}
		expecetedReturns := records.Citizen{
			Email:    "bjanardana@yahoo.com",
			Password: hashedPassword,
		}

		citizenRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(expecetedReturns, nil).Once()

		_, err := CitizenService.Login(context.Background(), result.Email, result.Password)
		assert.Nil(t, err)
	})

	t.Run("test case 2, invalid password", func(t *testing.T) {
		hashedPassword, _ := utilities.HashPassword("agus123456")
		result := Citizen{Email: "bjanardana@yahoo.com", Password: "asdasdasd"}
		expecetedReturn := records.Citizen{
			Email:    "bjanardana@yahoo.com",
			Password: hashedPassword,
		}

		citizenRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(expecetedReturn, nil).Once()

		_, err := CitizenService.Login(context.Background(), result.Email, result.Password)
		assert.Equal(t, err, errors.New("password doesn't match"))
	})
}

func TestUpdate(t *testing.T) {
	CitizenService := setup()
	t.Run("test case 1, valid update", func(t *testing.T) {
		ctxId := 1
		birthdays := time.Now()
		domain := Citizen{
			Birthday: birthdays,
			Address:  "jalan",
		}

		expectedReturn := records.Citizen{Address: "jalan"}
		expectedReturn.Birthday.Scan(birthdays)

		expectedReturnFamily := []records.FamilyMember{
			{
				Name: "Agus",
				Age:  12,
			},
			{
				Name: "Juna",
				Age:  12,
			},
		}
		expectedReturnEdit := records.FamilyMember{
			Id:        0,
			Name:      "Agus",
			Birthday:  birthdays,
			Nik:       "123123",
			Gender:    "Male",
			Age:       13,
			Handphone: "123123123",
			CitizenId: 1,
		}

		citizenRepository.On("Update", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("time.Time"), mock.AnythingOfType("string")).Return(expectedReturn, nil).Once()
		familyRepository.On("GetCitizenOwnFamily", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturnFamily, nil).Once()
		familyRepository.On("Update", mock.Anything, mock.AnythingOfType("int"), mock.Anything).Return(expectedReturnEdit, nil).Once()

		updateData, err := CitizenService.Update(context.Background(), ctxId, domain.Birthday, domain.Address)
		assert.Nil(t, err)
		assert.Equal(t, updateData.Birthday, domain.Birthday)
	})

	t.Run("test case 1, invalid update#1", func(t *testing.T) {
		ctxId := 1
		birthdays := time.Now()
		domain := Citizen{
			Birthday: birthdays,
			Address:  "jalan",
		}

		expectedReturn := records.Citizen{Address: "jalan"}
		expectedReturn.Birthday.Scan(birthdays)
		errors := errors.New("data not found")

		citizenRepository.On("Update", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("time.Time"), mock.AnythingOfType("string")).Return(expectedReturn, nil).Once()
		familyRepository.On("GetCitizenOwnFamily", mock.Anything, mock.AnythingOfType("int")).Return([]records.FamilyMember{}, errors).Once()

		_, err := CitizenService.Update(context.Background(), ctxId, domain.Birthday, domain.Address)
		assert.Equal(t, err, errors)

	})
}

func TestFindById(t *testing.T) {
	CitizenService := setup()
	t.Run("test case 1, valid find by id", func(t *testing.T) {
		ctxId := 1
		expectedReturn := records.Citizen{
			Id:              1,
			Name:            "agus",
			Email:           "bjanardana@gmail.com",
			Password:        "23123asdasd",
			NIK:             "123123",
			Address:         "jalan",
			HandphoneNumber: "08123123123",
			Gender:          "Male",
			Age:             13,
			VaccinePass:     "ADa",
		}

		citizenRepository.On("FindById", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturn, nil).Once()

		dataCitizen, err := CitizenService.CitizenFindById(context.Background(), ctxId)
		assert.Nil(t, err)
		assert.Equal(t, dataCitizen.Name, expectedReturn.Name)
	})

	t.Run("test case 1, data not found", func(t *testing.T) {
		ctxId := 1
		errors := errors.New("data not found")
		citizenRepository.On("FindById", mock.Anything, mock.AnythingOfType("int")).Return(records.Citizen{}, errors).Once()

		_, err := CitizenService.CitizenFindById(context.Background(), ctxId)
		assert.Equal(t, err, errors)
	})

}
