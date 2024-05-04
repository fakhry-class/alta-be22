package main

import (
	"be22/clean-arch/app/configs"
	"be22/clean-arch/app/databases"
	"be22/clean-arch/app/routers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := configs.InitConfig()
	dbMysql := databases.InitDBMysql(cfg)
	// dbPosgres := databases.InitDBPosgres(cfg)

	// create new instance echo
	e := echo.New()

	routers.InitRouter(e, dbMysql)
	e.Use(middleware.Logger())
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, uri=${uri}, status=${status}\n",
	// }))
	// start server and port
	e.Logger.Fatal(e.Start(":8080"))
}
