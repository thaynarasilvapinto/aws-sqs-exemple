package main

import (
	"src/config"
	"src/server"
	"src/sheduled"
)

func main() {
	config := config.Configuration()

	go sheduled.Bath("1", config)
	server.StartApllication(config)
}
