package VaccineController

import (
	"github.com/labstack/echo/v4"
)

type VaccineController interface {
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	FindVaccineById(c echo.Context) error
	FindVaccineOwnedByHF(c echo.Context) error
}
