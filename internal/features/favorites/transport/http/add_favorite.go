package favorites_transport_http

import (
	"net/http"

	core_logger "github.com/jabrail059/weather-dashboard/internal/core/logger"
	core_http_response "github.com/jabrail059/weather-dashboard/internal/core/transport/http/response"
	core_http_session "github.com/jabrail059/weather-dashboard/internal/core/transport/http/session"
)

type AddFavoriteResponse struct {
	Cities   []string `json:"cities" example:"Moscow,Berlin"`
	Presence bool     `json:"presence" example:"true"`
}

// AddFavorite 	godoc
// @Summary 	Добавить в избранные
// @Description Добавить населенный пункт в список избранных, сохраненных в системе
// @Tags 		favorites
// @Produce 	json
// @Param 		city		query		string			true				"Название населенного пункта"
// @Success 	200 		{object}	AddFavoriteResponse					"Успешное добавление населенного пункта в список избранных"
// @Failure 	400 		{object} 	core_http_response.ErrorResponse 	"Bad Request"
// @Failure 	404 		{object} 	core_http_response.ErrorResponse 	"Not Found"
// @Failure 	500 		{object} 	core_http_response.ErrorResponse 	"Internal Server Error"
// @Router 		/favorites 	[post]
func (h *FavoritesHTTPHander) AddFavorite(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
	city := GetCityQueryParam(r)

	cookie, isNew := core_http_session.GetOrCreate(r)
	if isNew {
		http.SetCookie(w, cookie)
	}
	value := cookie.Value

	favoritesList, err := h.favoritesService.AddFavorite(ctx, city, value)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed add city to favorites",
		)

		return
	}

	if favoritesList.CacheErr != nil {
		log.WarnError("cache error", favoritesList.CacheErr)
	}

	var response AddFavoriteResponse
	response.Cities = favoritesList.Cities
	response.Presence = favoritesList.Presence

	w.Header().Set("HX-Trigger", "favoritesChanged")
	responseHandler.JSONResponse(response, http.StatusOK)
}
