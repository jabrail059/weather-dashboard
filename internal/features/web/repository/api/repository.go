package web_api_repository

import "net/http"

type WebApiRepository struct {
	client   *http.Client
	getIpURL string
}

func NewWebApiRepository(
	client *http.Client,
	getIpURL string,
) *WebApiRepository {
	return &WebApiRepository{
		client:   client,
		getIpURL: getIpURL,
	}
}
