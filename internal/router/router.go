package router

import (
	"log/slog"
	"net/http"
	"order-service/internal/observability"
)

func GenerateRouter(logger *slog.Logger) http.Handler {
	mux := http.NewServeMux()
	configure(mux)
	return observability.RequestLogger(logger)(mux)
}

func configure(mux *http.ServeMux) {
	for _, route := range routes {
		mux.HandleFunc(route.method+" "+route.URI, route.function)
	}
}
