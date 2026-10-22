package favorites_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/jabrail059/weather-dashboard/internal/core/transport/http/errors"
)

func (r *FavoritesPostgresRepository) DeleteFavorite(
	ctx context.Context,
	city string,
	value string,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	DELETE FROM dashboard.favorites
	WHERE session_id=$1 AND name=$2;
	`

	cmdTag, err := r.pool.Exec(ctx, query, value, city)
	if err != nil {
		return fmt.Errorf(
			"exec query: %w",
			err,
		)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf(
			"favorite city='%s': %w",
			city,
			core_errors.ErrNotFound,
		)
	}

	return nil
}
