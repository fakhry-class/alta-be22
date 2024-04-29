package routers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {
	e.GET("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"message": "hello world",
		})
	})
}
