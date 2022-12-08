package delivery

import (
	"immersive-dashboard/features/feedback"
	"immersive-dashboard/middlewares"
	"immersive-dashboard/utils/helper"
	"net/http"

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
	e.POST("/feedbacks", handler.Create)
	// e.GET("/users/:id", handler.GetById)
	// e.PUT("/users/:id", handler.Update)
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
