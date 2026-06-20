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

func GetDailyForecast(ctx context.Context, latitude float64, longitude float64) (*models.Daily, error, int) {
	var res models.OpenMeteoResponse
	key := fmt.Sprintf("forecast:%v:%v", latitude, longitude)

	value, err := redis.Client().Get(ctx, key).Result()
	if err == nil {
		if err = json.Unmarshal([]byte(value), &res); err != nil {
			return nil, fmt.Errorf("Не удалось получить данные"), http.StatusInternalServerError
		}
		return &res.Daily, nil, http.StatusOK
	}

	apiURL := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%v&longitude=%v&daily=weather_code,temperature_2m_max,temperature_2m_min,sunrise,sunset&hourly=weather_code,temperature_2m&current=temperature_2m,apparent_temperature,wind_speed_10m,weather_code,is_day&timezone=auto", latitude, longitude)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("Не удалось создать запрос"), http.StatusInternalServerError
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Не удалось получить данные"), http.StatusBadRequest
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Возникла ошибка при получении данных"), resp.StatusCode
	}

	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, fmt.Errorf("Не удалось декодировать данные"), http.StatusBadRequest
	}

	jsonRes, err := json.Marshal(res)
	if err != nil {
		return nil, fmt.Errorf("Не удалось обработать данные"), http.StatusInternalServerError
	}

	if err = redis.Client().Set(ctx, key, jsonRes, time.Minute*45).Err(); err != nil {
		return nil, fmt.Errorf("Не удалось кешировать данные"), http.StatusInternalServerError
	}

	return &res.Daily, nil, http.StatusOK
}

func GetHourlyForecast(ctx context.Context, latitude string, longitude string, date string) (*models.Hourly, error, int) {
	var res models.OpenMeteoResponse
	key := fmt.Sprintf("forecast:%s:%s", latitude, longitude)

	value, err := redis.Client().Get(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("Отсутствуют данные о почасовой погоде"), http.StatusNotFound
	}
	if err := json.Unmarshal([]byte(value), &res); err != nil {
		return nil, fmt.Errorf("Не удалось получить данные"), http.StatusInternalServerError
	}
	return &res.Hourly, nil, http.StatusOK
}

func GetCurrentForecast(ctx context.Context, latitude float64, longitude float64) (*models.Current, error, int) {
	var res models.OpenMeteoResponse
	key := fmt.Sprintf("forecast:%v:%v", latitude, longitude)

	value, err := redis.Client().Get(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("Отсутствуют данные о текущей погоде"), http.StatusNotFound
	}
	if err := json.Unmarshal([]byte(value), &res); err != nil {
		return nil, fmt.Errorf("Не удалось получить данные"), http.StatusInternalServerError
	}
	return &res.Current, nil, http.StatusOK
}
