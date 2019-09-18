package main

import (
	"github.com/butzist/DevOpsDemo-template-Go/config"
	"github.com/butzist/DevOpsDemo-template-Go/handlers"
	"github.com/butzist/DevOpsDemo-template-Go/state"
	"github.com/gorilla/mux"
	"log"
)

func createRouter(cfg *config.Config) *mux.Router {
	var handler *handlers.Handler

	if stat, err := state.Create(cfg); err != nil {
		log.Fatal("Error initializing state", err)
	} else {
		handler = handlers.Create(stat)
	}

	// expose application and health endpoints
	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").Path("/health").HandlerFunc(handler.Health)

	return router
}
