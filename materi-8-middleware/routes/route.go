package routes

import (
	"materi8middleware/constants"
	controller "materi8middleware/controllers"
	m "materi8middleware/middleware"

	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/employees", controller.GetEmployee)
	m.LogMiddleware(e)
	e.POST("/employees", controller.CreateEmployee)
	e.POST("/login", controller.LoginEmployee)

	// basic auth
	eAuthBasic := e.Group("/auth")
	eAuthBasic.Use(mid.BasicAuth(m.BasicAuthDB))
	eAuthBasic.GET("/employees", controller.GetEmployee)

	//jwt auth
	eJwt := e.Group("/jwt")
	eJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eJwt.GET("/employees", controller.GetEmployee)

	return e
}
