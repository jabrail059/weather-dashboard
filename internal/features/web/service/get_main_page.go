package web_service

import (
	"context"

	"github.com/jabrail059/weather-dashboard/internal/core/domain"
)

func (s *WebService) GetCityByIP(
	ctx context.Context,
	ip string,
) (domain.CityReport, error) {
	report, err := s.webApiRepository.GetCityByIP(ctx, ip)
	if err != nil {
		return domain.CityReport{Detected: false}, nil
	}

	return report, nil
}
