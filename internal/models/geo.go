package models

type GeoRequest struct {
	Results []Result `json:"results"`
}

type Result struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type CityReport struct {
	Status    string  `json:"status"`
	City      string  `json:"city"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}
