package server

import (
	"fmt"
	"log"
	"net/http"
)

func StartApllication() {
	fmt.Println("Starting application...")
	router := initialiseRoutes()
	fmt.Println("application started on route http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
	fmt.Println("Starting application...")

}
