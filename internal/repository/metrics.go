package repository

import (
	"context"
	"order-service/internal/observability"
	"time"
)

func observeQuery[Generic any](ctx context.Context,
	operation string, fn func() (Generic, error)) (Generic, error) {

	span_name := "db." + operation
	_, span := observability.Tracer.Start(ctx, span_name)
	defer span.End()

	start := time.Now()
	result, err := fn()

	if err != nil {
		span.RecordError(err)
	}

	observability.DataBaseQueryDuration.WithLabelValues(operation).Observe(time.Since(start).Seconds())
	return result, err
}
