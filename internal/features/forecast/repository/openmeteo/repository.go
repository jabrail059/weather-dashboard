package forecast_openmeteo_repository

import "net/http"

type ForecastOpenmeteoRepository struct {
	client           *http.Client
	geocodingBaseURL string
	forecastBaseURL  string
}

func NewForecastOpenmeteoRepository(
	client *http.Client,
	geocodingBaseURL string,
	forecastBaseURL string,
) *ForecastOpenmeteoRepository {
	return &ForecastOpenmeteoRepository{
		client:           client,
		geocodingBaseURL: geocodingBaseURL,
		forecastBaseURL:  forecastBaseURL,
	}
}
