package favorites_transport_http

import (
	"net/http"

	core_logger "github.com/jabrail059/weather-dashboard/internal/core/logger"
	core_http_response "github.com/jabrail059/weather-dashboard/internal/core/transport/http/response"
	core_http_session "github.com/jabrail059/weather-dashboard/internal/core/transport/http/session"
)

type GetFavoritesResponse struct {
	Cities   []string `json:"cities" example:"Moscow,Berlin"`
	Presence bool     `json:"presence" example:"true"`
}

// Getfavorites godoc
// @Summary 	Получить избранные
// @Description Получить список избранных населенных пунктов, сохраненных в системе
// @Tags 		favorites
// @Produce 	html,json
// @Success 	200 		{string} 	string 								"HTML favorites fragment"
// @Failure 	400 		{object} 	core_http_response.ErrorResponse 	"Bad Request"
// @Failure 	404 		{object} 	core_http_response.ErrorResponse 	"Not Found"
// @Failure 	500 		{object} 	core_http_response.ErrorResponse 	"Internal Server Error"
// @Router 		/favorites 	[get]
func (h *FavoritesHTTPHander) GetFavorites(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	cookie, isNew := core_http_session.GetOrCreate(r)
	if isNew {
		http.SetCookie(w, cookie)
	}
	value := cookie.Value

	favoritesList, err := h.favoritesService.GetFavorites(ctx, value)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get favorites",
		)

		return
	}

	if favoritesList.CacheErr != nil {
		log.WarnError("cache error", favoritesList.CacheErr)
	}

	var response GetFavoritesResponse
	response.Cities = favoritesList.Cities
	response.Presence = favoritesList.Presence

	if err := responseHandler.RenderPartHTML(response, "favorites.html"); err != nil {
		responseHandler.ErrorResponse(
			err,
			"render favorites HTML template",
		)

		return
	}
}
