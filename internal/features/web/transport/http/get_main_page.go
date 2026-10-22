package web_transport_http

import (
	"net"
	"net/http"

	core_logger "github.com/jabrail059/weather-dashboard/internal/core/logger"
	core_http_response "github.com/jabrail059/weather-dashboard/internal/core/transport/http/response"
	core_http_session "github.com/jabrail059/weather-dashboard/internal/core/transport/http/session"
)

func (h *WebHTTPHandler) GetMainPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		ip = r.RemoteAddr
	}

	data, err := h.webService.GetCityByIP(ctx, ip)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get city by ip",
		)

		return
	}

	cookie, isNew := core_http_session.GetOrCreate(r)
	if isNew {
		http.SetCookie(w, cookie)
	}

	if err := responseHandler.RenderHTML(data, []string{"index.html"}); err != nil {
		responseHandler.ErrorResponse(
			err,
			"render main page HTML template",
		)

		return
	}
}
