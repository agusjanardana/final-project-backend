package HealthController

import "github.com/labstack/echo/v4"

type HealthFacilitatorController interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
	GetAllHealthFacilitator(c echo.Context) error
	FindById(c echo.Context) error
}
