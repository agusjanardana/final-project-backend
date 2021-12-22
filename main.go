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
	"vaccine-app-be/controllers/HealthController"
	"vaccine-app-be/drivers/repository/CitizenRepository"
	"vaccine-app-be/drivers/repository/HealthRepository"
	"vaccine-app-be/exceptions"
	"vaccine-app-be/services/CitizenService"
	"vaccine-app-be/services/HfService"
)

func main(){
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
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))


	cfg:= config.New()
	mysqlClient := mysql.New(cfg)
	defer mysqlClient.Close()

	//setup JWT
	EXPIRED, _ := strconv.Atoi(os.Getenv("JWT_EXPIRED"))

	configJWT := middleware2.ConfigJWT{
		SecretJWT: os.Getenv("JWT_SECRET"),
		ExpiredIn: EXPIRED,
	}

	//citizen
	citizenRepo := CitizenRepository.NewCitizenRepository(mysqlClient)
	citizenServ := CitizenService.NewCitizenService(citizenRepo, &configJWT)
	citizenCtrl := CitizenController.NewCitizenController(citizenServ)

	//HealthFa
	healthRepo := HealthRepository.NewHealthRepository(mysqlClient)
	healthServ := HfService.NewHealthService(healthRepo, &configJWT)
	healthCtrl := HealthController.NewHealthFacilitatorsController(healthServ)

	routesInit := routes.ControllerList{
		CitizenController: citizenCtrl,
		HealthController: healthCtrl,
	}
	routesInit.Registration(e)
	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}



