package controller

import (
	"net/http"

	"materi8middleware/config"
	"materi8middleware/middleware"
	model "materi8middleware/models"

	"github.com/labstack/echo"
)

func GetEmployee(c echo.Context) error {
	var employees []model.Employee

	err := config.DB.Find(&employees).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    employees,
	})
}

func CreateEmployee(c echo.Context) error {
	// binding data
	employee := model.Employee{}
	c.Bind(&employee)

	err := config.DB.Save(&employee).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create employee",
		"data":    employee,
	})
}

func LoginEmployee(c echo.Context) error {
	// binding data
	employee := model.Employee{}
	c.Bind(&employee)

	err := config.DB.Where("email = ? AND password = ?", employee.Email, employee.Password).First(&employee).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":     "failed login",
			"description": err.Error(),
		})
	}

	token, err := middleware.GenerateToken(employee.Id, employee.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":     "failed login",
			"description": err.Error(),
		})
	}

	employeeResponse := model.EmployeeResponse{
		employee.Id,
		employee.Name,
		employee.Age,
		employee.Origin,
		employee.Status,
		employee.Email,
		token,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success login",
		"employee": employeeResponse,
	})
}
