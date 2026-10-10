package forecast_service

import (
	"context"
	"fmt"

	"github.com/jabrail059/weather-dashboard/internal/core/domain"
	core_errors "github.com/jabrail059/weather-dashboard/internal/core/transport/http/errors"
)

func (s *ForecastService) GetForecast(
	ctx context.Context,
	city *string,
) (domain.Forecast, error) {
	if city == nil {
		return domain.Forecast{}, fmt.Errorf(
			"city must be not NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	geocodeCity, err := s.forecastRepository.GeocodeCity(ctx, *city)
	if err != nil {
		return domain.Forecast{}, fmt.Errorf(
			"get city geocode from repository: %w",
			err,
		)
	}

	latitude := geocodeCity.Latitude
	longitude := geocodeCity.Longitude

	forecast, err := s.forecastRepository.GetForecast(
		ctx,
		latitude,
		longitude,
	)
	if err != nil {
		return domain.Forecast{}, fmt.Errorf(
			"get forecast from repository: %w",
			err,
		)
	}

	return domain.NewForecast(
		geocodeCity.Name,
		latitude,
		longitude,
		forecast.CurrentForecast,
		forecast.Days,
	), nil
}
