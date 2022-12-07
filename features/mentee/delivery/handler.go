package delivery

import (
	"immersive-dashboard/features/mentee"
	"immersive-dashboard/features/user"
	"immersive-dashboard/middlewares"
	"immersive-dashboard/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MenteeDelivery struct {
	menteeService mentee.ServiceInterface
}

func New(service mentee.ServiceInterface, e *echo.Echo) {
	handler := &MenteeDelivery{
		menteeService: service,
	}

	e.GET("/mentees", handler.GetAll, middlewares.JWTMiddleware())
	e.POST("/mentees", handler.Create)

func (delivery *MenteeDelivery) GetAll(c echo.Context) error {
	results, err := delivery.menteeService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read all users", dataResponse))
}