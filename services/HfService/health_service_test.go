package HfService

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log"
	"testing"
	"vaccine-app-be/app/middleware"
	"vaccine-app-be/drivers/records"
	"vaccine-app-be/drivers/repository/HealthRepository/mocks"
	"vaccine-app-be/utilities"
)

var (
	healthRepository mocks.HealthRepository
	jwtAuth          *middleware.ConfigJWT
	service          HealthService
)

func setup() HealthService {
	jwtAuth := &middleware.ConfigJWT{
		SecretJWT: "testmock123",
		ExpiredIn: 2,
	}
	healthService := NewHealthService(&healthRepository, jwtAuth)

	return healthService
}

func TestRegister(t *testing.T) {
	HealthService := setup()
	t.Run("test case 1, valid test for register", func(t *testing.T) {
		domain := HealthFacilitator{
			Name:      "agus",
			Email:     "bjanardana@yahoo.com",
			Password:  "123456",
			Address:   "123asd",
			Latitude:  "123asdas",
			Longitude: "asdasd",
		}
		expectedReturn := records.HealthFacilitator{
			Name:  "agus",
			Email: "bjanardana@yahoo.com",
		}
		healthRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(records.HealthFacilitator{}, nil).Once()
		healthRepository.On("Register", mock.Anything, mock.Anything).Return(expectedReturn, nil).Once()

		actualResult, err := HealthService.Register(context.Background(), domain)
		assert.Nil(t, err)
		assert.Equal(t, expectedReturn.Name, actualResult.Name)
	})

	t.Run("test case 2, email already used", func(t *testing.T) {
		domain := HealthFacilitator{
			Name:      "agus",
			Email:     "bjanardana@yahoo.com",
			Password:  "123456",
			Address:   "123asd",
			Latitude:  "123asdas",
			Longitude: "asdasd",
		}
		expectedReturn := records.HealthFacilitator{
			Name:  "agus",
			Email: "bjanardana@yahoo.com",
		}
		healthRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(expectedReturn, nil).Once()
		healthRepository.On("Register", mock.Anything, mock.Anything).Return(expectedReturn, nil).Once()

		_, err := HealthService.Register(context.Background(), domain)
		assert.Equal(t, err, errors.New("email already used"))
		assert.Equal(t, expectedReturn.Email, domain.Email)
	})

	t.Run("test case 3, validation testing", func(t *testing.T) {
		domain := HealthFacilitator{
			Name:      "agus",
			Email:     "bjanardana@yahoo.com",
			Password:  "123456",
			Address:   "123asd",
			Latitude:  "123asdas",
			Longitude: "asdasd",
		}
		expectedReturn := records.HealthFacilitator{
			Name:  "agus",
			Email: "bjanardana@gmail.com",
		}
		healthRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(expectedReturn, nil).Once()
		healthRepository.On("Register", mock.Anything, mock.Anything).Return(expectedReturn, nil).Once()

		actualResult, err := HealthService.Register(context.Background(), domain)
		assert.NotEqualValues(t, err, expectedReturn.Email, actualResult.Email)
	})
}

func TestLogin(t *testing.T) {
	HealthService := setup()
	t.Run("test case 1, success login", func(t *testing.T) {
		hashedPassword, _ := utilities.HashPassword("agus123456")
		log.Println("test print password", hashedPassword)
		result := HealthFacilitator{Email: "bjanardana@yahoo.com", Password: "agus123456"}
		expecetedReturns := records.HealthFacilitator{
			Email:    "bjanardana@yahoo.com",
			Password: hashedPassword,
		}

		healthRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(expecetedReturns, nil).Once()

		_, err := HealthService.Login(context.Background(), result.Email, result.Password)
		assert.Nil(t, err)
	})

	t.Run("test case 2, invalid password", func(t *testing.T) {
		hashedPassword, _ := utilities.HashPassword("agus123456")
		result := HealthFacilitator{Email: "bjanardana@yahoo.com", Password: "asdasdasd"}
		expecetedReturn := records.HealthFacilitator{
			Email:    "bjanardana@yahoo.com",
			Password: hashedPassword,
		}

		healthRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(expecetedReturn, nil).Once()

		_, err := HealthService.Login(context.Background(), result.Email, result.Password)
		assert.Equal(t, err, errors.New("password doesn't match"))
	})
}

func TestGetAllHf(t *testing.T) {
	HealthService := setup()
	t.Run("test case 1, success login", func(t *testing.T) {
		expectedReturn := []records.HealthFacilitator{
			{
				Name:    "RS A",
				Address: "Jalan Kembang",
			},
			{
				Name:    "RS B",
				Address: "Jalan Melayu",
			},
		}

		healthRepository.On("GetAllHealthFacilitator", mock.Anything).Return(expectedReturn, nil).Once()

		facilitatorData, err := HealthService.GetAllHealthFacilitator(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, facilitatorData[0].Name, expectedReturn[0].Name)
	})
}

func TestFindById(t *testing.T) {
	HealthService := setup()
	t.Run("test case 1, success find by id", func(t *testing.T) {
		Id := 1
		expectedReturn := records.HealthFacilitator{
			Id:        1,
			Name:      "agus",
			Email:     "bjanardana@yahoo.com",
			Password:  "123456",
			Address:   "123asd",
			Latitude:  "123asdas",
			Longitude: "asdasd",
		}
		healthRepository.On("FindById", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturn, nil).Once()

		dataService, err := HealthService.FindById(context.Background(), Id)
		assert.Nil(t, err)
		assert.Equal(t, dataService.Id, expectedReturn.Id)

	})
}

func TestHfUpdate(t *testing.T) {
	HealthService := setup()
	t.Run("test case 1, success update", func(t *testing.T) {
		hfid := 1
		domain := HealthFacilitator{
			Id:        1,
			Name:      "agusEDIT",
			Email:     "bjanardana@yahoo.com",
			Password:  "123456",
			Address:   "123asd",
			Latitude:  "123asdas",
			Longitude: "asdasd",
		}
		expectedReturn := records.HealthFacilitator{
			Id:        1,
			Name:      "agus",
			Email:     "bjanardana@yahoo.com",
			Password:  "123456",
			Address:   "123asd",
			Latitude:  "123asdas",
			Longitude: "asdasd",
		}

		expectedReturnEdit := records.HealthFacilitator{
			Id:        1,
			Name:      "agusEDIT",
			Email:     "bjanardana@yahoo.com",
			Password:  "123456",
			Address:   "123asd",
			Latitude:  "123asdas",
			Longitude: "asdasd",
		}
		healthRepository.On("FindById", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturn, nil).Once()
		healthRepository.On("Update", mock.Anything, mock.AnythingOfType("int"), mock.Anything).Return(expectedReturnEdit, nil).Once()

		updateData, err := HealthService.Update(context.Background(), hfid, domain)
		assert.Nil(t, err)
		assert.Equal(t, updateData.Name, expectedReturnEdit.Name)
	})
}
