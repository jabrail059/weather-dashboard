package forecast_redis_repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jabrail059/weather-dashboard/internal/core/domain"
)

func (r *ForecastRedisRepository) AddForecast(
	ctx context.Context,
	forecast domain.Forecast,
) error {
	key := fmt.Sprintf("forecast:%v:%v", forecast.Latitude, forecast.Longitude)

	jsonRes, err := json.Marshal(forecast)
	if err != nil {
		return fmt.Errorf("json marshal: %w", err)
	}

	if err = r.client.Set(ctx, key, jsonRes, time.Minute*45).Err(); err != nil {
		return fmt.Errorf("redis client set: %w", err)
	}

	return nil
}
