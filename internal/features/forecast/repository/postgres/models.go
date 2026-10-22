package forecast_postgres_repository

import "github.com/jabrail059/weather-dashboard/internal/core/domain"

type GeocodeModel struct {
	ID        int
	City      string
	Latitude  float64
	Longitude float64
}

func geocodeDomainFromModel(
	geocodeModel GeocodeModel,
) domain.GeocodeResult {
	return domain.NewGeocodeResult(
		geocodeModel.ID,
		geocodeModel.City,
		geocodeModel.Latitude,
		geocodeModel.Longitude,
	)
}
