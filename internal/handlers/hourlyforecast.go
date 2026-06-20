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
	sunrise := r.URL.Query().Get("sunrise")
	sunset := r.URL.Query().Get("sunset")

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

	hours, err := service.BuildHourlyWeather(oneDayForecast, sunrise, sunset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := models.HourlyData{
		City:    city,
		Date:    date,
		Sunrise: sunrise,
		Sunset:  sunset,
		Hours:   hours,
	}

	if err := view.RenderSeveralTemplates(w, "hourly.html", "hourlybase.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
