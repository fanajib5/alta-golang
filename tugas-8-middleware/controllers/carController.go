package controller

import (
	"net/http"
	"strconv"

	"tugas8middleware/config"
	"tugas8middleware/middleware"
	model "tugas8middleware/models"

	"github.com/labstack/echo"
)

func GetCars(c echo.Context) error {
	var cars []model.Car

	// employee has access
	// without ADMIN VALIDATION

	errFind := config.DB.Find(&cars).Error
	if errFind != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": errFind.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all cars data",
		"data":    cars,
	})
}

func CreateCar(c echo.Context) error {
	a := &middleware.IsAdmin{}
	a.ValidateAdmin(c)
	if !a.Admin {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message":    "Oops! Akun anda bukan admin",
			"accessRole": "employee",
		})
	}

	car := model.Car{}
	c.Bind(&car)

	errSave := config.DB.Save(&car).Error
	if errSave != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": errSave.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create car",
		"data":    car,
	})
}

func UpdateCar(c echo.Context) error {
	a := &middleware.IsAdmin{}
	a.ValidateAdmin(c)
	if !a.Admin {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message":    "Oops! Akun anda bukan admin",
			"accessRole": "employee",
		})
	}

	var car = model.Car{}
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&car)

	updateDb := config.DB.Model(&car).Where("id = ?", id).UpdateColumns(map[string]interface{}{
		"type":   car.Type,
		"amount": car.Amount,
		"price":  car.Price,
		"status": car.Status,
	})

	if updateDb.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": updateDb.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "success update car",
		"newOrderData": car,
	})
}

func DeleteCar(c echo.Context) error {
	a := &middleware.IsAdmin{}
	a.ValidateAdmin(c)
	if !a.Admin {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message":    "Oops! Akun anda bukan admin",
			"accessRole": "employee",
		})
	}

	var car = model.Car{}
	id, _ := strconv.Atoi(c.Param("id"))

	errDelete := config.DB.Model(&car).Where("id = ?", id).Delete(id).Error
	if errDelete != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": errDelete.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete car",
	})
}
