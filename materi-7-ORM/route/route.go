package route

import (
	"materi7orm/controller"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	// ROUTING
	e.GET("/users", controller.GetUsersController)
	e.POST("/users", controller.CreateUserController)

	return e
}
