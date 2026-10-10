package core_http_server

import (
	"net/http"

	core_http_middleware "github.com/jabrail059/weather-dashboard/internal/core/transport/http/middleware"
)

type Route struct {
	Method     string
	Path       string
	Handler    http.HandlerFunc
	Middleware []core_http_middleware.Middleware
}

func (r *Route) WithMiddleware() http.Handler {
	return core_http_middleware.ChainMiddlerware(
		r.Handler,
		r.Middleware...,
	)
}
