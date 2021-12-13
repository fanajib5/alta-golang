package middleware

import (
	"tugas9unitTesting/config"
	model "tugas9unitTesting/models"

	"github.com/labstack/echo"
)

func BasicAuthDB(username, password string, c echo.Context) (bool, error) {
	var user model.User
	err := config.DB.Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
