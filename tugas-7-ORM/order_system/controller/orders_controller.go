package controller

import (
	"net/http"
	"strconv"

	"tugasOrderSystem/config"
	"tugasOrderSystem/model"

	"github.com/labstack/echo"
)

func GetAllOrders(c echo.Context) error {
	var orders []model.Order

	err := config.DB.Find(&orders).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get orders",
		"data":    orders,
	})
}

func CreateOrder(c echo.Context) error {
	var order = model.Order{}
	// binding data
	c.Bind(&order)

	err := config.DB.Create(&order).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create order",
		"data":    order,
	})
}

func UpdateOrder(c echo.Context) error {
	var order = model.Order{}
	id, _ := strconv.Atoi(c.Param("id"))
	// r := requestUser{}

	c.Bind(&order)

	updaetDb := config.DB.Model(&order).Where("id = ?", id).UpdateColumns(map[string]interface{}{
		"item":   order.Item,
		"amount": order.Amount,
		"price":  order.Price,
		"status": order.Status,
	})

	if updaetDb.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": updaetDb.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "success update order",
		"newOrderData": order,
	})
}

func DeleteOrder(c echo.Context) error {
	var order = model.Order{}
	id, _ := strconv.Atoi(c.Param("id"))

	err := config.DB.Model(&order).Where("id = ?", id).Delete(id).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete order",
	})
}
