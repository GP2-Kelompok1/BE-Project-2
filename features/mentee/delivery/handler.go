package delivery

import (
	"immersive-dashboard/features/mentee"
	"strconv"

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

	//delivery > service > repo > service > delivery
	e.GET("/mentees", handler.GetAll, middlewares.JWTMiddleware())
	e.POST("/mentees", handler.Create)
	e.GET("/mentees/:id", handler.GetById)
	e.PUT("/mentees/:id", handler.UpdateData)
	e.DELETE("/mentees/:id", handler.DeleteMentee)
}

// fungsi menampilkan semua user
func (delivery *MenteeDelivery) GetAll(c echo.Context) error {
	results, err := delivery.menteeService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read all mentees", dataResponse))
}

// fungsi pembuatan mentee
func (delivery *MenteeDelivery) Create(c echo.Context) error {
	menteeInput := MenteeRequest{}
	errBind := c.Bind(&menteeInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := toCore(menteeInput)
	err := delivery.menteeService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("success create data"))
}

func (delivery *MenteeDelivery) GetById(c echo.Context) error {
	id, errBind := strconv.Atoi(c.Param("id"))
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	IdMentee, err := delivery.menteeService.GetById(id)
	dataResponse := fromCore(IdMentee)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get users", dataResponse))
}

func (delivery *MenteeDelivery) UpdateData(c echo.Context) error {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error conv data "+errConv.Error()))
	}

	menteeInput := MenteeRequest{}
	errBind := c.Bind(&menteeInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := toCore(menteeInput)
	errUpt := delivery.menteeService.UpdateMentee(dataCore, id)
	if errUpt != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Db update "+errUpt.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success update data"))
}

func (delivery *MenteeDelivery) DeleteMentee(c echo.Context) error {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error conv data "+errConv.Error()))
	}

	errDel := delivery.menteeService.DeleteMentee(id)
	if errDel != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error delete mentee"+errDel.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete data"))

}
