package forecast_transport_http

import (
	"context"
	"net/http"
	"time"

	"github.com/jabrail059/weather-dashboard/internal/core/domain"
	core_http_server "github.com/jabrail059/weather-dashboard/internal/core/transport/http/server"
)

type ForecastHTTPHandler struct {
	forecastService ForecastService
}

type ForecastService interface {
	GetForecast(
		ctx context.Context,
		city *string,
	) (domain.Forecast, error)

	GetHourlyForecast(
		ctx context.Context,
		date *time.Time,
		city *string,
		sunrise *time.Time,
		sunset *time.Time,
	) (domain.HourlyRender, error)
}

func NewForecastHTTPHandler(
	forecastService ForecastService,
) *ForecastHTTPHandler {
	return &ForecastHTTPHandler{
		forecastService: forecastService,
	}
}

func (h *ForecastHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/weather",
			Handler: h.GetForecast,
		},
		{
			Method:  http.MethodGet,
			Path:    "/weather/hourly",
			Handler: h.GetHourlyForecast,
		},
	}
}
