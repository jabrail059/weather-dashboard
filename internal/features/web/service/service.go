package web_service

import (
	"context"

	"github.com/jabrail059/weather-dashboard/internal/core/domain"
)

type WebService struct {
	webRedisRepository WebRedisRepository
	webApiRepository   WebApiRepository
}

type WebRedisRepository interface {
	GetSearchHistory(
		ctx context.Context,
		value string,
	) (domain.History, error)
}

type WebApiRepository interface {
	GetCityByIP(
		ctx context.Context,
		ip string,
	) (domain.CityReport, error)
}

func NewWebService(
	webRedisRepository WebRedisRepository,
	webApiRepository WebApiRepository,
) *WebService {
	return &WebService{
		webRedisRepository: webRedisRepository,
		webApiRepository:   webApiRepository,
	}
}
