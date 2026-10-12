package forecast_transport_http

import (
	"fmt"
	"net/http"
	"time"

	core_logger "github.com/jabrail059/weather-dashboard/internal/core/logger"
	core_http_request "github.com/jabrail059/weather-dashboard/internal/core/transport/http/request"
	core_http_response "github.com/jabrail059/weather-dashboard/internal/core/transport/http/response"
)

func (h *ForecastHTTPHandler) GetHourlyForecast(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	queryParams, err := getQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get date/lat/lon/city/sunrise/sunset query param",
		)

		return
	}

	forecast, err := h.forecastService.GetHourlyForecast(
		ctx,
		queryParams.date,
		queryParams.city,
		queryParams.sunrise,
		queryParams.sunset,
	)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get hourly forecast",
		)

		return
	}

	files := []string{"hourly.html", "hourly_base.html"}
	if err := responseHandler.RenderHTML(forecast, files); err != nil {
		responseHandler.ErrorResponse(
			err,
			"render hourly forecast HTML template",
		)

		return
	}
}

type queryParams struct {
	date    *time.Time
	city    *string
	sunrise *time.Time
	sunset  *time.Time
}

func getQueryParams(r *http.Request) (*queryParams, error) {
	var (
		dateQueryParamKey    = "date"
		cityQueryParamKey    = "city"
		sunriseQueryParamKey = "sunrise"
		sunsetQueryParamKey  = "sunset"
	)

	date, err := core_http_request.GetTimeQueryParam(r, dateQueryParamKey, "2006-01-02")
	if err != nil {
		return nil, fmt.Errorf("get 'date' query param: %w", err)
	}

	city := core_http_request.GetStringQueryParam(r, cityQueryParamKey)

	sunrise, err := core_http_request.GetTimeQueryParam(r, sunriseQueryParamKey, "15:04")
	if err != nil {
		return nil, fmt.Errorf("get 'sunrise' query param: %w", err)
	}
	sunset, err := core_http_request.GetTimeQueryParam(r, sunsetQueryParamKey, "15:04")
	if err != nil {
		return nil, fmt.Errorf("get 'sunset' query param: %w", err)
	}

	return &queryParams{
		date:    date,
		city:    city,
		sunrise: sunrise,
		sunset:  sunset,
	}, nil
}
