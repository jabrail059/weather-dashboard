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

type HistoryList struct {
	Cities   []string
	Presence bool
	CacheErr error
}
