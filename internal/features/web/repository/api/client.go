package web_api_repository

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (r *WebApiRepository) getResponseFromIpApi(
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
			"get user city: %w",
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
