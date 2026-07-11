package observability

import (
	"context"
	"log/slog"
	"net/http"
	"order-service/internal/domain"
	"order-service/internal/models"
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

			next.ServeHTTP(rec.Writer, r)

			logger.Info("http_request",
				slog.String("request_id", reqID),
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Int("status", rec.Status),
				slog.Duration("duration", time.Since(start)),
			)
		})
	}
}
