package domain

import (
	"math"

	"github.com/jabrail059/weather-dashboard/internal/weather"
)

type CurrentForecast struct {
	Temperature  float64
	ApparentTemp float64
	WindSpeed    float64
	ImageSource  string
}

func NewCurrentForecast(
	temperature float64,
	apparentTemp float64,
	windSpeed float64,
	weatherCode int,
	isDay int,
) CurrentForecast {
	temperature = math.Round(temperature)
	apparentTemp = math.Round(apparentTemp)

	var imageSource string
	if weatherCode <= 2 {
		if isDay != 0 {
			imageSource = weather.ImageSources[weatherCode]
		} else {
			imageSource = weather.NightImageSources[weatherCode]
		}
	} else {
		imageSource = weather.ImageSources[weatherCode]
	}

	return CurrentForecast{
		Temperature:  temperature,
		ApparentTemp: apparentTemp,
		WindSpeed:    windSpeed,
		ImageSource:  imageSource,
	}
}
