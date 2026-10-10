package forecast_openmeteo_repository

import (
	"context"
	"fmt"

	"github.com/jabrail059/weather-dashboard/internal/core/domain"
)

const (
	dailyURLKey              = "daily"
	dailyForecastQueryParams = "weather_code,temperature_2m_max,temperature_2m_min,sunrise,sunset"
	currentURLKey            = "current"
	currentURLQueryParams    = "temperature_2m,apparent_temperature,wind_speed_10m,weather_code,is_day"
)

func (r *ForecastRepository) GetForecast(
	ctx context.Context,
	latitude float64,
	longitude float64,
) (domain.Forecast, error) {
	var openMeteoResponse OpenMeteoResponse
	//key := fmt.Sprintf("forecast:%v:%v", latitude, longitude)

	// value, err := redis.Client().Get(ctx, key).Result()
	// if err == nil {
	// 	if err = json.Unmarshal([]byte(value), &res); err != nil {
	// 		return nil, http.StatusInternalServerError, fmt.Errorf("не удалось получить данные")
	// 	}
	// 	return &res.Daily, http.StatusOK, nil
	// }

	forecastParams := map[string]string{
		dailyURLKey:   dailyForecastQueryParams,
		currentURLKey: currentURLQueryParams,
	}

	apiURL, err := r.buildForecastURL(
		latitude,
		longitude,
		forecastParams,
	)
	if err != nil {
		return domain.Forecast{}, fmt.Errorf("build forecast url: %w", err)
	}

	err = r.getResponseFromOpenMeteo(
		ctx,
		apiURL,
		&openMeteoResponse,
	)
	if err != nil {
		return domain.Forecast{}, fmt.Errorf("get response from openmeteo: %w", err)
	}

	forecast, err := forecastFromResponse(openMeteoResponse)
	if err != nil {
		return domain.Forecast{}, fmt.Errorf("get daily forecast: %w", err)
	}

	return forecast, nil

	// jsonRes, err := json.Marshal(openMeteoResponse)
	// if err != nil {
	// 	return nil, http.StatusInternalServerError, fmt.Errorf("не удалось обработать данные")
	// }

	// if err = redis.Client().Set(ctx, key, jsonRes, time.Minute*45).Err(); err != nil {
	// 	return nil, http.StatusInternalServerError, fmt.Errorf("не удалось кешировать данные")
	// }
}
