package forecast_postgres_repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jabrail059/weather-dashboard/internal/core/domain"
	core_errors "github.com/jabrail059/weather-dashboard/internal/core/transport/http/errors"
)

func (r *ForecastPostgresRepository) GetCity(
	ctx context.Context,
	city string,
) (domain.GeocodeResult, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, name, latitude, longitude
	FROM dashboard.cities
	WHERE name=$1;
	`

	row := r.pool.QueryRow(ctx, query, city)
	var geocodeModel GeocodeModel

	err := row.Scan(
		&geocodeModel.ID,
		&geocodeModel.City,
		&geocodeModel.Latitude,
		&geocodeModel.Longitude,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.GeocodeResult{}, fmt.Errorf("city %q: %w", city, core_errors.ErrNotFound)
		}
		return domain.GeocodeResult{}, fmt.Errorf("scan error: %w", err)
	}

	result := geocodeDomainFromModel(geocodeModel)

	return result, nil
}
