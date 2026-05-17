package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jabrail059/weather-dashboard/internal/models"
)

func GetCityByIP(ip string) (*models.CityReport, error) {
	resp, err := http.Get(fmt.Sprintf("http://ip-api.com/json/%s?lang=ru&fields=status,city,lat,lon", ip))
	if err != nil {
		return nil, fmt.Errorf("Не удалось получить данные")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
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
