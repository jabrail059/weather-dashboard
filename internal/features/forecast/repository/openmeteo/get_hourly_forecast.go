package forecast_openmeteo_repository

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jabrail059/weather-dashboard/internal/core/domain"
	core_http_url "github.com/jabrail059/weather-dashboard/internal/core/transport/url"
)

const (
	hourlURLKey               = "hourly"
	hourlyForecastQueryParams = "weather_code,temperature_2m"
)

func (r *ForecastRepository) GetHourlyForecast(
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

	apiURL, err := core_http_url.BuildURL(
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
	// key := fmt.Sprintf("forecast:%s:%s", latitude, longitude)

	// value, err := redis.Client().Get(ctx, key).Result()
	// if err != nil {
	// 	return nil, http.StatusNotFound, fmt.Errorf("отсутствуют данные о почасовой погоде")
	// }
	// if err := json.Unmarshal([]byte(value), &openMeteoResponse); err != nil {
	// 	return nil, http.StatusInternalServerError, fmt.Errorf("не удалось получить данные")
	// }
	return oneDayForecast, nil
}
