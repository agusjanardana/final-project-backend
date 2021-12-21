package CitizenService

import (
	"context"
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"vaccine-app-be/app/middleware"
	"vaccine-app-be/drivers/records"
	"vaccine-app-be/drivers/repository/CitizenRepository"
	"vaccine-app-be/utilities"
)

type CitizenServiceImpl struct {
	CitizenRepository CitizenRepository.CitizenRepository
	jwtAuth           *middleware.ConfigJWT
}

func NewCitizenService(CitizenRepository CitizenRepository.CitizenRepository, jwtAuth *middleware.ConfigJWT) CitizenService {
	return &CitizenServiceImpl{
		CitizenRepository: CitizenRepository,
		jwtAuth:           jwtAuth,
	}
}

func (service *CitizenServiceImpl) Register(ctx context.Context, citizen Citizen) (Citizen, error) {
	err := citizen.Validate()
	if err != nil {
		return citizen, err
	}
	//checking if emails is already used
	byEmail, err := service.CitizenRepository.FindByEmail(ctx, citizen.Email)
	if err != nil {
		return citizen, err
	}

	if byEmail.Email == citizen.Email {
		return citizen, errors.New("email already used")
	}


	password, err := utilities.HashPassword(citizen.Password)
	if err != nil {
		return citizen, err
	}
	citizen.Password = string(password)

	entityCitizen := new(records.Citizen)
	copier.Copy(&entityCitizen, &citizen)

	registeredUser, err := service.CitizenRepository.Register(ctx, *entityCitizen)
	if err != nil {
		return citizen, err
	}
	resp := Citizen{}
	copier.Copy(&resp, &registeredUser)

	return resp, nil
}

func (service *CitizenServiceImpl) Login(ctx context.Context, email, password string) (string, error) {
	if len(email) == 0 || len(password) == 0{
		return "" , errors.New("email or password blank")
	}

	byEmail, err := service.CitizenRepository.FindByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	fmt.Println("test", password, "dan", byEmail.Email)
	matchPassword := utilities.CheckPasswordHash(password, byEmail.Password)

	if !matchPassword{
		return "", errors.New("password doesn't match")
	}

	jwt := service.jwtAuth.GenerateToken(byEmail.Id, byEmail.Name)

	return jwt, nil
}
