package handlers

import (
	"net/http"

	"github.com/jabrail059/weather-dashboard/internal/models"
	"github.com/jabrail059/weather-dashboard/internal/service"
	"github.com/jabrail059/weather-dashboard/internal/session"
	"github.com/jabrail059/weather-dashboard/internal/view"
)

func GetFavorites(w http.ResponseWriter, r *http.Request) {
	cookie := session.GetOrCreate(r)
	http.SetCookie(w, cookie)

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
