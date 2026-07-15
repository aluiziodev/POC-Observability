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
)

func RequestIdFromContext(ctx context.Context) string {
	if id, ok := ctx.Value(domain.RequestIdKey).(string); ok {
		return id
	}

	return ""
}

func RequestLogger(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			reqID := newRequestID()
			ctx := context.WithValue(r.Context(), domain.RequestIdKey, reqID)
			r = r.WithContext(ctx)

			rec := &models.ResponseRecorder{Writer: w, Status: http.StatusOK}

			next.ServeHTTP(rec, r)

			duration := time.Since(start)

			path := routePattern(r)

			logger.Info("http_request",
				slog.String("request_id", reqID),
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
