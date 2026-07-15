package sqlite

import (
	"path/filepath"
	"testing"

	"github.com/jabrail059/weather-dashboard/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSqlite(t *testing.T) {
	path := filepath.Join(t.TempDir(), "test.db")

	storage, err := New(path)
	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, storage.Close())
	})

	input := models.Result{
		ID:        1,
		Name:      "Moscow",
		Latitude:  55.75204,
		Longitude: 37.61781,
	}

	err = storage.Save(t.Context(), &input)
	require.NoError(t, err)

	result, err := storage.Select(t.Context(), input.Name)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, input.Name, result.Name)
	assert.Equal(t, input.Latitude, result.Latitude)
	assert.Equal(t, input.Longitude, result.Longitude)
}
