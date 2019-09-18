package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var HealthCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "health_counter",
	Help: "Number of times the health endpoint has been called",
})

func InitMetrics() {
	prometheus.MustRegister(HealthCounter)

	// expose metrics
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(":9102", nil))
	}()
}
