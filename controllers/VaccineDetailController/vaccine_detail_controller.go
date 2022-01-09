package VaccineDetailController

import "github.com/labstack/echo/v4"

type VaccineDetailController interface {
	CitizenChooseSession(c echo.Context) error
}