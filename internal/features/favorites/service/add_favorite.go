package favorites_service

import (
	"context"
	"fmt"

	"github.com/jabrail059/weather-dashboard/internal/core/domain"
	core_errors "github.com/jabrail059/weather-dashboard/internal/core/transport/http/errors"
)

func (s FavoritesService) AddFavorite(
	ctx context.Context,
	city *string,
	value string,
) (domain.FavoritesList, error) {
	if city == nil {
		return domain.FavoritesList{}, fmt.Errorf(
			"city must be not NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	favorites, err := s.favoritesPostgresRepository.AddFavorite(ctx, *city, value)
	if err != nil {
		return domain.FavoritesList{Presence: false}, fmt.Errorf(
			"add favorite to repository: %w",
			err,
		)
	}

	if err := s.favoritesRedisRepository.InvalidateFavorites(ctx, value); err != nil {
		favorites.CacheErr = err
	}

	return favorites, nil
}
