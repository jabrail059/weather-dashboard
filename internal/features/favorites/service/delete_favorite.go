package favorites_service

import (
	"context"
	"fmt"

	"github.com/jabrail059/weather-dashboard/internal/core/domain"
	core_errors "github.com/jabrail059/weather-dashboard/internal/core/transport/http/errors"
)

func (s *FavoritesService) DeleteFavorite(
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

	if err := s.favoritesPostgresRepository.DeleteFavorite(ctx, *city, value); err != nil {
		return domain.FavoritesList{}, fmt.Errorf("delete favorite from repository: %w", err)
	}

	if err := s.favoritesRedisRepository.InvalidateFavorites(ctx, value); err != nil {
		return domain.FavoritesList{CacheErr: err}, nil
	}

	return domain.FavoritesList{}, nil
}
