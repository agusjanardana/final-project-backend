package exceptions

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func ErrorHandler(err error, c echo.Context) {
	log.Println(err.Error())
	c.JSON(http.StatusInternalServerError, echo.Map{"errors": err.Error()})
}