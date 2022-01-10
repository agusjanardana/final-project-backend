package VaccineSessionController

import "github.com/labstack/echo/v4"

type VaccineSessionController interface {
	CreateSession(c echo.Context) error
	GetSessionById(c echo.Context) error
	GetSessionOwnedByHf(c echo.Context) error
	DeleteSession(c echo.Context) error
	UpdateSession(c echo.Context) error
	GetAllVaccineSession(c echo.Context) error
	GetCitizenAndFamilySelectedSession(c echo.Context) error
}
