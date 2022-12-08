package delivery

import (
	"immersive-dashboard/features/feedback"
	"immersive-dashboard/middlewares"
	"immersive-dashboard/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FeedbackDelivery struct {
	feedbackService feedback.ServiceInterface
}

func New(service feedback.ServiceInterface, e *echo.Echo) {
	handler := &FeedbackDelivery{
		feedbackService: service,
	}

	e.GET("/feedbacks", handler.GetAll, middlewares.JWTMiddleware())
	e.POST("/feedbacks", handler.Create, middlewares.JWTMiddleware())
	// e.GET("/users/:id", handler.GetById)
	e.PUT("/users/:id", handler.UpdateData)
	// e.DELETE("/users/:id", handler.Delete)
}

func (delivery *FeedbackDelivery) GetAll(c echo.Context) error {
	results, err := delivery.feedbackService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read all feedbacks", dataResponse))
}

func (delivery *FeedbackDelivery) Create(c echo.Context) error {
	feedbackInput := FeedbackRequest{}
	errBind := c.Bind(&feedbackInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := toCore(feedbackInput)
	err := delivery.feedbackService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("success create data"))
}

// func (delivery *FeedbackDelivery) GetById(c echo.Context) error {
// 	id, errBind := strconv.Atoi(c.Param("id"))
// 	if errBind != nil {
// 		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
// 	}

// 	IdFeedback, err := delivery.feedbackService.GetById(id)
// 	dataResponse := fromCore(IdFeedback)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
// 	}
// 	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get users", dataResponse))
// }

func (delivery *FeedbackDelivery) UpdateData(c echo.Context) error {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error conv data "+errConv.Error()))
	}

	feedbackInput := FeedbackRequest{}
	errBind := c.Bind(&feedbackInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := toCore(feedbackInput)
	errUpt := delivery.feedbackService.UpdateFeedback(dataCore, id)
	if errUpt != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Db update "+errUpt.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success update data"))
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
