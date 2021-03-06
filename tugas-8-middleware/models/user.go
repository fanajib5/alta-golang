package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       int    `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Name     string `json:"name" form: "name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Admin    bool   `json:"admin" form:"admin"`
}
