package VaccineDetailController

import (
	"github.com/labstack/echo/v4"
)

type VaccineDetailController interface {
	CitizenChooseSession(c echo.Context) error
	GetDetailBySessionId(c echo.Context) error
	GetDetailById(c echo.Context) error
	GetDetailByFamilyId(c echo.Context) error
}