package main

import "net/http"

func Health(w http.ResponseWriter, _ *http.Request) {
	healthCounter.Inc()
	w.WriteHeader(http.StatusOK)
}
