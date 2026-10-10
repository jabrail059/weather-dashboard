package core_http_request

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	core_errors "github.com/jabrail059/weather-dashboard/internal/core/transport/http/errors"
)

func GetStringQueryParam(r *http.Request, key string) *string {
	val := r.URL.Query().Get(key)
	if val == "" {
		return nil
	}

	return &val
}

func GetFloatQueryParam(r *http.Request, key string) (*float64, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, fmt.Errorf(
			"no key=%s in query params: %w",
			key,
			core_errors.ErrInvalidArgument,
		)
	}

	val, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return nil, fmt.Errorf(
			"param=%s by key=%s not a valid float: %s: %v",
			param,
			key,
			err,
			core_errors.ErrInvalidArgument,
		)
	}

	return &val, nil
}

func GetTimeQueryParam(r *http.Request, key string, layout string) (*time.Time, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, fmt.Errorf(
			"no key=%s in query param: %w",
			key,
			core_errors.ErrInvalidArgument,
		)
	}

	val, err := time.Parse(layout, param)
	if err != nil {
		return nil, fmt.Errorf(
			"param=%s by key=%s not a valid date: %s: %v",
			param,
			key,
			err,
			core_errors.ErrInvalidArgument,
		)
	}

	return &val, nil
}
