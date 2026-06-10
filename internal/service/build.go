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
			return nil, fmt.Errorf("Возникла ошибка при преобразовании даты")
		}
		sunrise, err := time.Parse("2006-01-02T15:04", forecast.Sunrise[i])
		if err != nil {
			return nil, fmt.Errorf("Возникла ошибка при преобразовании даты")
		}
		sunset, err := time.Parse("2006-01-02T15:04", forecast.Sunset[i])
		if err != nil {
			return nil, fmt.Errorf("Возникла ошибка при преобразовании даты")
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

func BuildHourlyWeather(forecast *models.Hourly) ([]models.HourlyWeather, error) {
	hours := make([]models.HourlyWeather, 0)
	for i := 0; i < len(forecast.Time); i++ {
		parsedTime, err := time.Parse("2006-01-02T15:04", forecast.Time[i])
		if err != nil {
			return nil, fmt.Errorf("Возникла ошибка при преобразовании даты")
		}
		hourlyTemp := models.HourlyWeather{
			Time:        parsedTime.Format("15:04"),
			Temperature: math.Round(forecast.Temperature[i]),
			ImageSource: weather.ImageSources[forecast.WeatherCode[i]],
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
	return newForecast, nil
}
