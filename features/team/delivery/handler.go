package delivery

import (
	"immersive-dashboard/features/team"
	"immersive-dashboard/middlewares"
	"immersive-dashboard/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TeamDelivery struct {
	teamService team.ServiceInterface
}

func New(service team.ServiceInterface, e *echo.Echo) {
	handler := &TeamDelivery{
		teamService: service,
	}

	e.GET("/teams", handler.GetAll, middlewares.JWTMiddleware())
	e.POST("/teams", handler.Create, middlewares.JWTMiddleware())
	// e.GET("/users/:id", handler.GetById)
	// e.PUT("/users/:id", handler.Update)
	// e.DELETE("/users/:id", handler.Delete)
}

func (delivery *TeamDelivery) GetAll(c echo.Context) error {
	results, err := delivery.teamService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read all users", dataResponse))
}

func (delivery *TeamDelivery) Create(c echo.Context) error {
	teamInput := TeamRequest{}
	errBind := c.Bind(&teamInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := toCore(teamInput)
	err := delivery.teamService.Create(dataCore)
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
