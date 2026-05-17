package handlers

import (
	"html/template"
	"net"
	"net/http"

	"github.com/jabrail059/weather-dashboard/internal/service"
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
	report, err := service.GetCityByIP(ip)
	if err == nil && report.Status == "success" {
		data.AutoCity = report.City
		data.AutoDetected = true
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Ошибка шабонизации страницы", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Ошибка рендеринга страницы", http.StatusInternalServerError)
		return
	}
}
