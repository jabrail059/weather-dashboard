package favorites_postgres_repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jabrail059/weather-dashboard/internal/core/domain"
)

func (r *FavoritesPostgresRepository) AddFavorite(
	ctx context.Context,
	city string,
	value string,
) (domain.FavoritesList, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO dashboard.favorites (session_id, name, created_at, last_seen_at)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (session_id, name)
	DO UPDATE SET last_seen_at = EXCLUDED.last_seen_at
	RETURNING id, session_id, name, created_at, last_seen_at;
	`

	now := time.Now()
	row := r.pool.QueryRow(ctx, query, value, city, now, now)

	var favoritesModel FavoritesModel
	err := row.Scan(
		&favoritesModel.id,
		&favoritesModel.session_id,
		&favoritesModel.name,
		&favoritesModel.created_at,
		&favoritesModel.last_seen_at,
	)

	if err != nil {
		return domain.FavoritesList{}, fmt.Errorf("scan error: %w", err)
	}

	favorites := FavoriteListFromModel(favoritesModel)
	return favorites, nil
}
