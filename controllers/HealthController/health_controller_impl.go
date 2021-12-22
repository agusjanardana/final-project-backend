package HealthController

import (
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"net/http"
	"vaccine-app-be/controllers"
	"vaccine-app-be/controllers/HealthController/web"
	"vaccine-app-be/services/HfService"
)

type HealthFacilitatorCtrlImpl struct {
	healthService HfService.HealthService
}

func NewHealthFacilitatorsController(healthFacilitator HfService.HealthService) HealthFacilitatorController {
	return &HealthFacilitatorCtrlImpl{healthService: healthFacilitator}
}

func (controller *HealthFacilitatorCtrlImpl) Register(c echo.Context) error {
	ctx := c.Request().Context()

	req := web.HealthFaRegisterRequest{}

	err := c.Bind(&req)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}

	entityHealthF := HfService.HealthFacilitator{}
	copier.Copy(&entityHealthF, &req)
	register, err := controller.healthService.Register(ctx, entityHealthF)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}

	response := web.HealthFaRegisterResponse{}
	copier.Copy(&response, &register)

	return controllers.NewSuccessResponse(c, response)
}

func (controller *HealthFacilitatorCtrlImpl) Login(c echo.Context) error {
	ctx := c.Request().Context()
	req := web.HealthFaLoginRequest{}

	err := c.Bind(&req)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}

	login, err := controller.healthService.Login(ctx, req.Email, req.Password)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, echo.Map{
		"token" : login,
	})
}
