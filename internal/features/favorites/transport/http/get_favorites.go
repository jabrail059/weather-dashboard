package favorites_transport_http

import (
	"net/http"

	core_logger "github.com/jabrail059/weather-dashboard/internal/core/logger"
	core_http_response "github.com/jabrail059/weather-dashboard/internal/core/transport/http/response"
	core_http_session "github.com/jabrail059/weather-dashboard/internal/core/transport/http/session"
)

func (h *FavoritesHTTPHander) GetFavorites(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	cookie, isNew := core_http_session.GetOrCreate(r)
	if isNew {
		http.SetCookie(w, cookie)
	}
	value := cookie.Value

	cities, err := h.favoritesService.GetFavorites(ctx, value)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get favorites",
		)

		return
	}

	if cities.CacheErr != nil {
		log.WarnError("cache error", cities.CacheErr)
	}

	if err := responseHandler.RenderPartHTML(cities, "favorites.html"); err != nil {
		responseHandler.ErrorResponse(
			err,
			"render favorites HTML template",
		)

		return
	}
}
