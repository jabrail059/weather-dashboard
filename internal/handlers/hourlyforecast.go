package handlers

import (
	"net/http"

	"github.com/jabrail059/weather-dashboard/internal/models"
	"github.com/jabrail059/weather-dashboard/internal/service"
	"github.com/jabrail059/weather-dashboard/internal/view"
)

func GetHourlyForecast(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	lat := r.URL.Query().Get("lat")
	lon := r.URL.Query().Get("lon")
	city := r.URL.Query().Get("city")

	forecast, err, statusCode := service.GetHourlyForecast(r.Context(), lat, lon, date)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}
	oneDayForecast, err := service.BuildOneDayForecast(forecast, date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hours, err := service.BuildHourlyWeather(oneDayForecast)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		City  string
		Date  string
		Hours []models.HourlyWeather
	}{
		City:  city,
		Date:  date,
		Hours: hours,
	}

	if err := view.RenderSeveralTemplates(w, "hourly.html", "hourlybase.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
