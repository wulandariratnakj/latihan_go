package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id      uint   `gorm:"primaryKey" json:"id"`
	Name    string `json:"name"`    // name
	Age     int    `json:"age"`     // age
	Gender  string `json:"gender"`  // age
	Address string `json:"address"` // age
}

var DB *gorm.DB

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	initDB()
	e := echo.New()
	e.POST("/users", AddUsers)
	e.GET("/users", GetUsers)
	e.GET("/users/:id", GetDetailUsers)
	e.POST("/login", Login)
	e.Logger.Fatal(e.Start(":8080"))
}

func initDB() {
	dsn := "root:123@tcp(127.0.0.1:3306)/prakerja_2?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("connection DB")
	}
	migration()
}

func migration() {
	DB.AutoMigrate(&User{})
}

func AddUsers(c echo.Context) error {
	var user User
	c.Bind(&user)

	result := DB.Create(&user)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			"failed", nil,
		})
	}

	return c.JSON(http.StatusOK, Response{
		"sucess", user,
	})
}

func Login(c echo.Context) error {
	// email := c.FormValue("email")
	// password := c.FormValue("password")
	var loginRequest LoginRequest
	c.Bind(&loginRequest)

	return c.JSON(http.StatusOK, Response{
		"sucess", loginRequest,
	})
}

func GetDetailUsers(c echo.Context) error {
	id := c.Param("id")
	// logic
	return c.JSON(http.StatusOK, Response{
		"sucess", id,
	})
}

func GetUsers(c echo.Context) error {
	var users []User

	result := DB.Find(&users)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			"failed", nil,
		})
	}

	return c.JSON(http.StatusOK, Response{
		"sucess", users,
	})
}
