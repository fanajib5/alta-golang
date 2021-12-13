package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"tugas9unitTesting/config"
	model "tugas9unitTesting/models"
	"tugas9unitTesting/routes"

	"github.com/gavv/httpexpect"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func PrepareTestData() error {
	config.InitDB()

	config.DB.Create(model.User{
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Id:       99,
		Name:     "unit-test",
		Email:    "unit-test",
		Username: "unit-test",
		Password: "unit-test",
		Admin:    true,
	})

	return nil
}

func InitEcho() *echo.Echo {

	e := routes.New()

	return e
}

func TestIsUserLoggedin(t *testing.T) {
	PrepareTestData()

	handler := InitEcho()

	// run server using httptest
	server := httptest.NewServer(handler)

	defer server.Close()

	// create httpexpect for Mock instance
	e := httpexpect.New(t, server.URL)

	//==================GET JWT TOKEN FOR ADD IN HEADER REQUEST===================
	data := map[string]interface{}{
		"username": "unit-test",
		"password": "unit-test",
	}
	// get token
	obj := e.POST("/login").WithJSON(data).Expect().
		Status(http.StatusOK).JSON().Object()

	token := obj.Value("token").String().Raw()

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})

	t.Run("Expected Get validation the loggedin admin", func(t *testing.T) {
		auth.GET("/admin/").
			Expect().
			Status(http.StatusOK).JSON().Object()
	})
}
