package HfService

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"vaccine-app-be/app/middleware"
	"vaccine-app-be/drivers/records"
	"vaccine-app-be/drivers/repository/HealthRepository"
	"vaccine-app-be/utilities"
)

type HealthServiceImpl struct {
	HealthRepository HealthRepository.HealthRepository
	jwtAuth          *middleware.ConfigJWT
}

func NewHealthService(HealthRepository HealthRepository.HealthRepository, jwtAuth *middleware.ConfigJWT) HealthService {
	return &HealthServiceImpl{
		HealthRepository: HealthRepository,
		jwtAuth:          jwtAuth,
	}
}

func (service *HealthServiceImpl) Register(ctx context.Context, healthF HealthFacilitator) (HealthFacilitator, error) {
	err := healthF.Validate()
	if err != nil {
		return healthF, err
	}

	email, err := service.HealthRepository.FindByEmail(ctx, healthF.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return HealthFacilitator{}, err
	}

	if email.Email == healthF.Email {
		return HealthFacilitator{}, errors.New("email already used")
	}

	password, err := utilities.HashPassword(healthF.Password)
	if err != nil {
		return HealthFacilitator{}, err
	}
	healthF.Password = string(password)

	entityHf := new(records.HealthFacilitator)
	copier.Copy(entityHf, &healthF)
	register, err := service.HealthRepository.Register(ctx, *entityHf)
	if err != nil {
		return HealthFacilitator{}, err
	}

	respond := HealthFacilitator{}
	copier.Copy(&respond, &register)

	return respond, nil
}

func (service *HealthServiceImpl) Login(ctx context.Context, email, password string) (string, error) {
	if len(email) == 0 || len(password) == 0 {
		return "", errors.New("email or password blank")
	}

	byEmail, err := service.HealthRepository.FindByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	matchPassword := utilities.CheckPasswordHash(password, byEmail.Password)
	if !matchPassword {
		return "", errors.New("password doesn't match")
	}
	jwt := service.jwtAuth.GenerateToken(byEmail.Id, byEmail.Name, byEmail.Role)
	return jwt, nil
}

func (service *HealthServiceImpl) GetAllHealthFacilitator(ctx context.Context) ([]HealthFacilitator, error) {
	facilitator, err := service.HealthRepository.GetAllHealthFacilitator(ctx)
	if err != nil {
		return nil, err
	}

	var response []HealthFacilitator
	copier.Copy(&response, &facilitator)

	return response, nil
}

func (service *HealthServiceImpl) FindById(ctx context.Context, hfId int) (HealthFacilitator, error) {
	dataFacilitators, err := service.HealthRepository.FindById(ctx, hfId)
	if err != nil {
		return HealthFacilitator{}, err
	}
	response := HealthFacilitator{}
	copier.Copy(&response, &dataFacilitators)

	return response, nil
}

func (service *HealthServiceImpl) Update(ctx context.Context, hfId int, domain HealthFacilitator) (HealthFacilitator, error) {
	dataFindById, err := service.HealthRepository.FindById(ctx, hfId)
	if err != nil {
		return HealthFacilitator{}, err
	}

	if dataFindById.Id == hfId {
		entityRepo := records.HealthFacilitator{}
		copier.Copy(&entityRepo, &domain)
		dataRepo, err := service.HealthRepository.Update(ctx, hfId, entityRepo)
		if err != nil {
			return HealthFacilitator{}, err
		}
		response := HealthFacilitator{}
		copier.Copy(&response, &dataRepo)

		return response, nil
	}

	return HealthFacilitator{}, err
}
