package model

type Order struct {
	Id     int    `json:"id" form:"id"`
	Item   string `json:"item" form:"item"`
	Price  int    `json:"price" form: "price"`
	Amount int    `json:"amount" form: "amount"`
	Status string `jsonL"status" form:"status"`
}
