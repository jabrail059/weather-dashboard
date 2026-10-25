package favorites_transport_http

import (
	"net/http"

	core_logger "github.com/jabrail059/weather-dashboard/internal/core/logger"
	core_http_response "github.com/jabrail059/weather-dashboard/internal/core/transport/http/response"
	core_http_session "github.com/jabrail059/weather-dashboard/internal/core/transport/http/session"
)

// DeleteFavorites 	godoc
// @Summary 		Удалить из избранных
// @Description 	Удалить населенный пункт из списка избранных, сохраненных в системе
// @Tags 			favorites
// @Produce 		json
// @Param 			city 		query 		string 		true					"Название удаляемого населенного пункта"
// @Success 		204															"Успешное удаление населенного пункта"
// @Failure 		400 		{object} 	core_http_response.ErrorResponse 	"Bad Request"
// @Failure 		404 		{object} 	core_http_response.ErrorResponse 	"Not Found"
// @Failure 		500 		{object} 	core_http_response.ErrorResponse 	"Internal Server Error"
// @Router 			/favorites 	[delete]
func (h *FavoritesHTTPHander) DeleteFavorite(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	city := GetCityQueryParam(r)

	cookie, isNew := core_http_session.GetOrCreate(r)
	if isNew {
		http.SetCookie(w, cookie)
	}
	value := cookie.Value

	cities, err := h.favoritesService.DeleteFavorite(ctx, city, value)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to delete favorite city",
		)

		return
	}

	if cities.CacheErr != nil {
		log.WarnError("cache error", cities.CacheErr)
	}

	w.Header().Set("HX-Trigger", "favoritesChanged")
	responseHandler.NoContentResponse()
}
