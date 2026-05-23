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

type ReqDaily struct {
	Daily Daily `json:"daily"`
}

type Daily struct {
	Time           []string  `json:"time"`
	TemperatureMax []float64 `json:"temperature_2m_max"`
	TemperatureMin []float64 `json:"temperature_2m_min"`
	WeatherCode    []int     `json:"weather_code"`
}

type CityReport struct {
	Status    string  `json:"status"`
	City      string  `json:"city"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

type CitiesData struct {
	Cities   []string
	Presence bool
}

type DayWeather struct {
	Time               string
	TemperatureMax     float64
	TemperatureMin     float64
	WeatherDescription string
	ImageSource        string
}
