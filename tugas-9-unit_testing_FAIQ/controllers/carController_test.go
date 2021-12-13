package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
)

func TestGetCars(t *testing.T) {
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
	obj := e.POST("/login").
		WithJSON(data).
		Expect().
		Status(http.StatusOK).JSON().Object()

	token := obj.Value("token").String().Raw()

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})

	t.Run("Expected Find ALL car", func(t *testing.T) {
		auth.GET("/admin/cars").
			Expect().
			Status(http.StatusOK).JSON().Object()
	})
}

func TestCreateCar(t *testing.T) {

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
	obj := e.POST("/login").
		WithJSON(data).
		Expect().
		Status(http.StatusOK).JSON().Object()

	token := obj.Value("token").String().Raw()

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})

	t.Run("Expected Insert order", func(t *testing.T) {
		dataForInsert := map[string]interface{}{
			"item":   "Toyota Supra MK-III_test",
			"price":  19900000000,
			"amount": 2,
			"status": "paid_test",
		}

		auth.POST("/admin/cars").WithJSON(dataForInsert).
			Expect().
			Status(http.StatusOK)
	})

	t.Run("Expected Insert order, but NOT inserted that order", func(t *testing.T) {
		dataForInsert := map[string]interface{}{
			"item":   true,
			"price":  "must integer",
			"amount": "must integer too",
			"status": 9090,
		}

		auth.POST("/admin/cars").WithJSON(dataForInsert).
			Expect().
			Status(http.StatusOK)
	})

}

func TestUpdateCar(t *testing.T) {

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
	obj := e.POST("/login").
		WithJSON(data).
		Expect().
		Status(http.StatusOK).JSON().Object()

	token := obj.Value("token").String().Raw()

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})

	t.Run("Expected Update order", func(t *testing.T) {
		dataForUpdate := map[string]interface{}{
			"item":   "Honda Jazz-test-update",
			"price":  150000000,
			"amount": 2,
			"status": "purchased-test-update",
		}

		auth.PUT("/admin/cars/{id}").WithPath("id", 6).
			WithJSON(dataForUpdate).
			Expect().
			Status(http.StatusOK)
	})

	t.Run("Expected Update order, but NOT updated that car", func(t *testing.T) {
		dataForUpdate := map[string]interface{}{
			"item":   "Honda Jazz-test-update",
			"price":  150000000,
			"amount": 2,
			"status": "purchased-test-update",
		}

		auth.PUT("/admin/cars/{id}").WithPath("id", 999).
			WithJSON(dataForUpdate).
			Expect().
			Status(http.StatusOK)
	})

}

func TestDeleteCar(t *testing.T) {

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
	obj := e.POST("/login").
		WithJSON(data).
		Expect().
		Status(http.StatusOK).JSON().Object()

	token := obj.Value("token").String().Raw()

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})

	t.Run("Expected Delete order", func(t *testing.T) {
		auth.DELETE("/admin/cars/{id}").WithPath("id", 8).
			Expect().
			Status(http.StatusOK)
	})

	t.Run("Expected Delete order, but NOT deleted that car", func(t *testing.T) {
		auth.DELETE("/admin/cars/{id}").WithPath("id", 999).
			Expect().
			Status(http.StatusOK)
	})

}
