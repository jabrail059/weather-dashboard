package forecast_openmeteo_repository

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jabrail059/weather-dashboard/internal/core/domain"
	core_http_utils "github.com/jabrail059/weather-dashboard/internal/core/transport/utils"
)

const (
	hourlURLKey               = "hourly"
	hourlyForecastQueryParams = "weather_code,temperature_2m"
)

func (r *ForecastOpenmeteoRepository) GetHourlyForecast(
	ctx context.Context,
	latitude float64,
	longitude float64,
) (domain.OneDayForecast, error) {
	var hourlyResponse HourlyResponse

	params := map[string]string{
		"latitude":  strconv.FormatFloat(latitude, 'f', -1, 64),
		"longitude": strconv.FormatFloat(longitude, 'f', -1, 64),
		"timezone":  defaultTimezone,
		hourlURLKey: hourlyForecastQueryParams,
	}

	apiURL, err := core_http_utils.BuildURL(
		r.forecastBaseURL,
		params,
	)
	if err != nil {
		return domain.OneDayForecast{}, fmt.Errorf("build forecast url: %w", err)
	}

	err = r.getResponseFromOpenMeteo(
		ctx,
		apiURL,
		&hourlyResponse,
	)
	if err != nil {
		return domain.OneDayForecast{}, fmt.Errorf("get response from openmeteo: %w", err)
	}

	oneDayForecast := hourlyForecastFromResponse(hourlyResponse)
	return oneDayForecast, nil
}
