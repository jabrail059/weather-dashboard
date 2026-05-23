package handlers

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/jabrail059/weather-dashboard/internal/service"
	"github.com/jabrail059/weather-dashboard/internal/session"
	"github.com/jabrail059/weather-dashboard/internal/view"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		ip = r.RemoteAddr
	}
	data := struct {
		AutoCity     string
		AutoDetected bool
	}{}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*15)
	defer cancel()

	report, err := service.GetCityByIP(ctx, ip)
	if err == nil && report.Status == "success" {
		data.AutoCity = report.City
		data.AutoDetected = true
	}

	cookie := session.GetOrCreate(r)
	http.SetCookie(w, cookie)

	if err := view.RenderTemplate(w, "index.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
