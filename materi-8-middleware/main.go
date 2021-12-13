package main

import (
	"materi8middleware/config"
	"materi8middleware/routes"
)

func main() {
	config.InitDB()

	e := routes.New()
	// Run REST API Server
	e.Start(":8080")
}
