package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/jabrail059/weather-dashboard/internal/models"
	"github.com/jabrail059/weather-dashboard/storage"
)

func GeocodeCity(cityStorage storage.Storage, city string) (*models.GeoRequest, error, int) {
	var geo models.GeoRequest
	cityName := strings.TrimSpace(strings.ToLower(city))

	result, err := cityStorage.Select(context.Background(), cityName)
	if err == nil {
		geo = models.GeoRequest{
			Results: []models.Result{*result},
		}
		return &geo, nil, http.StatusOK
	}
	if err != storage.ErrCityNotSaved {
		return nil, fmt.Errorf("Не удалось получить данные"), http.StatusInternalServerError
	}

	resp, err := http.Get(fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=ru&format=json", url.QueryEscape(cityName)))
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
	geo.Results[0].Name = cityName
	err = cityStorage.Save(context.Background(), &geo.Results[0])
	if err != nil {
		return nil, fmt.Errorf("Не удалось сохранить данные"), http.StatusInternalServerError
	}

	return &geo, nil, http.StatusOK
}
