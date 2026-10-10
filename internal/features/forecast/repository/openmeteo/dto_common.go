package forecast_openmeteo_repository

import (
	"fmt"

	"github.com/jabrail059/weather-dashboard/internal/core/domain"
)

type OpenMeteoResponse struct {
	Daily   DailyResponse   `json:"daily"`
	Current CurrentResponse `json:"current"`
}

type DailyResponse struct {
	Time           []string  `json:"time"`
	TemperatureMax []float64 `json:"temperature_2m_max"`
	TemperatureMin []float64 `json:"temperature_2m_min"`
	WeatherCode    []int     `json:"weather_code"`
	Sunrise        []string  `json:"sunrise"`
	Sunset         []string  `json:"sunset"`
}

type HourlyResponse struct {
	Time        []string  `json:"time"`
	Temperature []float64 `json:"temperature_2m"`
	WeatherCode []int     `json:"weather_code"`
}

func hourlyForecastFromResponse(
	hourlyResponse HourlyResponse,
) domain.OneDayForecast {
	return domain.NewOneDayForecast(
		hourlyResponse.Time,
		hourlyResponse.Temperature,
		hourlyResponse.WeatherCode,
	)
}

type CurrentResponse struct {
	Temperature  float64 `json:"temperature_2m"`
	ApparentTemp float64 `json:"apparent_temperature"`
	WindSpeed    float64 `json:"wind_speed_10m"`
	WeatherCode  int     `json:"weather_code"`
	IsDay        int     `json:"is_day"`
}

func forecastFromResponse(
	openMeteoResponse OpenMeteoResponse,
) (domain.Forecast, error) {
	if err := validateDailyResponse(openMeteoResponse.Daily); err != nil {
		return domain.Forecast{}, fmt.Errorf(
			"validate daily response: %w",
			err,
		)
	}

	length := len(openMeteoResponse.Daily.Time)
	daysForecasts := make([]domain.DayForecast, length)

	for i := 0; i < length; i++ {
		oneDayForecast, err := domain.NewDayForecast(
			openMeteoResponse.Daily.Time[i],
			openMeteoResponse.Daily.TemperatureMax[i],
			openMeteoResponse.Daily.TemperatureMin[i],
			openMeteoResponse.Daily.WeatherCode[i],
			openMeteoResponse.Daily.Sunrise[i],
			openMeteoResponse.Daily.Sunset[i],
		)
		if err != nil {
			return domain.Forecast{}, fmt.Errorf("build one day forecast: %w", err)
		}

		daysForecasts[i] = oneDayForecast
	}

	currentForecast := currentForecastFromResponse(
		openMeteoResponse.Current,
	)

	return domain.Forecast{
		CurrentForecast: currentForecast,
		Days:            daysForecasts,
	}, nil
}

func currentForecastFromResponse(
	currentResponse CurrentResponse,
) domain.CurrentForecast {
	return domain.NewCurrentForecast(
		currentResponse.Temperature,
		currentResponse.ApparentTemp,
		currentResponse.WindSpeed,
		currentResponse.WeatherCode,
		currentResponse.IsDay,
	)
}

func validateDailyResponse(
	response DailyResponse,
) error {
	expectedLength := len(response.Time)

	if expectedLength == 0 {
		return fmt.Errorf("daily forecast is empty")
	}

	fields := []struct {
		name   string
		length int
	}{
		{"temperature_max", len(response.TemperatureMax)},
		{"temperature_min", len(response.TemperatureMin)},
		{"weather_code", len(response.WeatherCode)},
		{"sunrise", len(response.Sunrise)},
		{"sunset", len(response.Sunset)},
	}

	for _, field := range fields {
		if field.length != expectedLength {
			return fmt.Errorf(
				"daily response length mismatch: time=%d, %s=%d",
				expectedLength,
				field.name,
				field.length,
			)
		}
	}

	return nil
}

type GeocodeCityResponse struct {
	Results []ResultResponse `json:"results"`
}

type ResultResponse struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func geocodeResultsFromResponse(
	geocodeCityResponse GeocodeCityResponse,
) domain.GeocodeResult {
	id := geocodeCityResponse.Results[0].ID
	name := geocodeCityResponse.Results[0].Name
	latitude := geocodeCityResponse.Results[0].Latitude
	longitude := geocodeCityResponse.Results[0].Longitude

	return domain.NewGeocodeResult(
		id,
		name,
		latitude,
		longitude,
	)
}
