package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"os"
	"strconv"
	"vaccine-app-be/app/config"
	"vaccine-app-be/app/config/mysql"
	middleware2 "vaccine-app-be/app/middleware"
	"vaccine-app-be/app/routes"
	"vaccine-app-be/controllers/CitizenController"
	"vaccine-app-be/controllers/FamilyController"
	"vaccine-app-be/controllers/HealthController"
	"vaccine-app-be/controllers/VaccineController"
	"vaccine-app-be/controllers/VaccineDetailController"
	"vaccine-app-be/controllers/VaccineSessionController"
	"vaccine-app-be/drivers/repository/CitizenRepository"
	"vaccine-app-be/drivers/repository/FamilyRepository"
	"vaccine-app-be/drivers/repository/HealthRepository"
	"vaccine-app-be/drivers/repository/VaccineRepository"
	"vaccine-app-be/drivers/repository/VaccineSessionDetailRepository"
	"vaccine-app-be/drivers/repository/VaccineSessionRepository"
	"vaccine-app-be/exceptions"
	"vaccine-app-be/services/CitizenService"
	"vaccine-app-be/services/FamilyService"
	"vaccine-app-be/services/HfService"
	"vaccine-app-be/services/SessionDetailService"
	"vaccine-app-be/services/VaccineService"
	"vaccine-app-be/services/VaccineSessionService"
)

func main() {
	_ = godotenv.Load()
	e := echo.New()
	//handle error midleware
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:       1 << 10, // 1 KB
		DisableStackAll: true,
	}))
	e.HTTPErrorHandler = exceptions.ErrorHandler

	//handle CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderCookie},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	cfg := config.New()
	mysqlClient := mysql.New(cfg)
	defer mysqlClient.Close()

	//setup JWT
	EXPIRED, _ := strconv.Atoi(os.Getenv("JWT_EXPIRED"))

	configJWT := middleware2.ConfigJWT{
		SecretJWT: os.Getenv("JWT_SECRET"),
		ExpiredIn: EXPIRED,
	}

	//Family
	familyRepo := FamilyRepository.NewFamilyRepository(mysqlClient)
	familyServ := FamilyService.NewFamilyService(familyRepo)
	familyCtrl := FamilyController.NewFamilyControllerImpl(familyServ)

	//HealthFa
	healthRepo := HealthRepository.NewHealthRepository(mysqlClient)
	healthServ := HfService.NewHealthService(healthRepo, &configJWT)
	healthCtrl := HealthController.NewHealthFacilitatorsController(healthServ, familyServ)

	//vaccine
	vaccineRepo := VaccineRepository.NewVaccineRepository(mysqlClient)
	vaccineServ := VaccineService.NewVaccineRepository(vaccineRepo)
	vaccineCtrl := VaccineController.NewVaccineController(vaccineServ)

	sessionRepo := VaccineSessionRepository.NewVaccineSessionRepository(mysqlClient)
	detailRepo := VaccineSessionDetailRepository.NewSessionDetail(mysqlClient)

	//citizen
	citizenRepo := CitizenRepository.NewCitizenRepository(mysqlClient)
	citizenServ := CitizenService.NewCitizenService(citizenRepo, &configJWT, familyRepo, sessionRepo, detailRepo)
	citizenCtrl := CitizenController.NewCitizenController(citizenServ)

	//session
	sessionServ := VaccineSessionService.NewSessionService(sessionRepo, vaccineRepo, familyRepo, citizenRepo)
	sessionCtrl := VaccineSessionController.NewVaccineSessionController(sessionServ, healthServ, vaccineServ)

	//detail
	detailServ := SessionDetailService.NewSessionDetail(detailRepo, familyRepo, sessionRepo)
	detailCtrl := VaccineDetailController.NewSessionDetailController(detailServ)



	routesInit := routes.ControllerList{
		JWTMiddleware:            configJWT.Init(),
		CitizenController:        citizenCtrl,
		HealthController:         healthCtrl,
		FamilyController:         familyCtrl,
		VaccineController:        vaccineCtrl,
		VaccineSessionController: sessionCtrl,
		VaccineDetailController:  detailCtrl,
	}
	routesInit.Registration(e)
	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
