package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jabrail059/weather-dashboard/internal/models"
	"github.com/jabrail059/weather-dashboard/storage"
)

func GetForecast(latitude float64, longitude float64) (*models.ReqDaily, error, int) {
	var res models.ReqDaily
	key := fmt.Sprintf("forecast:%.4f:%.4f", latitude, longitude)

	value, err := storage.Redis().Get(context.Background(), key).Result()
	if err == nil {
		if err = json.Unmarshal([]byte(value), &res); err != nil {
			return nil, fmt.Errorf("Не удалось получить данные"), http.StatusInternalServerError
		}
		return &res, nil, http.StatusOK
	}

	resp, err := http.Get(fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%v&longitude=%v&daily=weather_code,temperature_2m_max,temperature_2m_min&timezone=auto", latitude, longitude))
	if err != nil {
		return nil, fmt.Errorf("Не удалось получить данные"), http.StatusBadRequest
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Возникла ошибка при получении данных"), http.StatusBadRequest
	}

	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, fmt.Errorf("Не удалось получить данные"), http.StatusBadRequest
	}

	jsonRes, err := json.Marshal(res)
	if err != nil {
		return nil, fmt.Errorf("Не удалось обработать данные"), http.StatusInternalServerError
	}

	if err = storage.Redis().Set(context.Background(), key, jsonRes, time.Minute*45).Err(); err != nil {
		return nil, fmt.Errorf("Не удалось кешировать данные"), http.StatusInternalServerError
	}

	return &res, nil, http.StatusOK
}
