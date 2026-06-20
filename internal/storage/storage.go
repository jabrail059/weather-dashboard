package storage

import (
	"context"
	"errors"

	"github.com/jabrail059/weather-dashboard/internal/models"
)

type Storage interface {
	Save(ctx context.Context, geoResult *models.Result) error
	Select(ctx context.Context, name string) (*models.Result, error)
	Close() error
}

var ErrCityNotSaved = errors.New("city not saved")
