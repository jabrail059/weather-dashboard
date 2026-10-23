package forecast_service

import (
	"context"
	"errors"
	"fmt"

	"github.com/jabrail059/weather-dashboard/internal/core/domain"
	core_redis_client "github.com/jabrail059/weather-dashboard/internal/core/repository/redis/client"
	core_errors "github.com/jabrail059/weather-dashboard/internal/core/transport/http/errors"
)

func (s *ForecastService) GetForecast(
	ctx context.Context,
	city *string,
	value string,
) (domain.Forecast, error) {
	if city == nil {
		return domain.Forecast{}, fmt.Errorf(
			"city must be not NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	geocodeCity, err := s.forecastPostgresRepository.GetCity(ctx, *city)
	cityNeedsSaving := errors.Is(err, core_errors.ErrNotFound)
	if err != nil {
		if !cityNeedsSaving {
			return domain.Forecast{}, fmt.Errorf("get city from postgres: %w", err)
		}
		geocodeCity, err = s.forecastOpenmeteoRepository.GeocodeCity(ctx, *city)
		if err != nil {
			return domain.Forecast{}, fmt.Errorf(
				"get city geocode from repository: %w",
				err,
			)
		}
	}

	latitude := geocodeCity.Latitude
	longitude := geocodeCity.Longitude

	var cacheErr error
	forecast, err := s.forecastRedisRepository.GetForecast(
		ctx, latitude, longitude,
	)

	needsCaching := err != nil

	if err != nil {
		if !errors.Is(err, core_redis_client.ErrKeyNotFound) {
			cacheErr = fmt.Errorf("get forecast from cache: %w", err)
		}
		forecast, err = s.forecastOpenmeteoRepository.GetForecast(
			ctx,
			latitude,
			longitude,
		)
		if err != nil {
			return domain.Forecast{CacheErr: cacheErr}, fmt.Errorf(
				"get forecast from repository: %w",
				err,
			)
		}
	}

	forecast = domain.NewForecast(
		geocodeCity.Name,
		latitude,
		longitude,
		forecast.CurrentForecast,
		forecast.Days,
	)
	forecast.CacheErr = cacheErr

	if cityNeedsSaving {
		if err := s.forecastPostgresRepository.SaveCity(ctx, *city, latitude, longitude); err != nil {
			return forecast, fmt.Errorf("postgres error: %w", err)
		}
	}

	if needsCaching {
		if err := s.forecastRedisRepository.AddForecast(ctx, forecast); err != nil {
			cacheErr = errors.Join(cacheErr, fmt.Errorf("save forecast in cache: %w", err))
		}
	}

	if err := s.forecastRedisRepository.AddSearchHistory(ctx, *city, value); err != nil {
		cacheErr = errors.Join(cacheErr, fmt.Errorf("save search history in cache: %w", err))
	}
	forecast.CacheErr = cacheErr

	return forecast, nil
}
