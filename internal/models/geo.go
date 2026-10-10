package models

type CityReport struct {
	Status    string  `json:"status"`
	City      string  `json:"city"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}
