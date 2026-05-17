package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jabrail059/weather-dashboard/internal/handlers"
	"github.com/jabrail059/weather-dashboard/internal/service"
	"github.com/jabrail059/weather-dashboard/storage/sqlite"
)

func main() {
	if err := RedisConnection(); err != nil {
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
	service.SetStorage(cityStorage)

	r := chi.NewRouter()
	r.Get("/", handlers.MainPage)
	r.Get("/weather", handlers.GeocodingAPI)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Println("Server is listening...")
	log.Fatal(http.ListenAndServe(":8081", r))
}
