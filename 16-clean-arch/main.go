package main

import (
	"be22/clean-arch/app/configs"
	"be22/clean-arch/app/databases"
	"be22/clean-arch/app/routers"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := configs.InitConfig()
	// dbMysql := databases.InitDBMysql(cfg)
	dbMysql := databases.InitDBPosgres(cfg)

	// create new instance echo
	e := echo.New()

	routers.InitRouter(e, dbMysql)
	// start server and port
	e.Logger.Fatal(e.Start(":8080"))
}
