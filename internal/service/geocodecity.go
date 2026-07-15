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

func GeocodeCity(ctx context.Context, cityStorage storage.Storage, city string) (*models.GeoRequest, int, error) {
	var geo models.GeoRequest
	cityName := normalizeCityName(city)

	result, err := cityStorage.Select(ctx, cityName)
	if err == nil {
		geo = models.GeoRequest{
			Results: []models.Result{*result},
		}
		return &geo, http.StatusOK, nil
	}
	if err != storage.ErrCityNotSaved {
		return nil, http.StatusInternalServerError, fmt.Errorf("не удалось получить данные")
	}

	apiURL := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=ru&format=json", url.QueryEscape(cityName))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("не удалось создать запрос")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("не удалось получить о погоде")
	}
	defer resp.Body.Close() //nolint:errcheck

	if resp.StatusCode != http.StatusOK {
		return nil, resp.StatusCode, fmt.Errorf("возникла ошибка при получении данных")
	}

	if err = json.NewDecoder(resp.Body).Decode(&geo); err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("не удалось получить данные")
	}
	if len(geo.Results) == 0 {
		return nil, http.StatusNotFound, fmt.Errorf("город не найден")
	}
	geo.Results[0].Name = cityName
	err = cityStorage.Save(ctx, &geo.Results[0])
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("не удалось сохранить данные")
	}

	return &geo, http.StatusOK, nil
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
