package delivery

import (
	"immersive-dashboard/features/class"
	"immersive-dashboard/middlewares"
	"immersive-dashboard/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ClassDelivery struct {
	classService class.ServiceInterface
}

func New(service class.ServiceInterface, e *echo.Echo) {
	handler := &ClassDelivery{
		classService: service,
	}

	e.GET("/classes", handler.GetAll, middlewares.JWTMiddleware())
	e.POST("/classes", handler.Create, middlewares.JWTMiddleware())
	e.GET("/users/:id", handler.GetById)
	e.PUT("/users/:id", handler.UpdateData)
	// e.DELETE("/users/:id", handler.Delete)
}

func (delivery *ClassDelivery) GetAll(c echo.Context) error {
	results, err := delivery.classService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read all users", dataResponse))
}

func (delivery *ClassDelivery) Create(c echo.Context) error {
	classInput := ClassRequest{}
	errBind := c.Bind(&classInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := toCore(classInput)
	err := delivery.classService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("success create data"))
}

func (delivery *ClassDelivery) GetById(c echo.Context) error {
	id, errBind := strconv.Atoi(c.Param("id"))
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	IdClass, err := delivery.classService.GetById(id)
	dataResponse := fromCore(IdClass)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get users", dataResponse))
}

func (delivery *ClassDelivery) UpdateData(c echo.Context) error {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error conv data "+errConv.Error()))
	}

	classInput := ClassRequest{}
	errBind := c.Bind(&classInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := toCore(classInput)
	errUpt := delivery.classService.UpdateClass(dataCore, id)
	if errUpt != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Db update "+errUpt.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success update data"))
}
