package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// create new instance echo
	e := echo.New()
	// define endpoint
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "POST Hello, World!")
	})
	e.GET("/articles", GetArticlesController)
	e.PUT("/articles", GetArticlesController)
	e.DELETE("/articles", GetArticlesController)
	// start server and port
	e.Logger.Fatal(e.Start(":8080"))
}

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// rendering data
func GetArticlesController(c echo.Context) error {
	// data dummy
	var data = []article{
		{1, "lorem", "lorem"},
		{2, "ipsum", "ipsum"},
		{3, "alterra", "academy"},
		{4, "BE", "22"},
	}
	return c.JSON(http.StatusOK, data)
}
