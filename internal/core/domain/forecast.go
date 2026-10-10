package domain

import (
	"fmt"
	"math"
	"time"

	"github.com/jabrail059/weather-dashboard/internal/weather"
)

type Forecast struct {
	City            string
	Latitude        float64
	Longitude       float64
	CurrentForecast CurrentForecast
	Days            []DayForecast
}

func NewForecast(
	city string,
	latitude float64,
	longitude float64,
	currentForecast CurrentForecast,
	days []DayForecast,
) Forecast {
	return Forecast{
		City:            city,
		Latitude:        latitude,
		Longitude:       longitude,
		CurrentForecast: currentForecast,
		Days:            days,
	}
}

type DayForecast struct {
	Date               string
	Time               string
	TemperatureMax     float64
	TemperatureMin     float64
	WeatherDescription string
	ImageSource        string
	Sunrise            string
	Sunset             string
}

func NewDayForecast(
	time string,
	temperatureMax float64,
	temperatureMin float64,
	weatherCode int,
	sunrise string,
	sunset string,
) (DayForecast, error) {
	parsedTime, parsedSunrise, parsedSunset, err := parseTimeSunriseSunset(
		time,
		sunrise,
		sunset,
	)
	if err != nil {
		return DayForecast{}, err
	}

	date := parsedTime.Format("2006-01-02")
	time = parsedTime.Format("January 2, 2006")
	temperatureMax = math.Round(temperatureMax)
	temperatureMin = math.Round(temperatureMin)
	weatherDescription := weather.WeatherDescriptions[weatherCode]
	imageSource := weather.ImageSources[weatherCode]
	sunrise = parsedSunrise.Format("15:04")
	sunset = parsedSunset.Format("15:04")

	return DayForecast{
		Date:               date,
		Time:               time,
		TemperatureMax:     temperatureMax,
		TemperatureMin:     temperatureMin,
		WeatherDescription: weatherDescription,
		ImageSource:        imageSource,
		Sunrise:            sunrise,
		Sunset:             sunset,
	}, nil
}

func parseTimeSunriseSunset(
	forecastTime string,
	sunrise string,
	sunset string,
) (*time.Time, *time.Time, *time.Time, error) {
	parsedTime, err := time.Parse("2006-01-02", forecastTime)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("date parse: %w", err)
	}

	parsedSunrise, err := time.Parse("2006-01-02T15:04", sunrise)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("sunrise parse: %w", err)
	}

	parsedSunset, err := time.Parse("2006-01-02T15:04", sunset)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("sunset parse: %w", err)
	}

	return &parsedTime, &parsedSunrise, &parsedSunset, nil
}
