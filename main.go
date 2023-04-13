package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var data []User = []User{
	{"alta", 10},
	{"academy", 11},
}

func main() {
	e := echo.New()
	e.GET("/users", GetUsers)
	e.GET("/users/:id", GetDetailUsers)
	e.POST("/login", Login)
	e.Logger.Fatal(e.Start(":8080"))
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
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")
	// logic
	return c.JSON(http.StatusOK, Response{
		"sucess", page + " & " + limit,
	})
}
