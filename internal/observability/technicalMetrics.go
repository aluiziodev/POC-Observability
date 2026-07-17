package observability

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	HttpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total de requisiçoes http recebidas",
		},
		[]string{"method", "path", "status"},
	)

	httpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration",
			Help:    "Histograma de duraçao das requsiçoes http",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)

	DataBaseQueryDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "database_query_duration",
			Help:    "Duraçao das consultas ao banco de dados",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"operation"},
	)
)
