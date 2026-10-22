package favorites_postgres_repository

import core_postgres_pool "github.com/jabrail059/weather-dashboard/internal/core/repository/postgres/pool"

type FavoritesPostgresRepository struct {
	pool core_postgres_pool.Pool
}

func NewFavoritesRepository(
	pool core_postgres_pool.Pool,
) *FavoritesPostgresRepository {
	return &FavoritesPostgresRepository{
		pool: pool,
	}
}
