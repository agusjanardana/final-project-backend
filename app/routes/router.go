package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"vaccine-app-be/controllers/CitizenController"
	"vaccine-app-be/controllers/FamilyController"
	"vaccine-app-be/controllers/HealthController"
)

type ControllerList struct {
	JWTMiddleware     middleware.JWTConfig
	CitizenController CitizenController.CitizenController
	HealthController  HealthController.HealthFacilitatorController
	FamilyController  FamilyController.FamilyController
}

func (c1 *ControllerList) Registration(e *echo.Echo) {
	apiV1 := e.Group("/api/v1")

	//	CITIZEN THINGS
	apiV1.POST("/citizen/registers", c1.CitizenController.Register)
	apiV1.POST("/citizen/logins", c1.CitizenController.Login)

	//	HEALTH FA THINGS
	apiV1.POST("/admin/registers", c1.HealthController.Register)
	apiV1.POST("/admin/logins", c1.HealthController.Login)

	//  FAMILY THINGS
	apiV1.GET("/families/:id", c1.FamilyController.GetFamilyById, middleware.JWTWithConfig(c1.JWTMiddleware))
	apiV1.POST("/families", c1.FamilyController.Create, middleware.JWTWithConfig(c1.JWTMiddleware))
	apiV1.PUT("/families/:id", c1.FamilyController.Update, middleware.JWTWithConfig(c1.JWTMiddleware))
	apiV1.DELETE("/families/:id", c1.FamilyController.Delete, middleware.JWTWithConfig(c1.JWTMiddleware))
	apiV1.GET("/family/citizens", c1.FamilyController.GetCitizenOwnFamily, middleware.JWTWithConfig(c1.JWTMiddleware))
}
