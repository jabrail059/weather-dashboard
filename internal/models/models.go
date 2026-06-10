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
	Daily   Daily   `json:"daily"`
	Hourly  Hourly  `json:"hourly"`
	Current Current `json:"current"`
}

type Hourly struct {
	Time        []string  `json:"time"`
	Temperature []float64 `json:"temperature_2m"`
	WeatherCode []int     `json:"weather_code"`
}

type Daily struct {
	Time           []string  `json:"time"`
	TemperatureMax []float64 `json:"temperature_2m_max"`
	TemperatureMin []float64 `json:"temperature_2m_min"`
	WeatherCode    []int     `json:"weather_code"`
	Sunrise        []string  `json:"sunrise"`
	Sunset         []string  `json:"sunset"`
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
	Date               string
	Time               string
	TemperatureMax     float64
	TemperatureMin     float64
	WeatherDescription string
	ImageSource        string
	Sunrise            string
	Sunset             string
}

type HourlyWeather struct {
	Time        string
	Temperature float64
	ImageSource string
}

type Current struct {
	Temperature  float64 `json:"temperature_2m"`
	ApparentTemp float64 `json:"apparent_temperature"`
	WindSpeed    float64 `json:"wind_speed_10m"`
	WeatherCode  int     `json:"weather_code"`
}

type CurrentWeather struct {
	Temperature  float64
	ApparentTemp float64
	WindSpeed    float64
	ImageSource  string
}
