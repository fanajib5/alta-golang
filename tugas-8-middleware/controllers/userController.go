package controller

import (
	"net/http"
	"tugas8middleware/config"
	"tugas8middleware/middleware"
	model "tugas8middleware/models"

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

	userResponse := middleware.UserResponse{
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

// func CreateUser(c echo.Context) error {
// 	// validate = validator.New()
// 	// binding data
// 	user := model.User{}
// 	c.Bind(&user)

// 	err := config.DB.Save(&user).Error
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"message": err.Error(),
// 		})
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success create user",
// 		"data":    user,
// 	})
// }

// // func readAllCookies(c echo.Context, token) error {
// // 	for _, cookie := range c.Cookies() {

// // 		if cookie.Value == token {
// // 			return c.JSON(http.StatusOK, map[string]interface{}{
// // 				"message":     "user is already logged in",
// // 				"token": cookie.Value,
// // 			})
// // 		}
// // 	}
// // 	return c.String(http.StatusOK, "read all the cookies")
// // }

// // func writeCookie(c echo.Context, token string) error {
// // 	cookie := new(http.Cookie)
// // 	cookie.Name = "Authorization"
// // 	cookie.Value = token
// // 	cookie.Expires = time.Now().Add(24 * time.Hour)
// // 	c.SetCookie(cookie)
// // 	return c.String(http.StatusOK, "write a cookie")
// // }
