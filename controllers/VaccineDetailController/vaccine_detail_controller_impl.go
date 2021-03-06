package VaccineDetailController

import (
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"vaccine-app-be/app/middleware"
	"vaccine-app-be/controllers"
	"vaccine-app-be/controllers/VaccineDetailController/web"
	"vaccine-app-be/services/SessionDetailService"
)

type VaccineDetailControllerImpl struct {
	sessionDetail SessionDetailService.SessionDetail
}

func NewSessionDetailController(sessionDetail SessionDetailService.SessionDetail) VaccineDetailController {
	return &VaccineDetailControllerImpl{sessionDetail: sessionDetail}
}
func (controller *VaccineDetailControllerImpl) CitizenChooseSession(c echo.Context) error {
	ctx := c.Request().Context()
	ctxId := middleware.GetUserId(c)
	sessionId := c.Param("id")
	atoi, err := strconv.Atoi(sessionId)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}
	session, err := controller.sessionDetail.CitizenChooseSession(ctx, ctxId, atoi)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusNotAcceptable, err)
	}

	var response []web.SessionDetailDo
	copier.Copy(&response, &session)

	return controllers.NewSuccessResponse(c, response)
}

func (controller *VaccineDetailControllerImpl) GetDetailBySessionId(c echo.Context) error {
	ctx := c.Request().Context()
	idParamSessionId := c.Param("id")
	dataToInteger, err := strconv.Atoi(idParamSessionId)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}
	dataService, err := controller.sessionDetail.GetDetailBySessionId(ctx, dataToInteger)
	if err != nil {
		return controllers.InternalServerError(c, http.StatusInternalServerError, err)
	}
	var response []web.SessionDetailDo

	copier.Copy(&response, dataService)

	return controllers.NewSuccessResponse(c, response)
}

func (controller *VaccineDetailControllerImpl) GetDetailById(c echo.Context) error {
	ctx := c.Request().Context()
	idParam := c.Param("id")
	dataToInteger, err := strconv.Atoi(idParam)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}
	dataService, err := controller.sessionDetail.GetDetailById(ctx, dataToInteger)
	if err != nil {
		return controllers.InternalServerError(c, http.StatusInternalServerError, err)
	}

	var response web.SessionDetailDo
	copier.Copy(&response, &dataService)
	return controllers.NewSuccessResponse(c, response)
}

func (controller *VaccineDetailControllerImpl) GetDetailByFamilyId(c echo.Context) error {
	ctx := c.Request().Context()
	idParamFamilyId := c.Param("id")
	dataToInteger, err := strconv.Atoi(idParamFamilyId)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}

	dataService, err := controller.sessionDetail.GetDetailByFamilyId(ctx, dataToInteger)
	if err != nil {
		return controllers.InternalServerError(c, http.StatusInternalServerError, err)
	}

	var response []web.SessionDetailDo

	copier.Copy(&response, &dataService)
	return controllers.NewSuccessResponse(c, response)
}
