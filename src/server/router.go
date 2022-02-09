package server

import (
	"src/api"

	"github.com/gorilla/mux"
)

func initialiseRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", api.HelloWorldHandler)

	return router
}
