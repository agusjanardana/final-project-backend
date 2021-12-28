package FamilyController

import (
	"github.com/labstack/echo/v4"
)

type FamilyController interface {
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	GetFamilyById(c echo.Context) error
	GetCitizenOwnFamily(c echo.Context) error
}
