package controller

import (
	"net/http"

	jwtV3 "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type UserResponse struct {
	Username string `json:"username" form:"username"`
	Name     string `json:"name" form: "name"`
	Admin    bool   `json:"admin" form:"admin"`
}

// Check the logged in user whether he is the Super Admin or the Employee
func IsUserLoggedin(c echo.Context) error {
	user := c.Get("user").(*jwtV3.Token)
	claims := user.Claims.(jwtV3.MapClaims)
	name := claims["name"]

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success logged in",
		"data":    name,
	})
}

// Check if admin is logged in
type IsAdmin struct {
	Admin bool `json:"admin"`
}

func (a *IsAdmin) ValidateAdmin(c echo.Context) {
	user := c.Get("user").(*jwtV3.Token)
	claims := user.Claims.(jwtV3.MapClaims)
	isAdmin := claims["admin"]

	if isAdmin == true {
		a.Admin = true
	}
}
