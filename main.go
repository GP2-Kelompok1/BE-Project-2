package main

import (
	"fmt"
	"immersive-dashboard/config"
	"immersive-dashboard/factory"
	"immersive-dashboard/utils/database/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
	// db := posgresql.InitDB(cfg)

	mysql.DBMigration(db)

	e := echo.New()

	factory.InitFactory(e, db)

	//middlwares
	e.Pre(middleware.RemoveTrailingSlash()) // fix garis miring yang typo2
	e.Use(middleware.CORS())                // akses agar ga di blok
	e.Use(middleware.Logger())              // log request response

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT)))
}
