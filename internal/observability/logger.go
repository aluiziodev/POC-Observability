package observability

import (
	"context"
	"log/slog"
	"os"
)

func NewLogger() *slog.Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	logger := slog.New(handler)
	slog.SetDefault(logger)

	return logger
}

func WarnContext(ctx context.Context, msg string, err error) {
	slog.WarnContext(ctx, msg,
		slog.String("request_id", requestIdFromContext(ctx)),
		slog.String("error", err.Error()))
}

func ErrorContext(ctx context.Context, msg string, err error) {
	slog.ErrorContext(ctx, msg,
		slog.String("request_id", requestIdFromContext(ctx)),
		slog.String("error", err.Error()))
}
