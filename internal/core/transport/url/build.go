package core_http_url

import (
	"fmt"
	"net/url"
)

func BuildURL(
	baseURL string,
	params map[string]string,
) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("parse base URL: %w", err)
	}

	query := u.Query()

	for key, value := range params {
		query.Set(key, value)
	}

	u.RawQuery = query.Encode()
	return u.String(), nil
}
