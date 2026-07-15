package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jabrail059/weather-dashboard/internal/models"
	"github.com/jabrail059/weather-dashboard/internal/storage/redis"
)

func GetDailyForecast(ctx context.Context, latitude float64, longitude float64) (*models.Daily, int, error) {
	var res models.OpenMeteoResponse
	key := fmt.Sprintf("forecast:%v:%v", latitude, longitude)

	value, err := redis.Client().Get(ctx, key).Result()
	if err == nil {
		if err = json.Unmarshal([]byte(value), &res); err != nil {
			return nil, http.StatusInternalServerError, fmt.Errorf("не удалось получить данные")
		}
		return &res.Daily, http.StatusOK, nil
	}

	apiURL := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%v&longitude=%v&daily=weather_code,temperature_2m_max,temperature_2m_min,sunrise,sunset&hourly=weather_code,temperature_2m&current=temperature_2m,apparent_temperature,wind_speed_10m,weather_code,is_day&timezone=auto", latitude, longitude)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("не удалось создать запрос")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("не удалось получить данные")
	}
	defer resp.Body.Close() //nolint:errcheck

	if resp.StatusCode != http.StatusOK {
		return nil, resp.StatusCode, fmt.Errorf("возникла ошибка при получении данных")
	}

	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("не удалось декодировать данные")
	}

	jsonRes, err := json.Marshal(res)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("не удалось обработать данные")
	}

	if err = redis.Client().Set(ctx, key, jsonRes, time.Minute*45).Err(); err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("не удалось кешировать данные")
	}

	return &res.Daily, http.StatusOK, nil
}

func GetHourlyForecast(ctx context.Context, latitude string, longitude string, date string) (*models.Hourly, int, error) {
	var res models.OpenMeteoResponse
	key := fmt.Sprintf("forecast:%s:%s", latitude, longitude)

	value, err := redis.Client().Get(ctx, key).Result()
	if err != nil {
		return nil, http.StatusNotFound, fmt.Errorf("отсутствуют данные о почасовой погоде")
	}
	if err := json.Unmarshal([]byte(value), &res); err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("не удалось получить данные")
	}
	return &res.Hourly, http.StatusOK, nil
}

func GetCurrentForecast(ctx context.Context, latitude float64, longitude float64) (*models.Current, int, error) {
	var res models.OpenMeteoResponse
	key := fmt.Sprintf("forecast:%v:%v", latitude, longitude)

	value, err := redis.Client().Get(ctx, key).Result()
	if err != nil {
		return nil, http.StatusNotFound, fmt.Errorf("отсутствуют данные о текущей погоде")
	}
	if err := json.Unmarshal([]byte(value), &res); err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("не удалось получить данные")
	}
	return &res.Current, http.StatusOK, nil
}
