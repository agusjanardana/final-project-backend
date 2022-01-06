package FamilyService

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
	"vaccine-app-be/drivers/records"
	"vaccine-app-be/drivers/repository/FamilyRepository/mocks"
)

var (
	familyRepository mocks.FamilyRepository
)

func setup() FamilyService {
	familyService := NewFamilyService(&familyRepository)

	return familyService
}

func TestCreate(t *testing.T) {
	FamilyService := setup()
	t.Run("test case 1, success create family members", func(t *testing.T) {
		domain := FamilyMember{
			Name:      "Agus",
			Birthday:  time.Now(),
			Nik:       "2313131",
			Gender:    "Male",
			Age:       15,
			Handphone: "087762827361",
			CitizenId: 1,
		}

		expectedReturn := records.FamilyMember{
			Id:        1,
			Name:      "Agus",
			Birthday:  time.Now(),
			Nik:       "2313131",
			Gender:    "Male",
			Age:       15,
			Handphone: "087762827361",
			CitizenId: 1,
		}
		familyRepository.On("Create", mock.Anything, mock.Anything).Return(expectedReturn, nil).Once()
		create, err := FamilyService.Create(context.Background(), 1, domain)
		assert.Nil(t, err)
		assert.Equal(t, create.Name, expectedReturn.Name)
	})

	t.Run("test case 2, trigger validate", func(t *testing.T) {
		domain := FamilyMember{
			Name:      "",
			Birthday:  time.Now(),
			Nik:       "2313131",
			Gender:    "Male",
			Age:       15,
			Handphone: "087762827361",
			CitizenId: 1,
		}

		_, err := FamilyService.Create(context.Background(), 1, domain)
		assert.Error(t, err)
	})
}

func TestGetFamilyById(t *testing.T) {
	FamilyService := setup()
	t.Run("test case 1, success get family by ID", func(t *testing.T) {
		expectedReturn := records.FamilyMember{
			Id:        1,
			Name:      "Agus",
			Birthday:  time.Now(),
			Nik:       "2313131",
			Gender:    "Male",
			Age:       15,
			Handphone: "087762827361",
			CitizenId: 1,
		}
		familyRepository.On("GetFamilyById", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturn, nil).Once()

		id, err := FamilyService.GetFamilyById(context.Background(), 1)
		assert.Nil(t, err)
		assert.Equal(t, id.Name, expectedReturn.Name)
	})
	t.Run("test case 2, failed id not found", func(t *testing.T) {
		err2 := errors.New("data not found")
		familyRepository.On("GetFamilyById", mock.Anything, mock.AnythingOfType("int")).Return(records.FamilyMember{}, err2).Once()

		_, err := FamilyService.GetFamilyById(context.Background(), 1)
		assert.Equal(t, err, err2)
	})
}

func TestGetCitizenOwnFamily(t *testing.T) {
	FamilyService := setup()
	t.Run("test case 1, success get family owned by a citizen", func(t *testing.T) {
		expectedReturn := []records.FamilyMember{
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
		familyRepository.On("GetCitizenOwnFamily", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturn, nil).Once()
		family, err := FamilyService.GetCitizenOwnFamily(context.Background(), 2)
		assert.Nil(t, err)
		assert.Equal(t, family[0].Name, expectedReturn[0].Name)
	})

	t.Run("test case 2, failed get data that owned by a citizen id not found", func(t *testing.T) {
		err2 := errors.New("data not found")
		familyRepository.On("GetCitizenOwnFamily", mock.Anything, mock.AnythingOfType("int")).Return([]records.FamilyMember{}, err2).Once()
		_, err := FamilyService.GetCitizenOwnFamily(context.Background(), 2)
		assert.Equal(t, err, err2)
	})
}

func TestUpdate(t *testing.T) {
	FamilyService := setup()
	t.Run("test case 1, success update data", func(t *testing.T) {
		ctzId := 1
		domain := FamilyMember{
			Id:        1,
			Name:      "Agus",
			Birthday:  time.Now(),
			Nik:       "123123asd",
			Gender:    "Male",
			Age:       15,
			Handphone: "087523123123",
		}
		expectedReturn := records.FamilyMember{
			Id:        1,
			Name:      "Agus",
			Birthday:  time.Now(),
			Nik:       "123123asd",
			Gender:    "Male",
			Age:       15,
			Handphone: "087523123123",
			CitizenId: 1,
		}
		familyRepository.On("GetFamilyById", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturn, nil).Once()
		familyRepository.On("Update", mock.Anything, mock.AnythingOfType("int"), mock.Anything).Return(expectedReturn, nil).Once()

		update, err := FamilyService.Update(context.Background(), ctzId, domain.Id, domain)
		assert.Nil(t, err)
		assert.Equal(t, update.Name, expectedReturn.Name)
	})

	t.Run("test case 2, data with that id not found", func(t *testing.T) {
		ctzId := 1
		domain := FamilyMember{
			Id:        1,
			Name:      "Agus",
			Birthday:  time.Now(),
			Nik:       "123123asd",
			Gender:    "Male",
			Age:       15,
			Handphone: "087523123123",
		}
		err2 := errors.New("data not found")
		familyRepository.On("GetFamilyById", mock.Anything, mock.AnythingOfType("int")).Return(records.FamilyMember{}, err2).Once()
		_, err := FamilyService.Update(context.Background(), ctzId, domain.Id, domain)
		assert.Equal(t, err, err2)
	})
}

func TestDelete(t *testing.T) {
	FamilyService := setup()
	t.Run("test case 1, success delete family memeber", func(t *testing.T) {
		citizenId := 1
		familyId := 1
		expectedReturn := records.FamilyMember{
			Id:        1,
			Name:      "Agus",
			Birthday:  time.Now(),
			Nik:       "123123asd",
			Gender:    "Male",
			Age:       15,
			Handphone: "087523123123",
			CitizenId: 1,
		}
		familyRepository.On("GetFamilyById", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturn, nil).Once()
		familyRepository.On("Delete", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(records.FamilyMember{}, nil).Once()

		s, err := FamilyService.Delete(context.Background(), familyId, citizenId)
		assert.Nil(t, err)
		assert.Equal(t, s, "success delete data")

	})
}
