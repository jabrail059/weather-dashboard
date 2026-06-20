package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jabrail059/weather-dashboard/internal/handlers"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", handlers.MainPage)
	r.Get("/weather", handlers.Weather)
	r.Post("/favorites/add", handlers.AddInFavorites)
	r.Get("/favorites", handlers.GetFavorites)
	r.Delete("/favorites/delete", handlers.DeleteFromFavorites)
	r.Get("/searchhistory", handlers.SearchHistory)
	r.Get("/weather/hourly", handlers.GetHourlyForecast)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	return r
}
