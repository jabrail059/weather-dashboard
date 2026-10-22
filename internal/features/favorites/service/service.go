package favorites_service

import (
	"context"

	"github.com/jabrail059/weather-dashboard/internal/core/domain"
)

type FavoritesService struct {
	favoritesPostgresRepository FavoritesPostgresRepository
	favoritesRedisRepository    FavoritesRedisRepository
}

type FavoritesPostgresRepository interface {
	GetFavorites(
		ctx context.Context,
		value string,
	) (domain.FavoritesList, error)

	AddFavorite(
		ctx context.Context,
		city string,
		value string,
	) (domain.FavoritesList, error)

	DeleteFavorite(
		ctx context.Context,
		city string,
		value string,
	) error
}

type FavoritesRedisRepository interface {
	GetFavorites(
		ctx context.Context,
		value string,
	) (domain.FavoritesList, error)

	AddFavorites(
		ctx context.Context,
		value string,
		cities []string,
	) error

	InvalidateFavorites(
		ctx context.Context,
		value string,
	) error
}

func NewFavoritesService(
	favoritesPostgresRepository FavoritesPostgresRepository,
	favoritesRedisRepository FavoritesRedisRepository,
) *FavoritesService {
	return &FavoritesService{
		favoritesPostgresRepository: favoritesPostgresRepository,
		favoritesRedisRepository:    favoritesRedisRepository,
	}
}
