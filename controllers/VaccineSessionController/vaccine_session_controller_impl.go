package VaccineSessionController

import (
	"errors"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"vaccine-app-be/app/middleware"
	"vaccine-app-be/controllers"
	"vaccine-app-be/controllers/VaccineSessionController/web"
	"vaccine-app-be/services/VaccineSessionService"
)

type VaccineSessionControllerImpl struct {
	vaccineSessionService VaccineSessionService.VaccineSessionService
}

func NewVaccineSessionController(vaccineSessionService VaccineSessionService.VaccineSessionService) VaccineSessionController {
	return &VaccineSessionControllerImpl{vaccineSessionService: vaccineSessionService}
}

func (controller *VaccineSessionControllerImpl) CreateSession(c echo.Context) error {
	ctx := c.Request().Context()
	ctxRole := middleware.GetUserRoles(c)
	ctxId := middleware.GetUserId(c)
	if ctxRole != "ADMIN" {
		return controllers.ForbiddenRequest(c, http.StatusForbidden, errors.New("doesn't have access"))
	}

	req := web.VaccineSessionCreateRequest{}
	req.HealthFacilitatorId = ctxId

	err := c.Bind(&req)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}

	entityService := VaccineSessionService.VaccineSession{}
	copier.Copy(&entityService, &req)
	session, err := controller.vaccineSessionService.CreateSession(ctx, entityService)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusNotAcceptable,err)
	}
	response := web.VaccineSessionCreateResponse{}
	copier.Copy(&response, &session)

	return controllers.NewSuccessResponse(c, response)
}

func (controller *VaccineSessionControllerImpl) GetSessionById(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}
	data, err := controller.vaccineSessionService.GetSessionById(ctx, atoi)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusInternalServerError, err)
	}

	response := web.VaccineSessionFindByIdResponse{}
	copier.Copy(&response, &data)
	return controllers.NewSuccessResponse(c, response)
}

func (controller *VaccineSessionControllerImpl) GetSessionOwnedByHf(c echo.Context) error {
	ctx := c.Request().Context()
	//ctxId := middleware.GetUserId(c)
	idParam := c.Param("id")
	dataToInteger, err := strconv.Atoi(idParam)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}

	data, err := controller.vaccineSessionService.GetSessionOwnedByHf(ctx, dataToInteger)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusInternalServerError, err)
	}

	var response []web.VaccineSessionCreateResponse
	copier.Copy(&response, &data)
	return controllers.NewSuccessResponse(c, response)
}

func (controller *VaccineSessionControllerImpl) DeleteSession(c echo.Context) error {
	ctx := c.Request().Context()
	ctxId := middleware.GetUserId(c)
	id := c.Param("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}
	_, err = controller.vaccineSessionService.DeleteSession(ctx, ctxId, atoi)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, echo.Map{
		"message": "success delete",
	})
}

func (controller *VaccineSessionControllerImpl) UpdateSession(c echo.Context) error {
	ctx := c.Request().Context()
	ctxId := middleware.GetUserId(c)
	id := c.Param("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}
	req := web.VaccineSessionCreateRequest{}
	err = c.Bind(&req)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}

	entityService := VaccineSessionService.VaccineSession{}
	copier.Copy(&entityService, req)
	session, err := controller.vaccineSessionService.UpdateSession(ctx, ctxId, atoi, entityService)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusInternalServerError, err)
	}

	response := web.VaccineSessionCreateResponse{}
	copier.Copy(&response, &session)

	return controllers.NewSuccessResponse(c, response)
}

func (controller *VaccineSessionControllerImpl) GetAllVaccineSession(c echo.Context) error {
	ctx := c.Request().Context()
	session, err := controller.vaccineSessionService.GetAllVaccineSession(ctx)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusInternalServerError, err)
	}
	var response []web.VaccineSessionCreateResponse
	copier.Copy(&response, &session)

	return controllers.NewSuccessResponse(c, response)
}

func (controller *VaccineSessionControllerImpl) GetCitizenAndFamilySelectedSession(c echo.Context) error {
	ctx := c.Request().Context()
	ctxCitizenId := middleware.GetUserId(c)

	sessionData, err := controller.vaccineSessionService.GetCitizenAndFamilySelectedSession(ctx, ctxCitizenId)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusInternalServerError, err)
	}
	var response []web.VaccineSessionCreateResponse
	copier.Copy(&response, &sessionData)

	citizenIdString := strconv.Itoa(ctxCitizenId)
	return controllers.NewSuccessResponse(c, echo.Map{
		"message": "Citizen id " + citizenIdString + " with his family selected this session",
		"session": response,
	})
}
