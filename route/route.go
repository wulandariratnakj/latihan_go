package route

import (
	"prakerja2/controller"

	"github.com/labstack/echo/v4"
)

func InitRoute() *echo.Echo {
	e := echo.New()
	e.POST("/users", controller.AddUsers)
	e.GET("/users", controller.GetUsers)
	e.GET("/users/:id", controller.GetDetailUsers)
	e.POST("/login", controller.Login)
	return e
}
