package controller

import (
	"net/http"
	"tugas9unitTesting/config"
	"tugas9unitTesting/middleware"
	model "tugas9unitTesting/models"

	"github.com/labstack/echo"
)

func LoginUser(c echo.Context) error {
	// binding data
	user := model.User{}
	c.Bind(&user)

	errLogin := config.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&user).Error
	if errLogin != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"message":     "failed login",
			"description": errLogin.Error(),
		})
	}

	token, errGen := middleware.GenerateToken(user.Id, user.Name, user.Admin)
	if errGen != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":     "failed login",
			"description": errGen.Error(),
		})
	}

	userResponse := UserResponse{
		Username: user.Username,
		Name:     user.Name,
		Admin:    user.Admin,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success login",
		"employee": userResponse,
		"token":    token,
	})
}
