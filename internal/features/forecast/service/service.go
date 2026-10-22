package forecast_service

import (
	"context"

	"github.com/jabrail059/weather-dashboard/internal/core/domain"
)

type ForecastService struct {
	forecastOpenmeteoRepository ForecastOpenmeteoRepository
	forecastPostgresRepository  ForecastPostgresRepository
	forecastRedisRepository     ForecastRedisRepository
}

type ForecastOpenmeteoRepository interface {
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
}

type ForecastPostgresRepository interface {
	GetCity(
		ctx context.Context,
		city string,
	) (domain.GeocodeResult, error)

	SaveCity(
		ctx context.Context,
		city string,
		latitude float64,
		longitude float64,
	) error
}

type ForecastRedisRepository interface {
	AddForecast(
		ctx context.Context,
		forecast domain.Forecast,
	) error

	GetForecast(
		ctx context.Context,
		latitude float64,
		longitude float64,
	) (domain.Forecast, error)

	AddSearchHistory(
		ctx context.Context,
		city string,
		value string,
	) error
}

func NewForecastService(
	forecastOpenmeteoRepository ForecastOpenmeteoRepository,
	forecastPostgresRepository ForecastPostgresRepository,
	forecastRedisRepository ForecastRedisRepository,
) *ForecastService {
	return &ForecastService{
		forecastOpenmeteoRepository: forecastOpenmeteoRepository,
		forecastPostgresRepository:  forecastPostgresRepository,
		forecastRedisRepository:     forecastRedisRepository,
	}
}
