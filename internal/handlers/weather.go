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

	geocode, err, statusCode := service.GeocodeCity(ctx, app.Storage(), city)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}
	forecast, err, statusCode := service.GetForecast(ctx, geocode.Results[0].Latitude, geocode.Results[0].Longitude)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}
	data := struct {
		City string
		Days []models.DayWeather
	}{City: city}

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
