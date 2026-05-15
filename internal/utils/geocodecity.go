package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jabrail059/weather-dashboard/internal/models"
	"github.com/jabrail059/weather-dashboard/storage"
)

func GeocodeCity(city string) (*models.GeoRequest, error, int) {
	var geo models.GeoRequest
	key := fmt.Sprintf("geocode:%s", city)

	value, err := storage.Redis().Get(context.Background(), key).Result()
	if err == nil {
		if err = json.Unmarshal([]byte(value), &geo); err != nil {
			return nil, fmt.Errorf("Не удалось получить данные"), http.StatusInternalServerError
		}
		return &geo, nil, http.StatusOK
	}

	resp, err := http.Get(fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=ru&format=json", city))
	if err != nil {
		return nil, fmt.Errorf("Не удалось получить о погоде"), http.StatusBadRequest
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Возникла ошибка при получении данных"), http.StatusBadRequest
	}

	if err = json.NewDecoder(resp.Body).Decode(&geo); err != nil {
		return nil, fmt.Errorf("Не удалось получить данные"), http.StatusBadRequest
	}
	if len(geo.Results) == 0 {
		return nil, fmt.Errorf("Город не найден"), http.StatusNotFound
	}

	jsonGeo, err := json.Marshal(geo)
	if err != nil {
		return nil, fmt.Errorf("Не удалось обработать данные"), http.StatusInternalServerError
	}

	if err = storage.Redis().Set(context.Background(), key, jsonGeo, time.Hour*24).Err(); err != nil {
		return nil, fmt.Errorf("Не удалось кешировать данные"), http.StatusInternalServerError
	}

	return &geo, nil, http.StatusOK
}
