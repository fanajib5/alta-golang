package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	Id    int    `json:"id" form:"id"`
	Age   int    `json:"age" form:"age"`
	Email string `json:"email" form:"email"`
	Name  string `json:"name" form: "name"`
}

func main() {
	e := echo.New()

	// ROUTING
	e.GET("/user/:id/:age", UserController)
	e.POST("/user", RegisterController)

	// Run REST API Server
	e.Start(":8000")
}

// response "Hello User!"
func UserController(e echo.Context) error {
	// mengambil URL params
	id, _ := strconv.Atoi(e.Param("id"))
	age, _ := strconv.Atoi(e.Param("age"))

	// mengambil query params
	search := e.QueryParam("search")
	sort := e.QueryParam("sort")

	user := User{Id: id, Age: age, Email: "faiq.najib@gmail.com", Name: "Faiq Najib"}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"user":   user,
		"search": search,
		"sort":   sort,
	})
}

//register method post
func RegisterController(e echo.Context) error {
	// email := e.FormValue("email")
	// name := e.FormValue("name")

	user := User{}

	e.Bind(&user)

	return e.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}
