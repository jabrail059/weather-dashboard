package forecast_service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jabrail059/weather-dashboard/internal/core/domain"
	core_errors "github.com/jabrail059/weather-dashboard/internal/core/transport/http/errors"
)

func (s *ForecastService) GetHourlyForecast(
	ctx context.Context,
	date *time.Time,
	city *string,
	sunrise *time.Time,
	sunset *time.Time,
) (domain.HourlyRender, error) {
	err := validateHourlyQueryParams(
		date,
		city,
		sunrise,
		sunset,
	)
	if err != nil {
		return domain.HourlyRender{}, fmt.Errorf(
			"validate hourly query params: %w",
			err,
		)
	}

	strDate := date.Format("2006-01-02")
	strSunrise := sunrise.Format("15:04")
	strSunset := sunset.Format("15:04")

	geocodeCity, err := s.forecastPostgresRepository.GetCity(ctx, *city)
	if err != nil {
		if !errors.Is(err, core_errors.ErrNotFound) {
			return domain.HourlyRender{}, fmt.Errorf("get city from postgres: %w", err)
		}
		geocodeCity, err = s.forecastOpenmeteoRepository.GeocodeCity(
			ctx,
			*city,
		)
		if err != nil {
			return domain.HourlyRender{}, fmt.Errorf(
				"get city geocode from repository: %w",
				err,
			)
		}
	}

	latitude := geocodeCity.Latitude
	longitude := geocodeCity.Longitude

	forecast, err := s.forecastOpenmeteoRepository.GetHourlyForecast(
		ctx,
		latitude,
		longitude,
	)
	if err != nil {
		return domain.HourlyRender{}, fmt.Errorf(
			"get one day of hourly forecast from repository: %w",
			err,
		)
	}

	hourlyForecast, err := domain.NewHourlyForecast(
		forecast,
		strDate,
		strSunrise,
		strSunset,
	)
	if err != nil {
		return domain.HourlyRender{}, fmt.Errorf(
			"get hourly forecast: %w",
			err,
		)
	}

	hourlyRender := domain.HourlyRender{
		City:    *city,
		Date:    strDate,
		Sunrise: strSunrise,
		Sunset:  strSunset,
		Hours:   hourlyForecast,
	}

	return hourlyRender, nil
}

func validateHourlyQueryParams(
	date *time.Time,
	city *string,
	sunrise *time.Time,
	sunset *time.Time,
) error {
	if date == nil {
		return fmt.Errorf(
			"date must be not NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}
	if city == nil {
		return fmt.Errorf(
			"city must be not NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}
	if sunrise == nil {
		return fmt.Errorf(
			"sunrise must be not NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}
	if sunset == nil {
		return fmt.Errorf(
			"sunset must be not NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}
