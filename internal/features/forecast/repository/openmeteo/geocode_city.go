package forecast_openmeteo_repository

import (
	"context"
	"fmt"
	"strings"
	"unicode"

	"github.com/jabrail059/weather-dashboard/internal/core/domain"
	core_errors "github.com/jabrail059/weather-dashboard/internal/core/transport/http/errors"
)

func (r *ForecastRepository) GeocodeCity(
	ctx context.Context,
	city string,
) (domain.GeocodeResult, error) {
	var geocodeCity GeocodeCityResponse
	normalizedCityName := normalizeCityName(city)

	// result, err := cityStorage.Select(ctx, cityName)
	// if err == nil {
	// 	geo = models.GeoRequest{
	// 		Results: []models.Result{*result},
	// 	}
	// 	return &geo, http.StatusOK, nil
	// }
	// if err != storage.ErrCityNotSaved {
	// 	return nil, http.StatusInternalServerError, fmt.Errorf("не удалось получить данные")
	// }

	apiURL, err := r.buildGeocodingURL(normalizedCityName)
	if err != nil {
		return domain.GeocodeResult{}, fmt.Errorf("build geocoding url: %w", err)
	}

	err = r.getResponseFromOpenMeteo(
		ctx,
		apiURL,
		&geocodeCity,
	)
	if err != nil {
		return domain.GeocodeResult{}, fmt.Errorf("get response from openmeteo: %w", err)
	}

	if len(geocodeCity.Results) == 0 {
		return domain.GeocodeResult{}, fmt.Errorf(
			"forecast for 'city'='%s' not found: %w",
			normalizedCityName,
			core_errors.ErrNotFound,
		)
	}

	geocodeCity.Results[0].Name = normalizedCityName

	geocodeResult := geocodeResultsFromResponse(geocodeCity)
	// err = cityStorage.Save(ctx, &geo.Results[0])
	// if err != nil {
	// 	return nil, http.StatusInternalServerError, fmt.Errorf("не удалось сохранить данные")
	// }

	return geocodeResult, nil
}

func normalizeCityName(city string) string {
	city = strings.TrimSpace(city)
	city = strings.ToLower(city)

	if city == "" {
		return city
	}

	runes := []rune(city)

	makeUpper := true
	for i, r := range runes {
		if makeUpper && unicode.IsLetter(r) {
			runes[i] = unicode.ToUpper(r)
			makeUpper = false
			continue
		}

		if r == ' ' || r == '-' {
			makeUpper = true
		}
	}

	return string(runes)
}
