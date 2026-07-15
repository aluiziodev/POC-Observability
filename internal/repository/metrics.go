package repository

import (
	"order-service/internal/observability"
	"time"
)

func observeQuery[Generic any](operation string, fn func() (Generic, error)) (Generic, error) {
	start := time.Now()
	result, err := fn()
	observability.DataBaseQueryDuration.WithLabelValues(operation).Observe(time.Since(start).Seconds())
	return result, err
}
