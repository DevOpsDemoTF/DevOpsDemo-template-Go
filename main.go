package main

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := parseConfig()
	initMetrics()

	log.Info("Service starting", cfg)

	// expose metrics
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(":9102", nil))
	}()

	// expose application and health endpoints
	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").Path("/health").HandlerFunc(Health)

	log.Fatal(http.ListenAndServe(":8080", router))
}
