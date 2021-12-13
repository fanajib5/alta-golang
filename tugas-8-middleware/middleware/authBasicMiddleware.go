package middleware

import (
	"tugas8middleware/config"
	model "tugas8middleware/models"

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
