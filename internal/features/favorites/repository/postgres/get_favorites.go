package favorites_postgres_repository

import (
	"context"
	"fmt"

	"github.com/jabrail059/weather-dashboard/internal/core/domain"
)

func (r *FavoritesPostgresRepository) GetFavorites(
	ctx context.Context,
	value string,
) (domain.FavoritesList, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, session_id, name, created_at, last_seen_at
	FROM dashboard.favorites
	WHERE session_id=$1;
	`

	rows, err := r.pool.Query(ctx, query, value)
	if err != nil {
		return domain.FavoritesList{}, fmt.Errorf("select favorites: %w", err)
	}
	defer rows.Close()

	var favoritesModels []FavoritesModel
	for rows.Next() {
		var favoritesModel FavoritesModel
		err := rows.Scan(
			&favoritesModel.id,
			&favoritesModel.session_id,
			&favoritesModel.name,
			&favoritesModel.created_at,
			&favoritesModel.last_seen_at,
		)
		if err != nil {
			return domain.FavoritesList{}, fmt.Errorf("scan error: %w", err)
		}

		favoritesModels = append(favoritesModels, favoritesModel)
	}

	if err := rows.Err(); err != nil {
		return domain.FavoritesList{}, fmt.Errorf("next rows: %w", err)
	}

	cities := FavoritesListFromModels(favoritesModels)
	return cities, nil
}
