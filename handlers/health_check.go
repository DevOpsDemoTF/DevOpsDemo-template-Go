package handlers

import (
	"github.com/butzist/DevOpsDemo-template-Go/metrics"
	"net/http"
)

func (h *Handler) Health(w http.ResponseWriter, _ *http.Request) {
	metrics.HealthCounter.Inc()

	status := http.StatusOK
	if !h.state.Healthy {
		status = http.StatusInternalServerError
	}
	w.WriteHeader(status)
}
