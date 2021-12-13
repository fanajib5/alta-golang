package main

import (
	"tugasOrderSystem/config"
	"tugasOrderSystem/route"
)

func main() {
	config.InitDB()

	e := route.New()

	e.Start(":8080")
}
