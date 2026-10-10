package forecast_openmeteo_repository

// func (r *ForecastRepository) GetCurrentForecast(
// 	ctx context.Context,
// 	latitude float64,
// 	longitude float64,
// ) (*domain.CurrentForecast, error) {
// 	var currentResponse CurrentResponse

// 	apiURL, err := r.buildForecastURL(
// 		latitude,
// 		longitude,
// 		currentURLKey,
// 		currentURLQueryParams,
// 	)
// 	if err != nil {
// 		return nil, fmt.Errorf("build forecast url: %w", err)
// 	}

// 	err = r.getResponseFromOpenMeteo(
// 		ctx,
// 		apiURL,
// 		&currentResponse,
// 	)
// 	if err != nil {
// 		return nil, fmt.Errorf("get response from openmeteo: %w", err)
// 	}

// 	currentForecast := currentForecastFromResponse(currentResponse)
// 	//key := fmt.Sprintf("forecast:%v:%v", latitude, longitude)

// 	// value, err := redis.Client().Get(ctx, key).Result()
// 	// if err != nil {
// 	// 	return nil, http.StatusNotFound, fmt.Errorf("отсутствуют данные о текущей погоде")
// 	// }
// 	// if err := json.Unmarshal([]byte(value), &openMeteoResponse); err != nil {
// 	// 	return nil, http.StatusInternalServerError, fmt.Errorf("не удалось получить данные")
// 	// }
// 	return &currentForecast, nil
// }
