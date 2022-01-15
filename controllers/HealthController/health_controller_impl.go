package HealthController

import (
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"vaccine-app-be/controllers"
	"vaccine-app-be/controllers/HealthController/web"
	"vaccine-app-be/services/FamilyService"
	"vaccine-app-be/services/HfService"
)

type HealthFacilitatorCtrlImpl struct {
	healthService HfService.HealthService
	familyService FamilyService.FamilyService
}

func NewHealthFacilitatorsController(healthFacilitator HfService.HealthService, familyService FamilyService.FamilyService) HealthFacilitatorController {
	return &HealthFacilitatorCtrlImpl{healthService: healthFacilitator, familyService: familyService}
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
		return controllers.InternalServerError(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, echo.Map{
		"token": login,
	})
}

func (controller *HealthFacilitatorCtrlImpl) GetAllHealthFacilitator(c echo.Context) error {
	ctx := c.Request().Context()

	facilitator, err := controller.healthService.GetAllHealthFacilitator(ctx)
	if err != nil {
		return controllers.InternalServerError(c, http.StatusInternalServerError, err)
	}

	var response []web.HealthFacilitator
	copier.Copy(&response, &facilitator)
	return controllers.NewSuccessResponse(c, response)
}

func (controller *HealthFacilitatorCtrlImpl) FindById(c echo.Context) error {
	ctx := c.Request().Context()
	requestId := c.Param("id")
	dataToInteger, err := strconv.Atoi(requestId)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}
	dataService, err := controller.healthService.FindById(ctx, dataToInteger)
	if err != nil {
		return controllers.InternalServerError(c, http.StatusInternalServerError, err)
	}

	response := web.HealthFacilitator{}
	copier.Copy(&response, &dataService)

	return controllers.NewSuccessResponse(c, response)
}

func (controller *HealthFacilitatorCtrlImpl) HealthUpdateFamilyMemberStatus(c echo.Context) error {
	ctx := c.Request().Context()
	idParam := c.Param("id")
	dataToInteger, err := strconv.Atoi(idParam)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}

	request := web.FamilyMemberUpdateStatusReq{}
	err = c.Bind(&request)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}

	serviceData := FamilyService.FamilyMember{}
	copier.Copy(&serviceData, &request)

	familyData, err := controller.familyService.HfUpdateStatusFamily(ctx, dataToInteger, serviceData)
	if err != nil {
		return controllers.InternalServerError(c, http.StatusInternalServerError, err)
	}

	response := web.FamilyMemberResponse{}
	copier.Copy(&response, &familyData)

	return controllers.NewSuccessResponse(c, response)
}

func (controller *HealthFacilitatorCtrlImpl) Update(c echo.Context) error {
	ctx := c.Request().Context()
	idParam := c.Param("id")
	dataToInteger, err := strconv.Atoi(idParam)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}
	request := web.HealthFaUpdateReq{}
	err = c.Bind(&request)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}

	entityHealthF := HfService.HealthFacilitator{}
	copier.Copy(&entityHealthF, &request)
	updateData, err := controller.healthService.Update(ctx, dataToInteger, entityHealthF)
	if err != nil {
		return err
	}

	response := web.HealthFaUpdateRes{}
	copier.Copy(&response, &updateData)

	return controllers.NewSuccessResponse(c, response)
}
