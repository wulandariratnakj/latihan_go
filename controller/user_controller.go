package controller

import (
	"net/http"
	"prakerja2/config"
	"prakerja2/model"

	"github.com/labstack/echo/v4"
)

func AddUsers(c echo.Context) error {
	var user model.User
	c.Bind(&user)

	result := config.DB.Create(&user)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			"failed", nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		"sucess", user,
	})
}

func GetDetailUsers(c echo.Context) error {
	id := c.Param("id")
	// logic
	return c.JSON(http.StatusOK, model.Response{
		"sucess", id,
	})
}

func GetUsers(c echo.Context) error {
	var users []model.User

	result := config.DB.Find(&users)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			"failed", nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		"sucess", users,
	})
}
