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
	"vaccine-app-be/drivers/repository/CitizenRepository"
	"vaccine-app-be/drivers/repository/FamilyRepository"
	"vaccine-app-be/drivers/repository/HealthRepository"
	"vaccine-app-be/drivers/repository/VaccineRepository"
	"vaccine-app-be/exceptions"
	"vaccine-app-be/services/CitizenService"
	"vaccine-app-be/services/FamilyService"
	"vaccine-app-be/services/HfService"
	"vaccine-app-be/services/VaccineService"
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

	//HealthFa
	healthRepo := HealthRepository.NewHealthRepository(mysqlClient)
	healthServ := HfService.NewHealthService(healthRepo, &configJWT)
	healthCtrl := HealthController.NewHealthFacilitatorsController(healthServ)

	//Family
	familyRepo := FamilyRepository.NewFamilyRepository(mysqlClient)
	familyServ := FamilyService.NewFamilyService(familyRepo)
	familyCtrl := FamilyController.NewFamilyControllerImpl(familyServ)

	//citizen
	citizenRepo := CitizenRepository.NewCitizenRepository(mysqlClient)
	citizenServ := CitizenService.NewCitizenService(citizenRepo, &configJWT, familyRepo)
	citizenCtrl := CitizenController.NewCitizenController(citizenServ)

	//vaccine
	vaccineRepo := VaccineRepository.NewVaccineRepository(mysqlClient)
	vaccineServ := VaccineService.NewVaccineRepository(vaccineRepo)
	vaccineCtrl := VaccineController.NewVaccineController(vaccineServ)

	routesInit := routes.ControllerList{
		JWTMiddleware:     configJWT.Init(),
		CitizenController: citizenCtrl,
		HealthController:  healthCtrl,
		FamilyController:  familyCtrl,
		VaccineController: vaccineCtrl,
	}
	routesInit.Registration(e)
	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
