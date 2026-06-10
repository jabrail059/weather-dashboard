package handlers

import (
	"context"
	"math"
	"net/http"
	"time"

	"github.com/jabrail059/weather-dashboard/internal/app"
	"github.com/jabrail059/weather-dashboard/internal/models"
	"github.com/jabrail059/weather-dashboard/internal/service"
	"github.com/jabrail059/weather-dashboard/internal/session"
	"github.com/jabrail059/weather-dashboard/internal/view"
	"github.com/jabrail059/weather-dashboard/internal/weather"
)

func Weather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")

	if r.Header.Get("HX-Request") != "true" {
		http.Error(w, "Не указан заголовок", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*15)
	defer cancel()

	cookie := session.GetOrCreate(r)
	http.SetCookie(w, cookie)

	geocode, err, statusCode := service.GeocodeCity(ctx, app.Storage(), city)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}
	forecast, err, statusCode := service.GetDailyForecast(ctx, geocode.Results[0].Latitude, geocode.Results[0].Longitude)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}
	currentForecast, err, statusCode := service.GetCurrentForecast(ctx, geocode.Results[0].Latitude, geocode.Results[0].Longitude)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	data := struct {
		City           string
		Latitude       float64
		Longitude      float64
		CurrentWeather models.CurrentWeather
		Days           []models.DayWeather
	}{
		City:      city,
		Latitude:  geocode.Results[0].Latitude,
		Longitude: geocode.Results[0].Longitude,
	}

	data.CurrentWeather = models.CurrentWeather{
		Temperature:  math.Round(currentForecast.Temperature),
		ApparentTemp: math.Round(currentForecast.ApparentTemp),
		WindSpeed:    currentForecast.WindSpeed,
		ImageSource:  weather.ImageSources[currentForecast.WeatherCode],
	}

	data.Days, err = service.BuildDayWeather(forecast)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := service.AddSearchHistory(r.Context(), city, cookie.Value); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Trigger", "searchHistoryChanged")

	if err := view.RenderPartOfTemplate(w, "base.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
