package routers

import (
	// "be22/clean-arch/features/product/data"
	// "be22/clean-arch/features/product/service"
	// "be22/clean-arch/features/product/handler"
	_userData "be22/clean-arch/features/user/data"
	_userHandler "be22/clean-arch/features/user/handler"
	_userService "be22/clean-arch/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {

	// factory
	userData := _userData.New(db)
	userService := _userService.New(userData)
	userHandlerAPI := _userHandler.New(userService)

	// productData := data.New(db)
	// productService := service.New(productData)
	// productHandlerAPI := handler.New(productService)

	// e.GET("/users", func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, map[string]any{
	// 		"message": "hello world",
	// 	})
	// })

	e.GET("/users", userHandlerAPI.GetAll)
	e.POST("/users", userHandlerAPI.Register)
	e.DELETE("/users/:id", userHandlerAPI.Delete)
	// e.GET("/product", productHandlerAPI.GetAll)
}
