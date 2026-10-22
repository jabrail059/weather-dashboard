package forecast_redis_repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jabrail059/weather-dashboard/internal/core/domain"
)

func (r *ForecastRedisRepository) GetForecast(
	ctx context.Context,
	latitude float64,
	longitude float64,
) (domain.Forecast, error) {
	key := fmt.Sprintf("forecast:%v:%v", latitude, longitude)

	var res domain.Forecast

	value, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return domain.Forecast{}, fmt.Errorf("redis client get: %w", err)
	}

	if err = json.Unmarshal([]byte(value), &res); err != nil {
		return domain.Forecast{}, fmt.Errorf("json unmarshal: %w", err)
	}
	return res, nil
}
