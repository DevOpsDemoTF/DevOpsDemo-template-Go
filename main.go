package main

import (
	"github.com/butzist/DevOpsDemo-template-Go/config"
	"github.com/butzist/DevOpsDemo-template-Go/metrics"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := config.ParseConfig()
	metrics.InitMetrics()

	log.Info("Service has been started", cfg)

	router := createRouter(cfg)
	log.Fatal(http.ListenAndServe(":8080", router))
}
