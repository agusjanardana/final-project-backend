package HealthController

import "github.com/labstack/echo/v4"

type HealthFacilitatorController interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
}
