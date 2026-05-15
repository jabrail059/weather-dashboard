package handlers

import (
	"fmt"
	"html/template"
	"log/slog"
	"math"
	"net/http"
	"time"

	"github.com/jabrail059/weather-dashboard/internal/utils"
	"github.com/jabrail059/weather-dashboard/internal/weather"
)

func GeocodingAPI(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	geocode, err, statusCode := utils.GeocodeCity(city)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}
	forecast, err, statusCode := utils.GetForecast(geocode.Results[0].Latitude, geocode.Results[0].Longitude)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	Data := make([]struct {
		Time               string
		TemperatureMax     float64
		TemperatureMin     float64
		WeatherDescription string
		ImageSource        string
	}, 0)

	for i := 0; i < len(forecast.Daily.Time); i++ {
		parsedTime, err := time.Parse("2006-01-02", forecast.Daily.Time[i])
		if err != nil {
			http.Error(w, "Возникла ошибка при преобразовании даты", http.StatusInternalServerError)
			return
		}
		formattedTime := parsedTime.Format("January 2, 2006")
		dailyTemp := struct {
			Time               string
			TemperatureMax     float64
			TemperatureMin     float64
			WeatherDescription string
			ImageSource        string
		}{
			Time:               formattedTime,
			TemperatureMax:     math.Round(forecast.Daily.TemperatureMax[i]),
			TemperatureMin:     math.Round(forecast.Daily.TemperatureMin[i]),
			WeatherDescription: weather.WeatherDescriptions[forecast.Daily.WeatherCode[i]],
			ImageSource:        weather.ImageSources[forecast.Daily.WeatherCode[i]],
		}
		Data = append(Data, dailyTemp)
	}

	if r.Header.Get("HX-Request") != "true" {
		http.Error(w, "Не указан заголовок", http.StatusBadRequest)
		return
	}

	tmpl, err := template.ParseFiles("templates/base.html")
	if err != nil {
		http.Error(w, "Возникла ошибка парсинга во время шаболнизации страницы", http.StatusInternalServerError)
		return
	}
	if err := tmpl.ExecuteTemplate(w, "base.html", Data); err != nil {
		slog.Info(fmt.Sprintf("Ошибка шаблонизации страницы %v", err.Error()))
		http.Error(w, "Возникла ошибка шаблонизации страницы", http.StatusInternalServerError)
		return
	}
}
