package main

import (
	"src/config"
	"src/server"
	"src/sheduled"
)

func main() {
	config := config.Configuration()

	go sheduled.Bath("1", config)
	go sheduled.Bath("2", config)
	go sheduled.Bath("3", config)
	go sheduled.Bath("4", config)
	go sheduled.Bath("5", config)
	go sheduled.Bath("6", config)
	go sheduled.Bath("7", config)
	go sheduled.Bath("8", config)
	go sheduled.Bath("9", config)
	go sheduled.Bath("10", config)
	server.StartApllication(config)
}
