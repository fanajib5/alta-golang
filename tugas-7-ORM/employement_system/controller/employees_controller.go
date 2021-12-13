package controller

import (
	"net/http"
	"strconv"

	"employementSystem/config"
	"employementSystem/model"

	"github.com/labstack/echo"
)

func GetAllEmployees(c echo.Context) error {
	var employees []model.Employee

	err := config.DB.Find(&employees).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get employees data",
		"data":    employees,
	})
}

func CreateEmployeeData(c echo.Context) error {
	var employee = model.Employee{}
	// binding data
	c.Bind(&employee)

	err := config.DB.Create(&employee).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create employee data",
		"data":    employee,
	})
}

func UpdateEmployeeData(c echo.Context) error {
	var employee = model.Employee{}
	id, _ := strconv.Atoi(c.Param("id"))
	// r := requestUser{}

	c.Bind(&employee)

	updaetDb := config.DB.Model(&employee).Where("id = ?", id).UpdateColumns(map[string]interface{}{
		"name":   employee.Name,
		"age":    employee.Age,
		"origin": employee.Origin,
		"status": employee.Status,
	})

	if updaetDb.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": updaetDb.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":         "success update employee data",
		"newEmployeeData": employee,
	})
}

func DeleteEmployeeData(c echo.Context) error {
	var employee = model.Employee{}
	id, _ := strconv.Atoi(c.Param("id"))

	err := config.DB.Model(&employee).Where("id = ?", id).Delete(id).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete employee",
	})
}
