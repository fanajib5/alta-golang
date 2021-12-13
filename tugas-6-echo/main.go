package main

import (
	"math"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type BmiResponse struct {
	Request  Request  `json:"Request"`
	Response Response `json:"Response"`
}

type Request struct {
	FullName string `json:"fullName" form:"fullName"`
	DoB      string `json:"DOB" form:"DOB"`
	Height   int    `json:"height" form:"height"`
	Weight   int    `json:"weight" form:"weight"`
}

type Response struct {
	FullName string  `json:"fullName"`
	Age      int     `json:"age"`
	BMI      float64 `json:"BMI"`
	Status   string  `json:"status"`
}

func YourBmiController(e echo.Context) error {
	userRequest := Request{}
	e.Bind(&userRequest)

	userDataWeight := userRequest.Weight
	weightFloat := float64(userDataWeight)
	userDataHeight := userRequest.Height
	heightFloat := float64(userDataHeight) / 100

	age := CalculateAge(userRequest.DoB)

	bmiResult := weightFloat / (heightFloat * heightFloat)
	userBmiStatus := DefiningBmi(bmiResult)

	apiResponse := Response{
		FullName: userRequest.FullName,
		Age:      age,
		BMI:      bmiResult,
		Status:   userBmiStatus,
	}

	jsonResponse := BmiResponse{
		Request:  userRequest,
		Response: apiResponse,
	}

	response := e.JSON(http.StatusOK, jsonResponse)

	return response
}

func CalculateAge(dob string) int {
	userDoB, _ := time.Parse("02/01/2006", dob)
	todayDate := time.Now()

	ageFloat := math.Floor(todayDate.Sub(userDoB).Hours() / 24 / 365)
	age := int(ageFloat)

	return age
}

func DefiningBmi(bmiValue float64) string {
	var bmiStatus string

	if bmiValue < 18.5 {
		bmiStatus = "Underweight"
	} else if 18.5 < bmiValue && bmiValue < 24.9 {
		bmiStatus = "Normal"
	} else if 24.9 < bmiValue && bmiValue > 29.9 {
		bmiStatus = "Overweight"
	} else {
		bmiStatus = "Obesity"
	}

	return bmiStatus
}

func main() {
	e := echo.New()

	e.GET("/your-bmi", YourBmiController)

	e.Start(":8000")
}
