package service

import (
	"fmt"
	"math"
	"time"

	"github.com/jabrail059/weather-dashboard/internal/models"
	"github.com/jabrail059/weather-dashboard/internal/weather"
)

func BuildDayWeather(forecast *models.ReqDaily) ([]models.DayWeather, error) {
	days := make([]models.DayWeather, 0)
	for i := 0; i < len(forecast.Daily.Time); i++ {
		parsedTime, err := time.Parse("2006-01-02", forecast.Daily.Time[i])
		if err != nil {
			return nil, fmt.Errorf("Возникла ошибка при преобразовании даты")
		}
		formattedTime := parsedTime.Format("January 2, 2006")
		dailyTemp := models.DayWeather{
			Time:               formattedTime,
			TemperatureMax:     math.Round(forecast.Daily.TemperatureMax[i]),
			TemperatureMin:     math.Round(forecast.Daily.TemperatureMin[i]),
			WeatherDescription: weather.WeatherDescriptions[forecast.Daily.WeatherCode[i]],
			ImageSource:        weather.ImageSources[forecast.Daily.WeatherCode[i]],
		}
		days = append(days, dailyTemp)
	}
	return days, nil
}
