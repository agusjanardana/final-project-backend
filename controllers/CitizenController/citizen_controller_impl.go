package CitizenController

import (
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"net/http"
	"vaccine-app-be/app/middleware"
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
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}

	login, err := citizenCtrl.citizenService.Login(ctx, req.Email, req.Password)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, echo.Map{
		"token": login,
	})
}

func (citizenCtrl *CitizenControllerImpl) Update(c echo.Context) error {
	ctx := c.Request().Context()

	req := web.CitizenUpdateRequest{}
	ctxCitizenId := middleware.GetUserId(c)
	err := c.Bind(&req)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}
	update, err := citizenCtrl.citizenService.Update(ctx, ctxCitizenId, req.Birthday, req.Address)
	if err != nil {
		return controllers.InternalServerError(c, http.StatusInternalServerError, err)
	}
	entity := web.CitizenUpdateResponse{}
	copier.Copy(&entity, &update)

	return controllers.NewSuccessResponse(c, entity)
}

func (citizenCtrl *CitizenControllerImpl) FindCitizenById(c echo.Context) error {
	ctx := c.Request().Context()
	ctxCitizenId := middleware.GetUserId(c)

	citizenData, err := citizenCtrl.citizenService.CitizenFindById(ctx, ctxCitizenId)
	if err != nil {
		return controllers.InternalServerError(c, http.StatusInternalServerError, err)
	}
	response := web.RespondFind{}
	copier.Copy(&response, &citizenData)

	return controllers.NewSuccessResponse(c, response)
}
