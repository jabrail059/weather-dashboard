package forecast_transport_http

import (
	"net/http"

	core_logger "github.com/jabrail059/weather-dashboard/internal/core/logger"
	core_http_request "github.com/jabrail059/weather-dashboard/internal/core/transport/http/request"
	core_http_response "github.com/jabrail059/weather-dashboard/internal/core/transport/http/response"
)

func (h *ForecastHTTPHandler) GetForecast(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	if !isHTMXRequest(r) {
		responseHandler.HTMXMethodErrorResponse()

		return
	}

	city := GetCityQueryParam(r)

	forecast, err := h.forecastService.GetForecast(ctx, city)

	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get forecast",
		)

		return
	}

	//TODO: передавать строковые ключи - плохо, придумать что нибудь с этим
	if err := responseHandler.RenderPartHTML(forecast, "base.html"); err != nil {
		responseHandler.ErrorResponse(
			err,
			"render part of HTML template",
		)

		return
	}
}

func GetCityQueryParam(r *http.Request) *string {
	var (
		cityQueryParamKey = "city"
	)
	city := core_http_request.GetStringQueryParam(r, cityQueryParamKey)

	return city
}

func isHTMXRequest(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true"
}
