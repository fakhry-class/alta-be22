package handler

import (
	"be22/clean-arch/app/middlewares"
	"be22/clean-arch/features/user"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.ServiceInterface
}

func New(us user.ServiceInterface) *UserHandler {
	return &UserHandler{
		userService: us,
	}
}

func (uh *UserHandler) Register(c echo.Context) error {
	// membaca data dari request body
	newUser := UserRequest{}
	errBind := c.Bind(&newUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error bind data: " + errBind.Error(),
		})
	}

	// mapping  dari request ke core
	inputCore := user.Core{
		Name:      newUser.Name,
		Email:     newUser.Email,
		Password:  newUser.Password,
		Phone:     newUser.Phone,
		Address:   newUser.Address,
		StoreName: newUser.StoreName,
	}
	// memanggil/mengirimkan data ke method service layer
	errInsert := uh.userService.Create(inputCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error insert data " + errInsert.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]any{
		"status":  "success",
		"message": "success add user",
	})
}

func (uh *UserHandler) GetAll(c echo.Context) error {
	result, err := uh.userService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error read data",
		})
	}
	var allUsersResponse []UserResponse
	for _, value := range result {
		allUsersResponse = append(allUsersResponse, UserResponse{
			ID:    value.ID,
			Name:  value.Name,
			Email: value.Email,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success read data",
		"results": allUsersResponse,
	})
}

func (uh *UserHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error convert id: " + errConv.Error(),
		})
	}
	err := uh.userService.Delete(uint(idConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error delete data " + err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success delete user",
	})
}

func (uh *UserHandler) Login(c echo.Context) error {
	var reqLoginData = LoginRequest{}
	errBind := c.Bind(&reqLoginData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error bind data: " + errBind.Error(),
		})
	}
	result, token, err := uh.userService.Login(reqLoginData.Email, reqLoginData.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error login " + err.Error(),
		})
	}
	//mapping
	var resultResponse = map[string]any{
		"id":    result.ID,
		"name":  result.Name,
		"token": token,
	}
	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success login",
		"data":    resultResponse,
	})
}

func (uh *UserHandler) Profile(c echo.Context) error {
	// extract id user from jwt token
	idToken := middlewares.ExtractTokenUserId(c)
	log.Println("idtoken:", idToken)
	result, err := uh.userService.GetById(uint(idToken))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error login " + err.Error(),
		})
	}
	resultResponse := UserResponse{
		ID:    result.ID,
		Name:  result.Name,
		Email: result.Email,
	}
	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success login",
		"data":    resultResponse,
	})
}
