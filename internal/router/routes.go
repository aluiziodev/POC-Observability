package router

import (
	"net/http"
	"order-service/internal/handler"
	"order-service/internal/observability"
)

type route struct {
	URI      string
	method   string
	function func(http.ResponseWriter, *http.Request)
}

var routes = []route{
	{
		URI:      "/health",
		method:   http.MethodGet,
		function: handler.Health,
	},
	{
		URI:      "/ready",
		method:   http.MethodGet,
		function: handler.Ready,
	},
	{
		URI:      "/metrics",
		method:   http.MethodGet,
		function: observability.MetricsHandler,
	},
	{
		URI:      "/orders",
		method:   http.MethodPost,
		function: handler.Create,
	},
	{
		URI:      "/orders",
		method:   http.MethodGet,
		function: handler.List,
	},
	{
		URI:      "/orders/{id}",
		method:   http.MethodGet,
		function: handler.Get,
	},
	{
		URI:      "/orders/{id}/status",
		method:   http.MethodPatch,
		function: handler.Update,
	},
}
