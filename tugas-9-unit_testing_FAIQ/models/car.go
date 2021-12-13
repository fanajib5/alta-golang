package model

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	Id     int    `json:"id" form:"id"`
	Type   string `json:"item" form:"item"`
	Price  int    `json:"price" form: "price"`
	Amount int    `json:"amount" form: "amount"`
	Status string `jsonL"status" form:"status"`
}
