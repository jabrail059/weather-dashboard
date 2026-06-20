package models

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

type CurrentWeather struct {
	Temperature  float64
	ApparentTemp float64
	WindSpeed    float64
	ImageSource  string
}

type CitiesData struct {
	Cities   []string
	Presence bool
}

type HourlyData struct {
	City    string
	Date    string
	Sunrise string
	Sunset  string
	Hours   []HourlyWeather
}

type DailyData struct {
	City           string
	Latitude       float64
	Longitude      float64
	CurrentWeather CurrentWeather
	Days           []DayWeather
}
