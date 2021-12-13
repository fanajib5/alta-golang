package controller

import (
	"net/http"
	"strconv"

	"tugas9unitTesting/config"
	model "tugas9unitTesting/models"

	"github.com/labstack/echo"
)

func GetOrders(c echo.Context) error {
	var orders []model.Order

	a := &IsAdmin{}
	a.ValidateAdmin(c)
	if !a.Admin {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message":    "Oops! Akun anda bukan admin",
			"accessRole": "employee",
		})
	}

	errFind := config.DB.Find(&orders).Error
	if errFind != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": errFind.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    orders,
	})
}

func CreateOrder(c echo.Context) error {
	a := &IsAdmin{}
	a.ValidateAdmin(c)
	if !a.Admin {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message":    "Oops! Akun anda bukan admin",
			"accessRole": "employee",
		})
	}

	order := model.Order{}
	c.Bind(&order)

	errSave := config.DB.Save(&order).Error
	if errSave != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": errSave.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create order",
		"data":    order,
	})
}

func UpdateOrder(c echo.Context) error {
	a := &IsAdmin{}
	a.ValidateAdmin(c)
	if !a.Admin {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message":    "Oops! Akun anda bukan admin",
			"accessRole": "employee",
		})
	}

	var order = model.Order{}
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&order)

	updateColumns := config.DB.Model(&order).Where("id = ?", id).UpdateColumns(map[string]interface{}{
		"item":   order.Item,
		"amount": order.Amount,
		"price":  order.Price,
		"status": order.Status,
	})

	if updateColumns.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": updateColumns.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "success update order",
		"newOrderData": order,
	})
}

func DeleteOrder(c echo.Context) error {
	a := &IsAdmin{}
	a.ValidateAdmin(c)
	if !a.Admin {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message":    "Oops! Akun anda bukan admin",
			"accessRole": "employee",
		})
	}

	var order = model.Order{}
	id, _ := strconv.Atoi(c.Param("id"))

	errDelete := config.DB.Model(&order).Where("id = ?", id).Delete(id).Error
	if errDelete != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": errDelete.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete order",
	})
}
