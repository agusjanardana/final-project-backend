package VaccineController

import (
	"errors"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"vaccine-app-be/app/middleware"
	"vaccine-app-be/controllers"
	"vaccine-app-be/controllers/VaccineController/web"
	"vaccine-app-be/services/VaccineService"
)

type VaccineControllerImpl struct {
	vaccineService VaccineService.VaccineService
}

func NewVaccineController(vaccineService VaccineService.VaccineService) VaccineController {
	return &VaccineControllerImpl{vaccineService: vaccineService}
}

func (controller *VaccineControllerImpl) Create(c echo.Context) error {
	ctx := c.Request().Context()

	ctxHfId := middleware.GetUserId(c)
	ctxRole := middleware.GetUserRoles(c)
	req := web.VaccineCreateRequest{}
	err := c.Bind(&req)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}
	req.HealthFacilitatorId = ctxHfId
	if ctxRole != "ADMIN" {
		return controllers.ForbiddenRequest(c, http.StatusForbidden, errors.New("doesn't have access"))
	}

	entityVaccine := VaccineService.Vaccine{}
	copier.Copy(&entityVaccine, req)
	create, err := controller.vaccineService.Create(ctx, entityVaccine)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}

	response := web.VaccineCreateResponse{}
	copier.Copy(&response, &create)

	return controllers.NewSuccessResponse(c, response)
}

func (controller *VaccineControllerImpl) Update(c echo.Context) error {
	ctx := c.Request().Context()

	vaccineId := c.Param("id")
	atoi, err := strconv.Atoi(vaccineId)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}

	ctxHfId := middleware.GetUserId(c)
	ctxRole := middleware.GetUserRoles(c)
	req := web.VaccineUpdateRequest{}
	err = c.Bind(&req)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}
	if ctxRole != "ADMIN" {
		return controllers.ForbiddenRequest(c, http.StatusForbidden, errors.New("doesn't have access"))
	}

	entityVaccine := VaccineService.Vaccine{}
	copier.Copy(&entityVaccine, req)
	create, err := controller.vaccineService.Update(ctx, ctxHfId, atoi, entityVaccine)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}
	response := web.VaccineUpdateResponse{}
	copier.Copy(&response, &create)

	return controllers.NewSuccessResponse(c, response)
}

func (controller *VaccineControllerImpl) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	vaccineId := c.Param("id")
	atoi, err := strconv.Atoi(vaccineId)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}

	ctxHfId := middleware.GetUserId(c)
	ctxRole := middleware.GetUserRoles(c)
	if ctxRole != "ADMIN" {
		return controllers.ForbiddenRequest(c, http.StatusForbidden, errors.New("doesn't have access"))
	}

	_, err = controller.vaccineService.Delete(ctx, ctxHfId, atoi)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, echo.Map{
		"message": "success delete",
	})
}

func (controller *VaccineControllerImpl) FindVaccineById(c echo.Context) error {
	ctx := c.Request().Context()
	vaccineId := c.Param("id")
	ctxRole := middleware.GetUserRoles(c)
	if ctxRole != "ADMIN" {
		return controllers.ForbiddenRequest(c, http.StatusForbidden, errors.New("doesn't have access"))
	}
	atoi, err := strconv.Atoi(vaccineId)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}
	id, err := controller.vaccineService.FindVaccineById(ctx, atoi)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}

	response := web.VaccineFindByIdResponse{}
	copier.Copy(&response, &id)
	return controllers.NewSuccessResponse(c, response)
}

func (controller *VaccineControllerImpl) FindVaccineOwnedByHF(c echo.Context) error {
	ctx := c.Request().Context()
	ctxRole := middleware.GetUserRoles(c)
	//ctxHfId := middleware.GetUserId(c)
	if ctxRole != "ADMIN" {
		return controllers.ForbiddenRequest(c, http.StatusForbidden, errors.New("doesn't have access"))
	}

	idParam := c.Param("id")
	dataToInteger, err := strconv.Atoi(idParam)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}

	hf, err := controller.vaccineService.FindVaccineOwnedByHF(ctx, dataToInteger)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}

	var response []web.VaccineFindByIdResponse
	copier.Copy(&response, &hf)

	return controllers.NewSuccessResponse(c, response)
}
