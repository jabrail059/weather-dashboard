package web_transport_http

import (
	"context"
	"net/http"

	"github.com/jabrail059/weather-dashboard/internal/core/domain"
	core_http_server "github.com/jabrail059/weather-dashboard/internal/core/transport/http/server"
)

type WebHTTPHandler struct {
	webService WebService
}

type WebService interface {
	GetCityByIP(
		ctx context.Context,
		ip string,
	) (domain.CityReport, error)

	GetSearchHistory(
		ctx context.Context,
		value string,
	) (domain.HistoryList, error)
}

func NewWebHTTPHadnler(
	webService WebService,
) *WebHTTPHandler {
	return &WebHTTPHandler{
		webService: webService,
	}
}

func (h *WebHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/",
			Handler: h.GetMainPage,
		},
		{
			Method:  http.MethodGet,
			Path:    "/history",
			Handler: h.GetSearchHistory,
		},
	}
}
