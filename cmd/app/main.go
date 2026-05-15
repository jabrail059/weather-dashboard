package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jabrail059/weather-dashboard/internal/handlers"
)

func main() {
	if err := RedisConnection(); err != nil {
		log.Fatal(err.Error())
	}

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/index.html")
	})
	r.Get("/weather", handlers.GeocodingAPI)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Println("Server is listening...")
	log.Fatal(http.ListenAndServe(":8081", r))
}
