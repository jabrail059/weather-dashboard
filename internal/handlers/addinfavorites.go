package handlers

import (
	"net/http"

	"github.com/jabrail059/weather-dashboard/internal/models"
	"github.com/jabrail059/weather-dashboard/internal/service"
	"github.com/jabrail059/weather-dashboard/internal/session"
	"github.com/jabrail059/weather-dashboard/internal/view"
)

func AddInFavorites(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	cookie := session.GetOrCreate(r)
	http.SetCookie(w, cookie)

	if err := service.SaveInFavorites(r.Context(), cookie.Value, city); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cities, err := service.GetCities(r.Context(), cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := models.CitiesData{
		Cities:   cities,
		Presence: len(cities) > 0,
	}

	if err := view.RenderPartOfTemplate(w, "favorites.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
