package domain

import (
	"fmt"
	"math"
	"strings"
	"time"

	core_errors "github.com/jabrail059/weather-dashboard/internal/core/transport/http/errors"
)

type HourlyRender struct {
	City    string
	Date    string
	Sunrise string
	Sunset  string
	Hours   []HourlyForecast
}

type OneDayForecast struct {
	Time        []string
	Temperature []float64
	WeatherCode []int
}

func NewOneDayForecast(
	time []string,
	temperature []float64,
	weatherCode []int,
) OneDayForecast {
	return OneDayForecast{
		Time:        time,
		Temperature: temperature,
		WeatherCode: weatherCode,
	}
}

type HourlyForecast struct {
	Time        string
	Temperature float64
	ImageSource string
}

func NewHourlyForecast(
	oneDayForecast OneDayForecast,
	date string,
	sunrise string,
	sunset string,
) ([]HourlyForecast, error) {
	forecast := &OneDayForecast{
		Time:        make([]string, 0),
		Temperature: make([]float64, 0),
		WeatherCode: make([]int, 0),
	}

	for i := 0; i < len(oneDayForecast.Time); i++ {
		if strings.HasPrefix(oneDayForecast.Time[i], date) {
			forecast.Time = append(forecast.Time, oneDayForecast.Time[i])
			forecast.Temperature = append(forecast.Temperature, oneDayForecast.Temperature[i])
			forecast.WeatherCode = append(forecast.WeatherCode, oneDayForecast.WeatherCode[i])
		}
	}

	if len(forecast.Time) == 0 {
		return nil, fmt.Errorf(
			"date: %s forecast not found: %w",
			date,
			core_errors.ErrNotFound,
		)
	}

	hours := make([]HourlyForecast, len(forecast.Time))

	for i := 0; i < len(forecast.Time); i++ {
		hour, err := validateHourlyForecast(
			forecast.Time[i],
			forecast.Temperature[i],
			forecast.WeatherCode[i],
			sunrise,
			sunset,
		)
		if err != nil {
			return nil, fmt.Errorf("validate hourly forecast: %w", err)
		}
		hours[i] = hour
	}
	return hours, nil

}

func validateHourlyForecast(
	date string,
	temperature float64,
	weatherCode int,
	sunrise string,
	sunset string,
) (HourlyForecast, error) {
	parsedTime, err := time.Parse("2006-01-02T15:04", date)
	if err != nil {
		return HourlyForecast{}, fmt.Errorf("time parse: %w", err)
	}
	formattedTime := parsedTime.Format("15:04")
	hour := HourlyForecast{
		Time:        formattedTime,
		Temperature: math.Round(temperature),
	}

	if weatherCode <= 2 {
		if formattedTime <= sunrise || formattedTime >= sunset {
			hour.ImageSource = NightImageSources[weatherCode]
		} else {
			hour.ImageSource = ImageSources[weatherCode]
		}
	} else {
		hour.ImageSource = ImageSources[weatherCode]
	}

	return hour, nil
}
