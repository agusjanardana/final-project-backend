package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

//berisi base response, 200, 400, 500 dst....

type BaseReponses struct {
	Meta struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func NewSuccessResponse(c echo.Context, data interface{}) error {
	response := BaseReponses{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = "success"
	response.Data = data

	return c.JSON(http.StatusOK, response)
}

func ForbiddenRequest(c echo.Context, error int, err error) error {
	response := BaseReponses{}
	response.Meta.Status = error
	response.Meta.Message = "forbidden access"

	return c.JSON(http.StatusForbidden, response)
}

func InternalServerError(c echo.Context, error int, err error) error {
	response := BaseReponses{}
	response.Meta.Status = error
	response.Meta.Message = "internal server error"

	return c.JSON(http.StatusInternalServerError, response)
}

func BadRequestResponse(c echo.Context, error int, err error) error {
	response := BaseReponses{}
	response.Meta.Status = error
	response.Meta.Message = string(err.Error())

	return c.JSON(http.StatusBadRequest, response)
}
