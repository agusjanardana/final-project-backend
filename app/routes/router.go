package routes

import (
	"github.com/labstack/echo/v4"
	"vaccine-app-be/controllers/CitizenController"
)

type ControllerList struct {
	CitizenController CitizenController.CitizenController
}

func (c1 *ControllerList) Registration(e *echo.Echo) {
	apiV1 := e.Group("/api/v1")

	//	CITIZEN THINGS
	apiV1.POST("/citizen/registers", c1.CitizenController.Register)
	apiV1.POST("/citizen/logins", c1.CitizenController.Login)
}
