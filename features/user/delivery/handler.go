package delivery

import (
	"immersive-dashboard/features/user"
	"immersive-dashboard/middlewares"
	"immersive-dashboard/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	userService user.ServiceInterface
}

func New(service user.ServiceInterface, e *echo.Echo) {
	handler := &UserDelivery{
		userService: service,
	}

	e.GET("/users", handler.GetAll, middlewares.JWTMiddleware())
	e.POST("/users", handler.Create, middlewares.JWTMiddleware())
	e.GET("/users/:id", handler.GetById, middlewares.JWTMiddleware())
	e.PUT("/users/:id", handler.UpdateData, middlewares.JWTMiddleware())
	// e.DELETE("/users/:id", handler.Delete)
}

func (delivery *UserDelivery) GetAll(c echo.Context) error {
	results, err := delivery.userService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read all users", dataResponse))
}

func (delivery *UserDelivery) Create(c echo.Context) error {
	userInput := UserRequest{}
	errBind := c.Bind(&userInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := toCore(userInput)
	err := delivery.userService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("success create data"))
}

// func (delivery *UserDelivery) GetById(c echo.Context) error {
// 	id, errconv := strconv.Atoi(c.Param("id"))
// 	if errconv != nil {
// 		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error read data "))
// 	}
// 	dataCore := toCore(id)
// }

// func (delivery *UserDelivery) Update(c echo.Context) error {

// }

// func (delivery *UserDelivery) Delete(c echo.Context) error {

// }
func (delivery *UserDelivery) GetById(c echo.Context) error {
	id, errBind := strconv.Atoi(c.Param("id"))
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	IdUser, err := delivery.userService.GetById(id)
	dataResponse := fromCore(IdUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get users", dataResponse))
}

func (delivery *UserDelivery) UpdateData(c echo.Context) error {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error conv data "+errConv.Error()))
	}

	userInput := UserRequest{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := toCore(userInput)
	errUpt := delivery.userService.UpdateUser(dataCore, id)
	if errUpt != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Db update "+errUpt.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success update data"))
}
