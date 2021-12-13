package main

import (
	"employementSystem/config"
	"employementSystem/route"
)

func main() {
	config.InitDB()

	e := route.New()

	e.Start(":8080")
}
