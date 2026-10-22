package favorites_transport_http

import (
	"context"
	"net/http"

	"github.com/jabrail059/weather-dashboard/internal/core/domain"
	core_http_server "github.com/jabrail059/weather-dashboard/internal/core/transport/http/server"
)

type FavoritesHTTPHander struct {
	favoritesService FavoritesService
}

type FavoritesService interface {
	GetFavorites(
		ctx context.Context,
		value string,
	) (domain.FavoritesList, error)

	AddFavorite(
		ctx context.Context,
		city *string,
		value string,
	) (domain.FavoritesList, error)

	DeleteFavorite(
		ctx context.Context,
		city *string,
		value string,
	) (domain.FavoritesList, error)
}

func NewFavoritesHTTPHander(
	favoritesService FavoritesService,
) *FavoritesHTTPHander {
	return &FavoritesHTTPHander{
		favoritesService: favoritesService,
	}
}

func (h *FavoritesHTTPHander) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/favorites",
			Handler: h.GetFavorites,
		},
		{
			Method:  http.MethodPost,
			Path:    "/favorites",
			Handler: h.AddFavorite,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/favorites",
			Handler: h.DeleteFavorite,
		},
	}
}
