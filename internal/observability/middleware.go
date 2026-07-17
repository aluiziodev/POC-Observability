package observability

import (
	"context"
	"log/slog"
	"net/http"
	"order-service/internal/domain"
	"order-service/internal/models"
	"strconv"
	"strings"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
)

func RequestLogger(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			propagated_ctx := otel.GetTextMapPropagator().Extract(r.Context(),
				propagation.HeaderCarrier(r.Header))

			ctx, span := Tracer.Start(propagated_ctx, r.Method+" "+r.URL.Path)
			defer span.End()

			reqID := newRequestID()
			ctx = context.WithValue(ctx, domain.RequestIdKey, reqID)
			r = r.WithContext(ctx)

			rec := &models.ResponseRecorder{Writer: w, Status: http.StatusOK}

			next.ServeHTTP(rec, r)

			duration := time.Since(start)

			path := routePattern(r)

			span_name := r.Method + " " + path
			span.SetName(span_name)
			span.SetAttributes(
				attribute.String("http_method", r.Method),
				attribute.String("http_route", path),
				attribute.Int("http_status", rec.Status),
			)

			if rec.Status >= 500 {
				span.SetStatus(codes.Error, http.StatusText(rec.Status))
			}

			logger.Info("http_request",
				slog.String("request_id", reqID),
				slog.String("trace_id", TraceIdFromContext(ctx)),
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Int("status", rec.Status),
				slog.Duration("duration", duration),
			)

			status := strconv.Itoa(rec.Status)
			HttpRequestsTotal.WithLabelValues(r.Method, path, status).Inc()
			httpRequestDuration.WithLabelValues(r.Method, path).Observe(duration.Seconds())
		})
	}
}

func routePattern(r *http.Request) string {
	if r.Pattern == "" {
		return r.URL.Path
	}

	if _, pattern, found := strings.Cut(r.Pattern, " "); found {
		return pattern
	}

	return r.Pattern
}
