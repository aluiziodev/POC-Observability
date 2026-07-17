package observability

import (
	"context"
	"log/slog"
	"order-service/internal/utils"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
	"go.opentelemetry.io/otel/trace"
)

var Tracer trace.Tracer = otel.Tracer("order-service")

func InitTracer(ctx context.Context, service_name string) (shutdown func(context.Context) error, err error) {
	endpoint := utils.GetEnvOrDefault("OTEL_EXPORTER_OTLP_ENDPOINT", "localhost:4318")

	exporter, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint(endpoint),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	resource, err := resource.New(ctx, resource.WithAttributes(
		semconv.ServiceName(service_name),
	))
	if err != nil {
		return nil, err
	}

	trace_provider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource),
	)

	otel.SetTracerProvider(trace_provider)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return trace_provider.Shutdown, nil
}

func TraceIdFromContext(ctx context.Context) string {
	span_context := trace.SpanContextFromContext(ctx)
	if !span_context.IsValid() {
		return ""
	}
	return span_context.TraceID().String()
}

func ShutDownTracer(shutDownTracer func(context.Context) error, logger *slog.Logger, timeout time.Duration) {
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := shutDownTracer(shutdownCtx); err != nil {
		logger.Error("Erro ao encerrar o tracer!", slog.String("error", err.Error()))
	}
}
