package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	// ID        string `gorm:"primarykey"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string    `json:"name" form:"name"`
	Email     string    `gorm:"unique" json:"email" form:"email"`
	Password  string    `json:"password" form:"password"`
	Phone     string    `json:"phone" form:"phone"`
	Address   string    `json:"address" form:"address"`
	StoreName string    `json:"store_name" form:"store_name"`
	Products  []Product `gorm:"foreignKey:UserID;references:ID"`
}
type Product struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	UserID      uint
	ProductName string
	Description string
	Price       float64
	Stock       int
	Type        string
}

/*
TODO: tambahkan table favorite
*/

func InitDB() {
	var err error
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// dsn := "root:qwerty123@tcp(127.0.0.1:3306)/db_be22_unit2?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBNAME"))
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Product{})
}

func main() {
	InitDB()
	fmt.Println("running")
	InitialMigration()

	// create new instance echo
	e := echo.New()
	e.GET("/users", GetAllUsersController)
	e.POST("/users", AddUserController)
	e.DELETE("/users/:id", DeleteUserByIdController)

	/*
		TODO:
		Update user by id
		CRUD product
		CRUD favorite
	*/

	// start server and port
	e.Logger.Fatal(e.Start(":8080"))
}

func GetAllUsersController(c echo.Context) error {
	var allUsers []User                          // var penampung data yg dibaca dari db
	tx := DB.Preload("Products").Find(&allUsers) // select * from users
	if tx.Error != nil {                         //check apakah terjadi error saat menjalankan query
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error read data",
		})
	}
	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success read data",
		"results": allUsers,
	})
}

func AddUserController(c echo.Context) error {
	newUser := User{}
	errBind := c.Bind(&newUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error bind data: " + errBind.Error(),
		})
	}
	// data newArticle simpan ke DB
	tx := DB.Create(&newUser) //menjalankan query insert into users(...) values(....)
	if tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error insert data " + tx.Error.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"status":  "success",
		"message": "success add user",
	})
}

func DeleteUserByIdController(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error convert id: " + errConv.Error(),
		})
	}

	tx := DB.Delete(&User{}, idConv)
	if tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error delete data " + tx.Error.Error(),
		})
	}
	if tx.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  "failed",
			"message": "data not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success delete user",
	})

}
