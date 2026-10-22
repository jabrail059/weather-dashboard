package favorites_service

import (
	"context"
	"fmt"

	"github.com/jabrail059/weather-dashboard/internal/core/domain"
)

func (s *FavoritesService) GetFavorites(
	ctx context.Context,
	value string,
) (domain.FavoritesList, error) {
	cached, cacheErr := s.favoritesRedisRepository.GetFavorites(ctx, value)
	if cacheErr == nil && cached.Presence {
		return cached, nil
	}

	favorites, err := s.favoritesPostgresRepository.GetFavorites(ctx, value)
	if err != nil {
		return domain.FavoritesList{CacheErr: cacheErr}, fmt.Errorf(
			"get favorites from repository: %w",
			err,
		)
	}

	if cacheErr == nil && favorites.Presence {
		cacheErr = s.favoritesRedisRepository.AddFavorites(ctx, value, favorites.Cities)
	}
	favorites.CacheErr = cacheErr
	return favorites, nil
}
