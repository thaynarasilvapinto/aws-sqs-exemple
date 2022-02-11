package server

import (
	"src/api"
	"src/config"

	"github.com/gorilla/mux"
)

func initialiseRoutes(config *config.Config) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", api.CreateMessage(config)).Methods("POST")

	return router
}
