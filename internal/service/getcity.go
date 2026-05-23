package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jabrail059/weather-dashboard/internal/models"
)

func GetCityByIP(ctx context.Context, ip string) (*models.CityReport, error) {
	apiURL := fmt.Sprintf("http://ip-api.com/json/%s?lang=ru&fields=status,city,lat,lon", ip)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("Не удалось создать запрос")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Не удалось получить данные")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Возникла ошибка при получении данных")
	}

	var report models.CityReport
	err = json.NewDecoder(resp.Body).Decode(&report)
	if err != nil {
		return nil, fmt.Errorf("Не удалось получить данные")
	}

	if report.Status != "success" {
		return nil, fmt.Errorf("Не удалось определить город")
	}

	return &report, nil
}
