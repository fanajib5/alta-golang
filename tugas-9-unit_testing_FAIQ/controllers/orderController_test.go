package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
)

func TestGetOrders(t *testing.T) {
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

	// fmt.Println("obj:", obj)
	// fmt.Println("obj value:", obj.Value("token"))

	token := obj.Value("token").String().Raw()

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})

	t.Run("Expected Find ALL get order", func(t *testing.T) {
		auth.GET("/admin/orders").
			Expect().
			Status(http.StatusOK).JSON().Object()
	})
}

func TestCreateOrder(t *testing.T) {

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

	t.Run("Expected Insert order, then call get one By ID and found that order", func(t *testing.T) {
		dataForInsert := map[string]interface{}{
			"item":   "Ransel-test",
			"price":  999000,
			"amount": 1,
			"status": "on process-test",
		}

		auth.POST("/admin/orders").WithJSON(dataForInsert).
			Expect().
			Status(http.StatusOK)
	})

	t.Run("Expected Insert order, then call get one By ID But NOT found that order", func(t *testing.T) {
		dataForInsert := map[string]interface{}{
			"item":   true,
			"price":  "must integer",
			"amount": "must integer too",
			"status": 9090,
		}

		auth.POST("/admin/orders").WithJSON(dataForInsert).
			Expect().
			Status(http.StatusOK)
	})

}

func TestUpdateOrder(t *testing.T) {

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
			"item":   "Ransel-test-update",
			"price":  900990,
			"amount": 9,
			"status": "on process-test-update",
		}

		auth.PUT("/admin/orders/{id}").WithPath("id", 8).
			WithJSON(dataForUpdate).
			Expect().
			Status(http.StatusOK)
	})

	t.Run("Expected Update order", func(t *testing.T) {
		dataForUpdate := map[string]interface{}{
			"item":   "Ransel-test-update",
			"price":  900990,
			"amount": 9,
			"status": "on process-test-update",
		}

		auth.PUT("/admin/orders/{id}").WithPath("id", 7).
			WithJSON(dataForUpdate).
			Expect().
			Status(http.StatusOK)
	})

}

func TestDeleteOrder(t *testing.T) {

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
		auth.DELETE("/admin/orders/{id}").WithPath("id", 8).
			Expect().
			Status(http.StatusOK)
	})

	t.Run("Expected Delete order", func(t *testing.T) {
		auth.DELETE("/admin/orders/{id}").WithPath("id", 7).
			Expect().
			Status(http.StatusOK)
	})

}
