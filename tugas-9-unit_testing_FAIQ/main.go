package main

import (
	"tugas9unitTesting/config"
	"tugas9unitTesting/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	// Run REST API Server
	e.Start(":8080")
}
