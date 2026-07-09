package service

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/jabrail059/weather-dashboard/internal/models"
	"github.com/jabrail059/weather-dashboard/internal/weather"
)

func BuildDayWeather(forecast *models.Daily) ([]models.DayWeather, error) {
	days := make([]models.DayWeather, 0)
	for i := 0; i < len(forecast.Time); i++ {
		parsedTime, err := time.Parse("2006-01-02", forecast.Time[i])
		if err != nil {
			return nil, fmt.Errorf("Возникла ошибка при преобразовании даты: %v", err)
		}
		sunrise, err := time.Parse("2006-01-02T15:04", forecast.Sunrise[i])
		if err != nil {
			return nil, fmt.Errorf("Возникла ошибка при преобразовании даты: %v", err)
		}
		sunset, err := time.Parse("2006-01-02T15:04", forecast.Sunset[i])
		if err != nil {
			return nil, fmt.Errorf("Возникла ошибка при преобразовании даты: %v", err)
		}
		dailyTemp := models.DayWeather{
			Date:               parsedTime.Format("2006-01-02"),
			Time:               parsedTime.Format("January 2, 2006"),
			TemperatureMax:     math.Round(forecast.TemperatureMax[i]),
			TemperatureMin:     math.Round(forecast.TemperatureMin[i]),
			WeatherDescription: weather.WeatherDescriptions[forecast.WeatherCode[i]],
			ImageSource:        weather.ImageSources[forecast.WeatherCode[i]],
			Sunrise:            sunrise.Format("15:04"),
			Sunset:             sunset.Format("15:04"),
		}
		days = append(days, dailyTemp)
	}
	return days, nil
}

func BuildHourlyWeather(forecast *models.Hourly, sunrise string, sunset string) ([]models.HourlyWeather, error) {
	hours := make([]models.HourlyWeather, 0)
	for i := 0; i < len(forecast.Time); i++ {
		parsedTime, err := time.Parse("2006-01-02T15:04", forecast.Time[i])
		if err != nil {
			return nil, fmt.Errorf("Возникла ошибка при преобразовании даты: %v", err)
		}
		formattedTime := parsedTime.Format("15:04")
		hourlyTemp := models.HourlyWeather{
			Time:        formattedTime,
			Temperature: math.Round(forecast.Temperature[i]),
		}

		if forecast.WeatherCode[i] <= 2 {
			if formattedTime <= sunrise || formattedTime >= sunset {
				hourlyTemp.ImageSource = weather.NightImageSources[forecast.WeatherCode[i]]
			} else {
				hourlyTemp.ImageSource = weather.ImageSources[forecast.WeatherCode[i]]
			}
		} else {
			hourlyTemp.ImageSource = weather.ImageSources[forecast.WeatherCode[i]]
		}
		hours = append(hours, hourlyTemp)
	}
	return hours, nil
}

func BuildOneDayForecast(forecast *models.Hourly, date string) (*models.Hourly, error) {
	newForecast := &models.Hourly{
		Time:        make([]string, 0),
		Temperature: make([]float64, 0),
		WeatherCode: make([]int, 0),
	}

	for i := 0; i < len(forecast.Time); i++ {
		if strings.HasPrefix(forecast.Time[i], date) {
			newForecast.Time = append(newForecast.Time, forecast.Time[i])
			newForecast.Temperature = append(newForecast.Temperature, forecast.Temperature[i])
			newForecast.WeatherCode = append(newForecast.WeatherCode, forecast.WeatherCode[i])
		}
	}

	if len(newForecast.Time) == 0 {
		return nil, fmt.Errorf("Прогноз на %s не найден", date)
	}

	return newForecast, nil
}

func BuildCurrentWeather(forecast *models.Current) models.CurrentWeather {
	CurrentWeather := models.CurrentWeather{
		Temperature:  math.Round(forecast.Temperature),
		ApparentTemp: math.Round(forecast.ApparentTemp),
		WindSpeed:    forecast.WindSpeed,
	}

	if forecast.WeatherCode <= 2 {
		if forecast.IsDay != 0 {
			CurrentWeather.ImageSource = weather.ImageSources[forecast.WeatherCode]
		} else {
			CurrentWeather.ImageSource = weather.NightImageSources[forecast.WeatherCode]
		}
	} else {
		CurrentWeather.ImageSource = weather.ImageSources[forecast.WeatherCode]
	}

	return CurrentWeather
}
