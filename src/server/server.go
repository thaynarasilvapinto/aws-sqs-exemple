package server

import (
	"fmt"
	"log"
	"net/http"
	"src/config"
)

func StartApllication(config *config.Config) {
	fmt.Println("Starting application...")
	router := initialiseRoutes(config)
	fmt.Println("application started on route http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
	fmt.Println("Starting application...")

}
