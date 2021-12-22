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
	if err != nil{
		return healthF, err
	}

	email, err := service.HealthRepository.FindByEmail(ctx, healthF.Email)
	if err != nil && !errors.Is(err,gorm.ErrRecordNotFound) {
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
	if len(email) == 0 || len(password) == 0{
		return "" , errors.New("email or password blank")
	}

	byEmail, err := service.HealthRepository.FindByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	matchPassword := utilities.CheckPasswordHash(password, byEmail.Password)
	if !matchPassword{
		return "", errors.New("password doesn't match")
	}
	jwt := service.jwtAuth.GenerateToken(byEmail.Id, byEmail.Name)
	return jwt, nil
}
