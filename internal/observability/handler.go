package observability

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var metricsHandler = promhttp.Handler()

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	metricsHandler.ServeHTTP(w, r)
}
