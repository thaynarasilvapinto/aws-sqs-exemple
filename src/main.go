package main

import (
	"src/server"
	"src/sheduled"
)

func main() {
	go sheduled.Bath()

	server.StartApllication()
}
