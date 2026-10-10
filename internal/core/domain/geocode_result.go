package domain

type GeocodeResult struct {
	ID        int
	Name      string
	Latitude  float64
	Longitude float64
}

func NewGeocodeResult(
	id int,
	name string,
	latitude float64,
	longitude float64,
) GeocodeResult {
	return GeocodeResult{
		ID:        id,
		Name:      name,
		Latitude:  latitude,
		Longitude: longitude,
	}
}
