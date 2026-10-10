package forecast_service

import (
	"context"

	"github.com/jabrail059/weather-dashboard/internal/core/domain"
)

type ForecastService struct {
	forecastRepository ForecastRepository
}

type ForecastRepository interface {
	GeocodeCity(
		ctx context.Context,
		city string,
	) (domain.GeocodeResult, error)

	GetForecast(
		ctx context.Context,
		latitude float64,
		longitude float64,
	) (domain.Forecast, error)

	GetHourlyForecast(
		ctx context.Context,
		latitude float64,
		longitude float64,
	) (domain.OneDayForecast, error)

	// GetCurrentForecast(
	// 	ctx context.Context,
	// 	latitude float64,
	// 	longitude float64,
	// ) (*domain.CurrentForecast, error)
}

func NewForecastService(
	forecastRepository ForecastRepository,
) *ForecastService {
	return &ForecastService{
		forecastRepository: forecastRepository,
	}
}
