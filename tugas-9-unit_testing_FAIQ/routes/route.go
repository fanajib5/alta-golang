package routes

import (
	"tugas9unitTesting/constants"
	controller "tugas9unitTesting/controllers"
	m "tugas9unitTesting/middleware"

	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	m.LogMiddleware(e)
	e.POST("/login", controller.LoginUser)

	eJwt := e.Group("/admin")
	eJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eJwt.GET("/", controller.IsUserLoggedin)

	eJwt.GET("/orders", controller.GetOrders)
	eJwt.POST("/orders", controller.CreateOrder)
	eJwt.PUT("/orders/:id", controller.UpdateOrder)
	eJwt.DELETE("/orders/:id", controller.DeleteOrder)

	eJwt.GET("/cars", controller.GetCars)
	eJwt.POST("/cars", controller.CreateCar)
	eJwt.PUT("/cars/:id", controller.UpdateCar)
	eJwt.DELETE("/cars/:id", controller.DeleteCar)

	return e
}
