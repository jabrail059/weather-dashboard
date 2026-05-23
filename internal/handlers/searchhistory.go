package handlers

import (
	"net/http"

	"github.com/jabrail059/weather-dashboard/internal/service"
	"github.com/jabrail059/weather-dashboard/internal/session"
	"github.com/jabrail059/weather-dashboard/internal/view"
)

func SearchHistory(w http.ResponseWriter, r *http.Request) {
	cookie := session.GetOrCreate(r)
	http.SetCookie(w, cookie)

	historyData, err := service.GetSearchHistory(r.Context(), cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := view.RenderPartOfTemplate(w, "searchhistory.html", historyData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
