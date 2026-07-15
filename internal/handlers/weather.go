package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/jabrail059/weather-dashboard/internal/app"
	"github.com/jabrail059/weather-dashboard/internal/models"
	"github.com/jabrail059/weather-dashboard/internal/service"
	"github.com/jabrail059/weather-dashboard/internal/session"
	"github.com/jabrail059/weather-dashboard/internal/view"
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

	geocode, statusCode, err := service.GeocodeCity(ctx, app.Storage(), city)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}
	forecast, statusCode, err := service.GetDailyForecast(ctx, geocode.Results[0].Latitude, geocode.Results[0].Longitude)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}
	currentForecast, statusCode, err := service.GetCurrentForecast(ctx, geocode.Results[0].Latitude, geocode.Results[0].Longitude)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	data := models.DailyData{
		City:           geocode.Results[0].Name,
		Latitude:       geocode.Results[0].Latitude,
		Longitude:      geocode.Results[0].Longitude,
		CurrentWeather: service.BuildCurrentWeather(currentForecast),
	}

	data.Days, err = service.BuildDayWeather(forecast)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := service.AddSearchHistory(r.Context(), data.City, cookie.Value); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Trigger", "searchHistoryChanged")

	if err := view.RenderPartOfTemplate(w, "base.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
