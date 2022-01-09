package CitizenService

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"log"
	"time"
	"vaccine-app-be/app/middleware"
	"vaccine-app-be/drivers/records"
	"vaccine-app-be/drivers/repository/CitizenRepository"
	"vaccine-app-be/drivers/repository/FamilyRepository"
	"vaccine-app-be/utilities"
)

type CitizenServiceImpl struct {
	CitizenRepository CitizenRepository.CitizenRepository
	jwtAuth           *middleware.ConfigJWT
	FamilyRepository  FamilyRepository.FamilyRepository
}

func NewCitizenService(CitizenRepository CitizenRepository.CitizenRepository, jwtAuth *middleware.ConfigJWT, FamilyRepository FamilyRepository.FamilyRepository) CitizenService {
	return &CitizenServiceImpl{
		CitizenRepository: CitizenRepository,
		jwtAuth:           jwtAuth,
		FamilyRepository:  FamilyRepository,
	}
}

func (service *CitizenServiceImpl) Register(ctx context.Context, citizen Citizen) (Citizen, error) {
	err := citizen.Validate()
	if err != nil {
		return citizen, err
	}
	//checking if emails is already used
	byEmail, err := service.CitizenRepository.FindByEmail(ctx, citizen.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
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
	//handle auto registered family members
	data := resp.ToRecordFamily()
	log.Println(data)
	_, err = service.FamilyRepository.Create(ctx, data)
	if err != nil {
		return Citizen{}, err
	}
	return resp, nil
}

func (service *CitizenServiceImpl) Login(ctx context.Context, email, password string) (string, error) {
	if len(email) == 0 || len(password) == 0 {
		return "", errors.New("email or password blank")
	}

	byEmail, err := service.CitizenRepository.FindByEmail(ctx, email)
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

func (service *CitizenServiceImpl) Update(ctx context.Context, userId int, birthDay time.Time, address string) (Citizen, error) {
	update, err := service.CitizenRepository.Update(ctx, userId, birthDay, address)
	if err != nil {
		return Citizen{}, err
	}
	entity := Citizen{}
	copier.Copy(&entity, &update)
	record := records.FamilyMember{Birthday: birthDay, Address: address}
	data, err := service.FamilyRepository.GetCitizenOwnFamily(ctx, userId)
	if err != nil {
		return Citizen{}, err
	}
	_, err = service.FamilyRepository.Update(ctx, data[0].Id, record)
	if err != nil {
		return Citizen{}, err
	}
	return entity, nil
}

func (service *CitizenServiceImpl) CitizenFindById(ctx context.Context, userId int) (Citizen, error) {
	citizenData, err := service.CitizenRepository.FindById(ctx, userId)
	if err != nil {
		return Citizen{}, err
	}

	entityCitizenRespond := Citizen{}
	copier.Copy(&entityCitizenRespond, &citizenData)

	return entityCitizenRespond, nil
}
