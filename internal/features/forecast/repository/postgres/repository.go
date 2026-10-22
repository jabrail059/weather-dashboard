package forecast_postgres_repository

import core_postgres_pool "github.com/jabrail059/weather-dashboard/internal/core/repository/postgres/pool"

type ForecastPostgresRepository struct {
	pool core_postgres_pool.Pool
}

func NewForecastPostgresRepository(
	pool core_postgres_pool.Pool,
) *ForecastPostgresRepository {
	return &ForecastPostgresRepository{
		pool: pool,
	}
}
