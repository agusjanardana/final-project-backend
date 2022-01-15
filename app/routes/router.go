package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"vaccine-app-be/controllers/CitizenController"
	"vaccine-app-be/controllers/FamilyController"
	"vaccine-app-be/controllers/HealthController"
	"vaccine-app-be/controllers/VaccineController"
	"vaccine-app-be/controllers/VaccineDetailController"
	"vaccine-app-be/controllers/VaccineSessionController"
)

type ControllerList struct {
	JWTMiddleware            middleware.JWTConfig
	CitizenController        CitizenController.CitizenController
	HealthController         HealthController.HealthFacilitatorController
	FamilyController         FamilyController.FamilyController
	VaccineController        VaccineController.VaccineController
	VaccineSessionController VaccineSessionController.VaccineSessionController
	VaccineDetailController  VaccineDetailController.VaccineDetailController
}

func (c1 *ControllerList) Registration(e *echo.Echo) {
	apiV1 := e.Group("/api/v1")

	//	CITIZEN THINGS
	apiV1.POST("/citizen/registers", c1.CitizenController.Register)
	apiV1.POST("/citizen/logins", c1.CitizenController.Login)
	apiV1.PUT("/citizens", c1.CitizenController.Update, middleware.JWTWithConfig(c1.JWTMiddleware))
	apiV1.GET("/citizen/profiles", c1.CitizenController.FindCitizenById, middleware.JWTWithConfig(c1.JWTMiddleware))
	apiV1.GET("/citizen/sessions", c1.VaccineSessionController.GetCitizenAndFamilySelectedSession, middleware.JWTWithConfig(c1.JWTMiddleware))

	//	HEALTH FA THINGS
	apiV1.POST("/admin/registers", c1.HealthController.Register)
	apiV1.POST("/admin/logins", c1.HealthController.Login)
	apiV1.GET("/admins", c1.HealthController.GetAllHealthFacilitator)
	apiV1.GET("/admin/:id", c1.HealthController.FindById)
	apiV1.PUT("/admin/:id", c1.HealthController.Update, middleware.JWTWithConfig(c1.JWTMiddleware))
	apiV1.PUT("/admin/family/:id", c1.HealthController.HealthUpdateFamilyMemberStatus, middleware.JWTWithConfig(c1.JWTMiddleware))

	//  FAMILY THINGS
	apiV1.GET("/families/:id", c1.FamilyController.GetFamilyById, middleware.JWTWithConfig(c1.JWTMiddleware))
	apiV1.POST("/families", c1.FamilyController.Create, middleware.JWTWithConfig(c1.JWTMiddleware))
	apiV1.PUT("/families/:id", c1.FamilyController.Update, middleware.JWTWithConfig(c1.JWTMiddleware))
	apiV1.DELETE("/families/:id", c1.FamilyController.Delete, middleware.JWTWithConfig(c1.JWTMiddleware))
	apiV1.GET("/family/citizens", c1.FamilyController.GetCitizenOwnFamily, middleware.JWTWithConfig(c1.JWTMiddleware))

	//	VACCINE THINGS
	apiV1.GET("/vaccine/:id", c1.VaccineController.FindVaccineById, middleware.JWTWithConfig(c1.JWTMiddleware))
	apiV1.GET("/vaccines/:id", c1.VaccineController.FindVaccineOwnedByHF, middleware.JWTWithConfig(c1.JWTMiddleware))
	apiV1.POST("/vaccines", c1.VaccineController.Create, middleware.JWTWithConfig(c1.JWTMiddleware))
	apiV1.PUT("/vaccine/:id", c1.VaccineController.Update, middleware.JWTWithConfig(c1.JWTMiddleware))
	apiV1.DELETE("/vaccine/:id", c1.VaccineController.Delete, middleware.JWTWithConfig(c1.JWTMiddleware))

	// Vaccine Session
	apiV1.POST("/vaccine/sessions", c1.VaccineSessionController.CreateSession, middleware.JWTWithConfig(c1.JWTMiddleware))
	apiV1.GET("/vaccine/session/:id", c1.VaccineSessionController.GetSessionById, middleware.JWTWithConfig(c1.JWTMiddleware))
	apiV1.GET("/vaccine/session/owned/:id", c1.VaccineSessionController.GetSessionOwnedByHf, middleware.JWTWithConfig(c1.JWTMiddleware))
	apiV1.DELETE("/vaccine/session/:id", c1.VaccineSessionController.DeleteSession, middleware.JWTWithConfig(c1.JWTMiddleware))
	apiV1.PUT("/vaccine/session/:id", c1.VaccineSessionController.UpdateSession, middleware.JWTWithConfig(c1.JWTMiddleware))
	apiV1.GET("/vaccine/sessions", c1.VaccineSessionController.GetAllVaccineSession, middleware.JWTWithConfig(c1.JWTMiddleware))

	// Booking Session Detail
	apiV1.POST("/session/bookings/:id", c1.VaccineDetailController.CitizenChooseSession, middleware.JWTWithConfig(c1.JWTMiddleware))
	apiV1.GET("/session/family/:id", c1.VaccineDetailController.GetDetailByFamilyId, middleware.JWTWithConfig(c1.JWTMiddleware))
	apiV1.GET("/session/detail/:id", c1.VaccineDetailController.GetDetailById, middleware.JWTWithConfig(c1.JWTMiddleware))
	apiV1.GET("/session/:id", c1.VaccineDetailController.GetDetailBySessionId, middleware.JWTWithConfig(c1.JWTMiddleware))
}
