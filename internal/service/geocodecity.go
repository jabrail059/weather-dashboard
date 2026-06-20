package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"unicode"

	"github.com/jabrail059/weather-dashboard/internal/models"
	"github.com/jabrail059/weather-dashboard/internal/storage"
)

func GeocodeCity(ctx context.Context, cityStorage storage.Storage, city string) (*models.GeoRequest, error, int) {
	var geo models.GeoRequest
	cityName := normalizeCityName(city)

	result, err := cityStorage.Select(ctx, cityName)
	if err == nil {
		geo = models.GeoRequest{
			Results: []models.Result{*result},
		}
		return &geo, nil, http.StatusOK
	}
	if err != storage.ErrCityNotSaved {
		return nil, fmt.Errorf("Не удалось получить данные"), http.StatusInternalServerError
	}

	apiURL := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=ru&format=json", url.QueryEscape(cityName))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("Не удалось создать запрос"), http.StatusInternalServerError
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Не удалось получить о погоде"), http.StatusBadRequest
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Возникла ошибка при получении данных"), resp.StatusCode
	}

	if err = json.NewDecoder(resp.Body).Decode(&geo); err != nil {
		return nil, fmt.Errorf("Не удалось получить данные"), http.StatusBadRequest
	}
	if len(geo.Results) == 0 {
		return nil, fmt.Errorf("Город не найден"), http.StatusNotFound
	}
	geo.Results[0].Name = cityName
	err = cityStorage.Save(ctx, &geo.Results[0])
	if err != nil {
		return nil, fmt.Errorf("Не удалось сохранить данные"), http.StatusInternalServerError
	}

	return &geo, nil, http.StatusOK
}

func normalizeCityName(city string) string {
	city = strings.TrimSpace(city)
	city = strings.ToLower(city)

	if city == "" {
		return city
	}

	runes := []rune(city)

	makeUpper := true
	for i, r := range runes {
		if makeUpper && unicode.IsLetter(r) {
			runes[i] = unicode.ToUpper(r)
			makeUpper = false
			continue
		}

		if r == ' ' || r == '-' {
			makeUpper = true
		}
	}

	return string(runes)
}
