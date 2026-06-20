package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/jabrail059/weather-dashboard/internal/models"
	"github.com/jabrail059/weather-dashboard/internal/storage"
	_ "modernc.org/sqlite"
)

type Storage struct {
	db *sql.DB
}

func New(connStr string) (*Storage, error) {
	db, err := sql.Open("sqlite", connStr)
	if err != nil {
		return nil, fmt.Errorf("Не удалось открыть базу данных: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("Не удалось подключиться к базе данных: %w", err)
	}

	slog.Info("SQLite успешно подключен")
	return &Storage{db: db}, nil
}

func (s *Storage) Save(ctx context.Context, geoResult *models.Result) error {
	q := `INSERT INTO cities (name, latitude, longitude) VALUES (?, ?, ?)`
	if _, err := s.db.ExecContext(ctx, q, geoResult.Name, geoResult.Latitude, geoResult.Longitude); err != nil {
		return fmt.Errorf("Не удалось сохранить данные: %w", err)
	}
	return nil
}

func (s *Storage) Select(ctx context.Context, name string) (*models.Result, error) {
	q := `SELECT latitude, longitude FROM cities WHERE name=?`
	var result = &models.Result{Name: name}
	err := s.db.QueryRowContext(ctx, q, name).Scan(&result.Latitude, &result.Longitude)
	if err == sql.ErrNoRows {
		return nil, storage.ErrCityNotSaved
	}
	if err != nil {
		return nil, fmt.Errorf("Не удалось получить данные из базы данных: %w", err)
	}

	return result, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}
