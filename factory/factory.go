package factory

import (
	authDelivery "immersive-dashboard/features/auth/delivery"
	authRepo "immersive-dashboard/features/auth/repository"
	authService "immersive-dashboard/features/auth/service"

	userDelivery "immersive-dashboard/features/user/delivery"
	userRepo "immersive-dashboard/features/user/repository"
	userService "immersive-dashboard/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {

	authRepoFactory := authRepo.New(db)
	authServiceFactory := authService.New(authRepoFactory)
	authDelivery.New(authServiceFactory, e)

	userRepoFactory := userRepo.New(db)
	userServiceFactory := userService.New(userRepoFactory)
	userDelivery.New(userServiceFactory, e)

}
