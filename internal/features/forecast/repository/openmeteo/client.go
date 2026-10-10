package forecast_openmeteo_repository

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const (
	defaultTimezone = "auto"
	defaultCount    = "1"
	defaultLanguage = "ru"
	defaultFormat   = "json"
)

func (r *ForecastRepository) getResponseFromOpenMeteo(
	ctx context.Context,
	apiURL string,
	response any,
) error {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return fmt.Errorf(
			"create new request: %w",
			err,
		)
	}

	clientDo, err := r.client.Do(request)
	if err != nil {
		return fmt.Errorf(
			"get forecast from openmeteo: %w",
			err,
		)
	}
	defer clientDo.Body.Close() //nolint:errcheck

	if clientDo.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"clientDo status code: %d",
			clientDo.StatusCode,
		)
	}

	if err = json.NewDecoder(clientDo.Body).Decode(response); err != nil {
		return fmt.Errorf(
			"decode clientDo: %w",
			err,
		)
	}

	return nil
}

func (r *ForecastRepository) buildForecastURL(
	latitude float64,
	longitude float64,
	forecastParams map[string]string,
) (string, error) {
	u, err := url.Parse(r.forecastBaseURL)
	if err != nil {
		return "", fmt.Errorf("parse Open Meteo URL: %w", err)
	}

	query := u.Query()
	query.Set("latitude", strconv.FormatFloat(latitude, 'f', -1, 64))
	query.Set("longitude", strconv.FormatFloat(longitude, 'f', -1, 64))
	query.Set("timezone", defaultTimezone)

	for key, value := range forecastParams {
		query.Set(key, value)
	}

	u.RawQuery = query.Encode()
	return u.String(), nil
}

func (r *ForecastRepository) buildGeocodingURL(
	name string,
) (string, error) {
	u, err := url.Parse(r.geocodingBaseURL)
	if err != nil {
		return "", fmt.Errorf("parse Open Meteo URL: %w", err)
	}

	query := u.Query()
	query.Set("name", name)
	query.Set("count", defaultCount)
	query.Set("language", defaultLanguage)
	query.Set("format", defaultFormat)

	u.RawQuery = query.Encode()
	return u.String(), nil
}
