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
	"vaccine-app-be/drivers/repository/VaccineSessionDetailRepository"
	"vaccine-app-be/drivers/repository/VaccineSessionRepository"
	"vaccine-app-be/utilities"
)

type CitizenServiceImpl struct {
	CitizenRepository        CitizenRepository.CitizenRepository
	jwtAuth                  *middleware.ConfigJWT
	FamilyRepository         FamilyRepository.FamilyRepository
	vaccineSessionRepository VaccineSessionRepository.VaccineSessionRepository
	sessionDetail            VaccineSessionDetailRepository.VaccineSessionDetail
}

func NewCitizenService(CitizenRepository CitizenRepository.CitizenRepository, jwtAuth *middleware.ConfigJWT, FamilyRepository FamilyRepository.FamilyRepository, vaccineSessionRepository VaccineSessionRepository.VaccineSessionRepository, sessionDetail VaccineSessionDetailRepository.VaccineSessionDetail,
) CitizenService {
	return &CitizenServiceImpl{
		CitizenRepository:        CitizenRepository,
		jwtAuth:                  jwtAuth,
		FamilyRepository:         FamilyRepository,
		vaccineSessionRepository: vaccineSessionRepository,
		sessionDetail:            sessionDetail,
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

func (service *CitizenServiceImpl) GetCitizenRelationWithHealthFacilitators(ctx context.Context, hfId int) ([]Citizen, error) {
	dataSessionHf, err := service.vaccineSessionRepository.FindSessionOwnedByHf(ctx, hfId)
	if err != nil {
		return nil, err
	}

	var idFamily []int
	for i := 0; i < len(dataSessionHf); i++ {
		dataDetail, err := service.sessionDetail.GetDetailBySessionId(ctx, dataSessionHf[i].Id)
		if err != nil {
			return nil, err
		}
		idFamily = append(idFamily, dataDetail[i].FamilyMemberId)
	}

	var idCitizen []int
	for i := 0; i < len(idFamily); i++ {
		dataFamily, err := service.FamilyRepository.GetFamilyById(ctx, idFamily[i])
		if err != nil {
			return nil, err
		}

		skip := false
		for _, u := range idCitizen {
			if dataFamily.CitizenId == u{
				skip = true
				break
			}
		}
		if !skip{
			idCitizen = append(idCitizen, dataFamily.CitizenId)
		}
	}

	data := make([]Citizen, len(idCitizen)-1)
	for i := 0;i < len(idCitizen); i++ {
		dataCitizenRepo, err := service.CitizenRepository.FindById(ctx, idCitizen[i])
		if err != nil {
			return nil, err
		}
		response := Citizen{}
		copier.Copy(&response, &dataCitizenRepo)
		data = append(data, response)
	}
	return data, nil
}
