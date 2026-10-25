package web_api_repository

import (
	"context"
	"fmt"

	"github.com/jabrail059/weather-dashboard/internal/core/domain"
	core_http_utils "github.com/jabrail059/weather-dashboard/internal/core/transport/http/utils"
)

type CityReport struct {
	Status string `json:"status"`
	City   string `json:"city"`
}

func (r *WebApiRepository) GetCityByIP(
	ctx context.Context,
	ip string,
) (domain.CityReport, error) {
	params := map[string]string{
		"lang":   "ru",
		"fields": "status,city",
	}
	getIpURL := fmt.Sprintf("%s/%s", r.getIpURL, ip)

	apiURL, err := core_http_utils.BuildURL(getIpURL, params)
	if err != nil {
		return domain.CityReport{}, fmt.Errorf(
			"build ip-api url: %w",
			err,
		)
	}

	var response CityReport
	err = r.getResponseFromIpApi(
		ctx,
		apiURL,
		&response,
	)

	var report domain.CityReport
	if err == nil && response.Status == "success" {
		report.City = response.City
		report.Detected = true
	}

	return report, nil
}
