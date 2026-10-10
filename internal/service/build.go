package service

// func BuildHourlyWeather(forecast *models.Hourly, sunrise string, sunset string) ([]models.HourlyWeather, error) {
// 	hours := make([]models.HourlyWeather, 0)
// 	for i := 0; i < len(forecast.Time); i++ {
// 		parsedTime, err := time.Parse("2006-01-02T15:04", forecast.Time[i])
// 		if err != nil {
// 			return nil, fmt.Errorf("возникла ошибка при преобразовании даты: %v", err)
// 		}
// 		formattedTime := parsedTime.Format("15:04")
// 		hourlyTemp := models.HourlyWeather{
// 			Time:        formattedTime,
// 			Temperature: math.Round(forecast.Temperature[i]),
// 		}

// 		if forecast.WeatherCode[i] <= 2 {
// 			if formattedTime <= sunrise || formattedTime >= sunset {
// 				hourlyTemp.ImageSource = weather.NightImageSources[forecast.WeatherCode[i]]
// 			} else {
// 				hourlyTemp.ImageSource = weather.ImageSources[forecast.WeatherCode[i]]
// 			}
// 		} else {
// 			hourlyTemp.ImageSource = weather.ImageSources[forecast.WeatherCode[i]]
// 		}
// 		hours = append(hours, hourlyTemp)
// 	}
// 	return hours, nil
// }

// func BuildOneDayForecast(forecast *models.Hourly, date string) (*models.Hourly, error) {
// 	newForecast := &models.Hourly{
// 		Time:        make([]string, 0),
// 		Temperature: make([]float64, 0),
// 		WeatherCode: make([]int, 0),
// 	}

// 	for i := 0; i < len(forecast.Time); i++ {
// 		if strings.HasPrefix(forecast.Time[i], date) {
// 			newForecast.Time = append(newForecast.Time, forecast.Time[i])
// 			newForecast.Temperature = append(newForecast.Temperature, forecast.Temperature[i])
// 			newForecast.WeatherCode = append(newForecast.WeatherCode, forecast.WeatherCode[i])
// 		}
// 	}

// 	if len(newForecast.Time) == 0 {
// 		return nil, fmt.Errorf("прогноз на %s не найден", date)
// 	}

// 	return newForecast, nil
// }
