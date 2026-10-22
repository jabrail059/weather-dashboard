package favorites_transport_http

import (
	"net/http"

	core_http_request "github.com/jabrail059/weather-dashboard/internal/core/transport/http/request"
)

func GetCityQueryParam(r *http.Request) *string {
	var (
		cityQueryParamKey = "city"
	)
	city := core_http_request.GetStringQueryParam(r, cityQueryParamKey)

	return city
}
