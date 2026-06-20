package models

type OpenMeteoResponse struct {
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

type Current struct {
	Temperature  float64 `json:"temperature_2m"`
	ApparentTemp float64 `json:"apparent_temperature"`
	WindSpeed    float64 `json:"wind_speed_10m"`
	WeatherCode  int     `json:"weather_code"`
	IsDay        int     `json:"is_day"`
}
