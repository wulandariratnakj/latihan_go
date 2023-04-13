package controller

import (
	"net/http"
	"prakerja2/model"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	// email := c.FormValue("email")
	// password := c.FormValue("password")
	var loginRequest model.LoginRequest
	c.Bind(&loginRequest)

	return c.JSON(http.StatusOK, model.Response{
		"sucess", loginRequest,
	})
}
