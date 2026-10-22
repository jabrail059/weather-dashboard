package web_transport_http

import (
	"net/http"

	core_logger "github.com/jabrail059/weather-dashboard/internal/core/logger"
	core_http_response "github.com/jabrail059/weather-dashboard/internal/core/transport/http/response"
	core_http_session "github.com/jabrail059/weather-dashboard/internal/core/transport/http/session"
)

func (h *WebHTTPHandler) GetSearchHistory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	cookie, isNew := core_http_session.GetOrCreate(r)
	if isNew {
		http.SetCookie(w, cookie)
	}
	value := cookie.Value

	history, err := h.webService.GetSearchHistory(ctx, value)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get search history",
		)

		return
	}

	if history.CacheErr != nil {
		log.WarnError("cache error", history.CacheErr)
	}

	if err := responseHandler.RenderPartHTML(history, "search_history.html"); err != nil {
		responseHandler.ErrorResponse(
			err,
			"render history HTML template",
		)

		return
	}
}
