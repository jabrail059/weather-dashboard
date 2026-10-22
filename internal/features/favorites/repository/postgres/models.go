package favorites_postgres_repository

import (
	"time"

	"github.com/jabrail059/weather-dashboard/internal/core/domain"
)

type FavoritesModel struct {
	id           int
	session_id   string
	name         string
	created_at   time.Time
	last_seen_at time.Time
}

func FavoritesListFromModels(
	models []FavoritesModel,
) domain.FavoritesList {
	cities := make([]string, len(models))
	for i, model := range models {
		cities[i] = model.name
	}

	return domain.FavoritesList{
		Cities:   cities,
		Presence: len(cities) > 0,
	}
}

func FavoriteListFromModel(
	model FavoritesModel,
) domain.FavoritesList {
	return domain.FavoritesList{
		Cities:   []string{model.name},
		Presence: true,
	}
}
