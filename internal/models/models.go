package models

type Result struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GeoRequest struct {
	Results []Result `json:"results"`
}

type ReqHourly struct {
	Hourly Hourly `json:"hourly"`
}

type Hourly struct {
	Time        []string  `json:"time"`
	Temperature []float64 `json:"temperature_2m"`
}
