package CitizenController

import "github.com/labstack/echo/v4"

type CitizenController interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
	Update(c echo.Context) error
}
