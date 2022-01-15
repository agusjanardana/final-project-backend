package VaccineService

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"vaccine-app-be/drivers/records"
	"vaccine-app-be/drivers/repository/VaccineRepository/mocks"
)

var (
	vaccineRepository mocks.VaccineRepository
)

func setup() VaccineService {
	vaccineService := NewVaccineRepository(&vaccineRepository)

	return vaccineService
}

func TestCreate(t *testing.T) {
	VaccineService := setup()
	t.Run("test case 1, success create vaccine", func(t *testing.T) {
		domain := Vaccine{
			HealthFacilitatorId: 1,
			Name:                "Sinovac",
			Stock:               200,
		}

		expectedReturn := records.Vaccine{
			Id:                  1,
			HealthFacilitatorId: 1,
			Name:                "Sinovac",
			Stock:               200,
		}
		vaccineRepository.On("Create", mock.Anything, mock.Anything).Return(expectedReturn, nil).Once()

		CreatedData, err := VaccineService.Create(context.Background(), domain)
		assert.Nil(t, err)
		assert.Equal(t, CreatedData.Id, expectedReturn.Id)
	})

	t.Run("test case 2, validate request failed", func(t *testing.T) {
		domain := Vaccine{
			HealthFacilitatorId: 1,
			Stock:               200,
		}
		create, err := VaccineService.Create(context.Background(), domain)
		assert.Empty(t, create)
		assert.NotEmpty(t, err)
	})
}

func TestUpdate(t *testing.T) {
	VaccineService := setup()
	t.Run("test case 1, success update vaccine", func(t *testing.T) {
		vaccineId := 1
		HealthFacilitatorsId := 1
		domain := Vaccine{
			Name:  "Sinovac EDIT",
			Stock: 200,
		}

		expectedReturnFind := records.Vaccine{
			Id:                  1,
			HealthFacilitatorId: 1,
			Name:                "Sinovac",
			Stock:               200,
		}

		expectedReturnEdit := records.Vaccine{
			Id:                  1,
			HealthFacilitatorId: 1,
			Name:                "Sinovac EDIT",
			Stock:               200,
		}

		vaccineRepository.On("FindVaccineById", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturnFind, nil).Once()
		vaccineRepository.On("Update", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.Anything).Return(expectedReturnEdit, nil).Once()

		updateData, err := VaccineService.Update(context.Background(), HealthFacilitatorsId, vaccineId, domain)
		assert.Equal(t, updateData.Name, expectedReturnEdit.Name)
		assert.Nil(t, err)
	})

	t.Run("test case 2, vaccine not found", func(t *testing.T) {

		err2 := errors.New("data not found")
		vaccineRepository.On("FindVaccineById", mock.Anything, mock.AnythingOfType("int")).Return(records.Vaccine{}, err2).Once()

		domain := Vaccine{
			Name:  "Sinovac EDIT",
			Stock: 200,
		}

		updateData, err := VaccineService.Update(context.Background(), 1, 1, domain)

		assert.Equal(t, err, err2)
		assert.Empty(t, updateData)
	})
}

func TestDelete(t *testing.T) {
	VaccineService := setup()
	t.Run("test case 1, success delete data", func(t *testing.T) {
		Id := 1
		HfId := 1
		expectedReturnFind := records.Vaccine{
			Id:                  1,
			HealthFacilitatorId: 1,
			Name:                "Sinovac",
			Stock:               200,
		}
		vaccineRepository.On("FindVaccineById", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturnFind, nil).Once()
		vaccineRepository.On("Delete", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(records.Vaccine{}, nil).Once()

		dataDelete, err := VaccineService.Delete(context.Background(), HfId, Id)
		assert.Nil(t, err)
		assert.Equal(t, dataDelete, "success delete data")
	})

	t.Run("test case 2, not success delete id not found", func(t *testing.T) {
		Id := 1
		HfId := 1
		err2 := errors.New("data not found")
		vaccineRepository.On("FindVaccineById", mock.Anything, mock.AnythingOfType("int")).Return(records.Vaccine{}, err2).Once()

		dataDelete, err := VaccineService.Delete(context.Background(), HfId, Id)
		assert.Equal(t, err, err2)
		assert.Empty(t, dataDelete)
	})
}

func TestFindVaccineById(t *testing.T) {
	VaccineService := setup()
	t.Run("test case 1, success find by id", func(t *testing.T) {

		expectedReturnFind := records.Vaccine{
			Id:                  1,
			HealthFacilitatorId: 1,
			Name:                "Sinovac",
			Stock:               200,
		}

		vaccineRepository.On("FindVaccineById", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturnFind, nil).Once()
		data, err := VaccineService.FindVaccineById(context.Background(), 1)
		assert.Equal(t, data.Id, expectedReturnFind.Id)
		assert.Nil(t, err)
	})

	t.Run("test case 2, not success id not found", func(t *testing.T) {
		Id := 1
		err2 := errors.New("data not found")
		vaccineRepository.On("FindVaccineById", mock.Anything, mock.AnythingOfType("int")).Return(records.Vaccine{}, err2).Once()

		dataService, err := VaccineService.FindVaccineById(context.Background(), Id)
		assert.Equal(t, err, err2)
		assert.Empty(t, dataService)
	})
}

func TestFindVaccineOwnedByHF(t *testing.T) {
	VaccineService := setup()
	t.Run("test case 1, success find vaccine owned by a HF", func(t *testing.T) {
		expectedReturnFind := []records.Vaccine{
			{
				Id:                  1,
				HealthFacilitatorId: 1,
				Name:                "Sinovac",
				Stock:               200,
			},
			{
				Id:                  2,
				HealthFacilitatorId: 1,
				Name:                "Sinovac",
				Stock:               200,
			},
		}
		vaccineRepository.On("FindVaccineOwnedByHF", mock.Anything, mock.AnythingOfType("int")).Return(expectedReturnFind, nil).Once()

		dataFind, err := VaccineService.FindVaccineOwnedByHF(context.Background(), 1)
		assert.Nil(t, err)
		assert.Equal(t, dataFind[0].Id, expectedReturnFind[0].Id)
	})

	t.Run("test case 2, failed HF not have any vaccine", func(t *testing.T) {
		err2 := errors.New("data not found")
		vaccineRepository.On("FindVaccineOwnedByHF", mock.Anything, mock.AnythingOfType("int")).Return([]records.Vaccine{}, err2).Once()

		dataFind, err := VaccineService.FindVaccineOwnedByHF(context.Background(), 1)
		assert.Equal(t, err, err2)
		assert.Empty(t, dataFind)
	})

}
