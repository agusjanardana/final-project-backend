package routes

import (
	"github.com/labstack/echo/v4"
	"vaccine-app-be/controllers/CitizenController"
	"vaccine-app-be/controllers/HealthController"
)

type ControllerList struct {
	CitizenController CitizenController.CitizenController
	HealthController  HealthController.HealthFacilitatorController
}

func (c1 *ControllerList) Registration(e *echo.Echo) {
	apiV1 := e.Group("/api/v1")

	//	CITIZEN THINGS
	apiV1.POST("/citizen/registers", c1.CitizenController.Register)
	apiV1.POST("/citizen/logins", c1.CitizenController.Login)

	//	HEALTH FA THINGS
	apiV1.POST("/admin/registers", c1.HealthController.Register)
	apiV1.POST("/admin/logins", c1.HealthController.Login)
}
