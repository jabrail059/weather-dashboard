package web_service

import (
	"context"

	"github.com/jabrail059/weather-dashboard/internal/core/domain"
)

func (s *WebService) GetSearchHistory(
	ctx context.Context,
	value string,
) (domain.HistoryList, error) {
	history, err := s.webRedisRepository.GetSearchHistory(ctx, value)
	if err != nil {
		return domain.HistoryList{
			Presence: false,
			CacheErr: err,
		}, nil
	}

	return history, nil
}
