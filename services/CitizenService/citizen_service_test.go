package CitizenService

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	middleware2 "vaccine-app-be/app/middleware"
	"vaccine-app-be/drivers/records"
	"vaccine-app-be/drivers/repository/CitizenRepository/mocks"
	"vaccine-app-be/utilities"
)

var (
	citizenRepository mocks.CitizenRepository
	jwtAuth           *middleware2.ConfigJWT
	service           CitizenService
)

func setup() CitizenService {
	jwtAuth := &middleware2.ConfigJWT{
		SecretJWT: "testmock123",
		ExpiredIn: 2,
	}
	citizenService := NewCitizenService(&citizenRepository, jwtAuth)

	return citizenService
}

func TestRegister(t *testing.T){
	CitizenService := setup()
	t.Run("test case 1, valid test for register", func(t *testing.T){
		domain := Citizen{
			Name: "agus",
			Email: "bjanardana@yahoo.com",
			Password: "123456",
			NIK : "1234567",
		}
		expectedReturn := records.Citizen{
			Name: "agus",
			Email : "bjanardana@yahoo.com",
			NIK : "1234567",
		}
		citizenRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(records.Citizen{},nil).Once()
		citizenRepository.On("Register", mock.Anything, mock.Anything).Return(expectedReturn, nil).Once()

		actualResult , err := CitizenService.Register(context.Background(), domain)
		assert.Nil(t, err)
		assert.Equal(t, expectedReturn.Name, actualResult.Name)
	})

	t.Run("test case 2, email already used", func(t *testing.T){
		domain := Citizen{
			Name: "agus",
			Email: "bjanardana@yahoo.com",
			Password: "123456",
			NIK : "1234567",
		}
		expectedReturn := records.Citizen{
			Name: "agus",
			Email : "bjanardana@yahoo.com",
			NIK : "1234567",
		}
		citizenRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(expectedReturn,nil).Once()
		citizenRepository.On("Register", mock.Anything, mock.Anything).Return(expectedReturn, nil).Once()

		_ , err := CitizenService.Register(context.Background(), domain)
		assert.Equal(t, err, errors.New("email already used"))
	})

	t.Run("test case 3, validation testing", func(t *testing.T){
		domain := Citizen{
			Name: "agus",
			Email: "bjanardana",
			Password: "123456",
			NIK : "1234567",
		}
		expectedReturn := records.Citizen{
			Name: "agus",
			Email : "bjanardana@gmail.com",
			NIK : "1234567",
		}
		citizenRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(expectedReturn,nil).Once()
		citizenRepository.On("Register", mock.Anything, mock.Anything).Return(expectedReturn, nil).Once()

		actualResult , err := CitizenService.Register(context.Background(), domain)
		assert.NotEqualValues(t, err, expectedReturn.Email, actualResult.Email)
	})
}

func TestLogin(t *testing.T) {
	CitizenService := setup()
	t.Run("test case 1, valid login", func(t *testing.T){
		hashedPassword, _ := utilities.HashPassword("agus123456")
		expecetedReturn := records.Citizen{
			Email: "bjanardana@gmail.com",
			Password: hashedPassword,
		}

		citizenRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(expecetedReturn, nil).Once()

		_, err := CitizenService.Login(context.Background(), "bjanardana@gmail.com", "agus123456")
		assert.Nil(t, err)
	})

	t.Run("test case 2, invalid password", func(t *testing.T){
		hashedPassword, _ := utilities.HashPassword("agus123456")
		expecetedReturn := records.Citizen{
			Email: "bjanardana@gmail.com",
			Password: hashedPassword,
		}

		citizenRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(expecetedReturn, nil).Once()

		_, err := CitizenService.Login(context.Background(), "bjanardana@gmail.com", "asdasdasd")
		assert.Equal(t, err, errors.New("password doesn't match"))
	})
}