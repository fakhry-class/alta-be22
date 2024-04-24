package main

import (
	"net/http"
	"strconv"

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

	// URL Param
	e.GET("/articles/:articleid", GetArticleByIdController)
	// Query Params
	e.GET("/products", GetAllProductsController)
	// form value
	e.POST("/products", AddProductController)
	// binding
	e.POST("/articles", AddArticleController)

	// start server and port
	e.Logger.Fatal(e.Start(":8080"))
}

type article struct {
	ID      int    `json:"id" form:"id"`
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// rendering data
func GetArticlesController(c echo.Context) error {
	qTitle := c.QueryParam("title")
	qPage := c.QueryParam("page")
	// data dummy
	var data = []article{
		{1, "lorem", "lorem"},
		{2, "ipsum", "ipsum"},
		{3, "alterra", "academy"},
		{4, "BE", "22"},
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"query":   qTitle,
		"page":    qPage,
		"results": data,
	})
}

// url param
func GetArticleByIdController(c echo.Context) error {
	id := c.Param("articleid")

	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success read url param",
		"id":      id,
	})
}

// query param
func GetAllProductsController(c echo.Context) error {
	qName := c.QueryParam("name")
	qPage := c.QueryParam("page")
	return c.JSON(http.StatusOK, map[string]any{
		"status":     "success",
		"message":    "success read query param",
		"name_query": qName,
		"page":       qPage,
	})
}

// request body via form value
func AddProductController(c echo.Context) error {
	name := c.FormValue("name")
	price := c.FormValue("price")
	// file, fileHeader, err := c.Request().FormFile("image")

	priceFloat, errConv := strconv.ParseFloat(price, 64)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error convert data: " + errConv.Error(),
		})
	}

	newProduct := Product{}
	newProduct.Name = name
	newProduct.Price = priceFloat

	return c.JSON(http.StatusCreated, map[string]any{
		"status":  "success",
		"message": "success add product via form value",
		"result":  newProduct,
	})
}

// request body via binding/bind
func AddArticleController(c echo.Context) error {
	newArticle := article{}
	errBind := c.Bind(&newArticle)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error bind data: " + errBind.Error(),
		})
	}
	// data newArticle simpan ke DB

	return c.JSON(http.StatusCreated, map[string]any{
		"status":  "success",
		"message": "success add article via bind json",
		"result":  newArticle,
	})
}
