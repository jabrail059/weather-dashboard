package service

import (
	"github.com/jabrail059/weather-dashboard/storage"
)

var db storage.Storage

func SetStorage(database storage.Storage) {
	db = database
}

func Storage() storage.Storage {
	return db
}
