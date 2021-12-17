package CitizenController

import (
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"net/http"
	"vaccine-app-be/controllers"
	"vaccine-app-be/controllers/CitizenController/web"
	"vaccine-app-be/services/CitizenService"
)

type CitizenControllerImpl struct {
	citizenService CitizenService.CitizenService
}

func NewCitizenController(citizenService CitizenService.CitizenService) CitizenController {
	return &CitizenControllerImpl{citizenService: citizenService}
}
func (citizenCtrl *CitizenControllerImpl) Register(c echo.Context) error {
	ctx := c.Request().Context()

	req := web.RequestRegister{}
	err := c.Bind(&req)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}
	entityCitizen := new(CitizenService.Citizen)
	copier.Copy(&entityCitizen, &req)
	registered, err := citizenCtrl.citizenService.Register(ctx, *entityCitizen)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}

	responses := web.ResponseRegister{}
	copier.Copy(&responses, &registered)

	return controllers.NewSuccessResponse(c, responses)
}

func (citizenCtrl *CitizenControllerImpl) Login(c echo.Context) error {
	ctx := c.Request().Context()

	req := web.CitizenLoginRequest{}
	err := c.Bind(&req)
	if err != nil{
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}

	login, err := citizenCtrl.citizenService.Login(ctx, req.Email, req.Password)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, echo.Map{
		"token" : login,
	})
}
