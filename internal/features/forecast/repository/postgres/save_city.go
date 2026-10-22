package forecast_postgres_repository

import (
	"context"
	"fmt"
)

func (r *ForecastPostgresRepository) SaveCity(
	ctx context.Context,
	city string,
	latitude float64,
	longitude float64,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO dashboard.cities (name, latitude, longitude)
	VALUES ($1, $2, $3)
	ON CONFLICT (name) DO NOTHING;
	`

	if _, err := r.pool.Exec(ctx, query, city, latitude, longitude); err != nil {
		return fmt.Errorf("save city: %w", err)
	}

	return nil

}
