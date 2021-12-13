package main

import (
	"tugas8middleware/config"
	"tugas8middleware/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	// Run REST API Server
	e.Start(":8080")
}
