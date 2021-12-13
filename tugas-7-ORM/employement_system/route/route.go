package route

import (
	"employementSystem/controller"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	// ROUTING
	e.GET("/employees", controller.GetAllEmployees)
	e.POST("/employees", controller.CreateEmployeeData)
	e.PUT("/employees/:id", controller.UpdateEmployeeData)
	e.DELETE("/employees/:id", controller.DeleteEmployeeData)

	return e
}
