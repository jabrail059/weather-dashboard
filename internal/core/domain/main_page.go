package domain

type CityReport struct {
	City     string
	Detected bool
}

func NewCityReport(
	city string,
	detected bool,
) CityReport {
	return CityReport{
		City:     city,
		Detected: detected,
	}
}

type History struct {
	Cities   []string
	Presence bool
}
