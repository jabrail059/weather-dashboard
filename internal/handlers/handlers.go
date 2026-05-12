package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"

	"github.com/jabrail059/weather-dashboard/internal/models"
)

func GeocodingAPIHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("city")
	firstResp, err := http.Get(fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=ru&format=json", name))
	if err != nil {
		http.Error(w, "Не удалось получить данные о погоде", http.StatusBadRequest)
		return
	}
	defer firstResp.Body.Close()

	if firstResp.StatusCode != 200 {
		http.Error(w, "Возникла ошибка при получении данных", http.StatusBadRequest)
		return
	}

	var geo models.GeoRequest
	if err = json.NewDecoder(firstResp.Body).Decode(&geo); err != nil {
		http.Error(w, "Не удалось получить данные", http.StatusBadRequest)
		return
	}
	if len(geo.Results) == 0 {
		http.Error(w, "Город не найден", http.StatusNotFound)
		return
	}

	secondResp, err := http.Get(fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%v&longitude=%v&hourly=temperature_2m", geo.Results[0].Latitude, geo.Results[0].Longitude))
	if err != nil {
		http.Error(w, "Не удалось получить данные", http.StatusBadRequest)
		return
	}
	defer secondResp.Body.Close()

	if secondResp.StatusCode != 200 {
		http.Error(w, "Возникла ошибка при получении данных", http.StatusBadRequest)
		return
	}

	var res models.ReqHourly
	if err = json.NewDecoder(secondResp.Body).Decode(&res); err != nil {
		http.Error(w, "Не удалось получить данные", http.StatusBadRequest)
		return
	}

	Data := make([]struct {
		Time        string
		Temperature float64
	}, 0)

	for i := 0; i < len(res.Hourly.Temperature); i++ {
		dailyTemp := struct {
			Time        string
			Temperature float64
		}{
			Time:        res.Hourly.Time[i],
			Temperature: res.Hourly.Temperature[i],
		}
		Data = append(Data, dailyTemp)
	}

	if r.Header.Get("HX-Request") != "true" {
		http.Error(w, "Не указан заголовок", http.StatusBadRequest)
	}

	tmpl, err := template.ParseFiles("templates/base.html")
	if err != nil {
		http.Error(w, "Возникла ошибка парсинга во время шаболнизации страницы", http.StatusInternalServerError)
		return
	}
	if err := tmpl.ExecuteTemplate(w, "base.html", Data); err != nil {
		slog.Info(fmt.Sprintf("Ошибка шаблонизации страницы %v", err.Error()))
		http.Error(w, "Возникла ошибка шаблонизации страницы", http.StatusInternalServerError)
		return
	}
}
