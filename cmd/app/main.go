package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jabrail059/weather-dashboard/internal/app"
	"github.com/jabrail059/weather-dashboard/internal/handlers"
	"github.com/jabrail059/weather-dashboard/storage/redis"
	"github.com/jabrail059/weather-dashboard/storage/sqlite"
)

func main() {
	if err := redis.Connection(); err != nil {
		log.Fatal(err.Error())
	}

	dbPath := os.Getenv("SQLITE_PATH")
	if dbPath == "" {
		dbPath = "./data/weather.db"
	}
	cityStorage, err := sqlite.New(dbPath)
	if err != nil {
		log.Fatal("Не удалось подключиться к sqlite")
	}
	app.SetStorage(cityStorage)

	r := chi.NewRouter()
	r.Get("/", handlers.MainPage)
	r.Get("/weather", handlers.Weather)
	r.Post("/favorites/add", handlers.AddInFavorites)
	r.Get("/favorites", handlers.GetFavorites)
	r.Delete("/favorites/delete", handlers.DeleteFromFavorites)
	r.Get("/searchhistory", handlers.SearchHistory)
	r.Get("/weather/hourly", handlers.GetHourlyForecast)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Println("Server is listening...")
	log.Fatal(http.ListenAndServe(":8081", r))
}
