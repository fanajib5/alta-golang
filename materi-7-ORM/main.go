package main

import (
	"materi7orm/config"
	"materi7orm/route"
)

func main() {
	config.InitDB()

	e := route.New()
	// Run REST API Server
	e.Start(":8000")
}
