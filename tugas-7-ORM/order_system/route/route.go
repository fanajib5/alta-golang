package route

import (
	"tugasOrderSystem/controller"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	// ROUTING
	e.GET("/orders", controller.GetAllOrders)
	e.POST("/orders", controller.CreateOrder)
	e.PUT("/orders/:id", controller.UpdateOrder)
	e.DELETE("/orders/:id", controller.DeleteOrder)

	return e
}
