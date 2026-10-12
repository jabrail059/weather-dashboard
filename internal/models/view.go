package models

type CitiesData struct {
	Cities   []string
	Presence bool
}

type HourlyWeather struct {
	Time        string
	Temperature float64
	ImageSource string
}

type HourlyData struct {
}
