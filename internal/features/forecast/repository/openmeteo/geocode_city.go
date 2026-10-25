package forecast_openmeteo_repository

import (
	"context"
	"fmt"
	"strings"
	"unicode"

	"github.com/jabrail059/weather-dashboard/internal/core/domain"
	core_errors "github.com/jabrail059/weather-dashboard/internal/core/transport/http/errors"
	core_http_utils "github.com/jabrail059/weather-dashboard/internal/core/transport/http/utils"
)

const (
	defaultCount    = "1"
	defaultLanguage = "ru"
	defaultFormat   = "json"
)

func (r *ForecastOpenmeteoRepository) GeocodeCity(
	ctx context.Context,
	city string,
) (domain.GeocodeResult, error) {
	var geocodeCity GeocodeCityResponse
	normalizedCityName := normalizeCityName(city)

	params := map[string]string{
		"name":     normalizedCityName,
		"count":    defaultCount,
		"language": defaultLanguage,
		"format":   defaultFormat,
	}

	apiURL, err := core_http_utils.BuildURL(
		r.geocodingBaseURL,
		params,
	)
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
