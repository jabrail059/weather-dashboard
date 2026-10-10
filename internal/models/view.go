package models

type HourlyWeather struct {
	Time        string
	Temperature float64
	ImageSource string
}

type CitiesData struct {
	Cities   []string
	Presence bool
}

type HourlyData struct {
}
