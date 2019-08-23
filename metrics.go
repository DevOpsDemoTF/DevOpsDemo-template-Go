package main

import "github.com/prometheus/client_golang/prometheus"

var healthCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "health_counter",
	Help: "Number of times the health endpoint has been called",
})

func initMetrics() {
	prometheus.MustRegister(healthCounter)
}
