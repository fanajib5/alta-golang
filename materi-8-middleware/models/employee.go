package model

type Employee struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Age      int    `json:"age" form: "age"`
	Origin   string `json:"origin" form: "origin"`
	Status   string `json:"status" form:"status"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type EmployeeResponse struct {
	Id     int    `json:"id" form:"id"`
	Name   string `json:"name" form:"name"`
	Age    int    `json:"age" form: "age"`
	Origin string `json:"origin" form: "origin"`
	Status string `json:"status" form:"status"`
	Email  string `json:"email" form:"email"`
	Token  string `json:"token" form:"token"`
}
