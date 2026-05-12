package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jabrail059/weather-dashboard/internal/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/index.html")
	})
	r.Get("/weather", handlers.GeocodingAPIHandler)

	log.Println("Server is listening...")
	log.Fatal(http.ListenAndServe(":8081", r))
}
