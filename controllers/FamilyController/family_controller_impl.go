package FamilyController

import (
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"vaccine-app-be/app/middleware"
	"vaccine-app-be/controllers"
	"vaccine-app-be/controllers/FamilyController/web"
	"vaccine-app-be/services/FamilyService"
)

type FamilyControllerImpl struct {
	familyService FamilyService.FamilyService
}

func NewFamilyControllerImpl(familyService FamilyService.FamilyService) FamilyController {
	return &FamilyControllerImpl{familyService: familyService}
}

func (controller *FamilyControllerImpl) Create(c echo.Context) error {
	ctx := c.Request().Context()
	//get citizen ID
	ctxCitizenId := middleware.GetUserId(c)
	req := web.FamilyMemberRequest{}
	err := c.Bind(&req)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}

	entityFamily := FamilyService.FamilyMember{}
	copier.Copy(&entityFamily, &req)
	create, err := controller.familyService.Create(ctx, ctxCitizenId, entityFamily)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}

	response := web.FamilyMemberResponse{}
	copier.Copy(&response, &create)

	return controllers.NewSuccessResponse(c, response)
}

func (controller *FamilyControllerImpl) Update(c echo.Context) error {
	ctx := c.Request().Context()
	//get citizen ID
	id := c.Param("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}
	req := web.FamilyMemberRequest{}
	err = c.Bind(&req)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}
	entityFamily := FamilyService.FamilyMember{}
	copier.Copy(&entityFamily, &req)
	update, err := controller.familyService.Update(ctx, atoi, entityFamily)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}

	response := web.FamilyMemberResponse{}
	copier.Copy(&response, &update)

	return controllers.NewSuccessResponse(c, response)
}

func (controller *FamilyControllerImpl) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	//get citizen ID
	ctxCitizenId := middleware.GetUserId(c)
	atoi, err := strconv.Atoi(id)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}

	_, err = controller.familyService.Delete(ctx, atoi, ctxCitizenId)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, echo.Map{
		"message": "success delete",
	})
}

func (controller *FamilyControllerImpl) GetFamilyById(c echo.Context) error {
	ctx := c.Request().Context()
	//get citizen ID
	id := c.Param("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}
	byId, err := controller.familyService.GetFamilyById(ctx, atoi)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}
	response := web.FamilyMemberResponse{}
	copier.Copy(&response, &byId)

	return controllers.NewSuccessResponse(c, response)
}

func (controller *FamilyControllerImpl) GetCitizenOwnFamily(c echo.Context) error {
	ctx := c.Request().Context()
	//get citizen ID
	ctxCitizenId := middleware.GetUserId(c)
	family, err := controller.familyService.GetCitizenOwnFamily(ctx, ctxCitizenId)
	if err != nil {
		return controllers.BadRequestResponse(c, http.StatusBadRequest, err)
	}

	var response []web.FamilyMemberResponse
	copier.Copy(&response, &family)

	return controllers.NewSuccessResponse(c, response)
}
