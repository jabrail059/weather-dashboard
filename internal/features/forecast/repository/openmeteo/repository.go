package forecast_openmeteo_repository

import "net/http"

type ForecastRepository struct {
	client           *http.Client
	geocodingBaseURL string
	forecastBaseURL  string
}

func NewForecastRepository(
	client *http.Client,
	geocodingBaseURL string,
	forecastBaseURL string,
) *ForecastRepository {
	return &ForecastRepository{
		client:           client,
		geocodingBaseURL: geocodingBaseURL,
		forecastBaseURL:  forecastBaseURL,
	}
}
