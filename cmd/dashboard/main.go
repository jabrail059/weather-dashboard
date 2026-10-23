package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	core_config "github.com/jabrail059/weather-dashboard/internal/core/config"
	core_logger "github.com/jabrail059/weather-dashboard/internal/core/logger"
	core_pgx_pool "github.com/jabrail059/weather-dashboard/internal/core/repository/postgres/pool/pgx"
	core_goredis_client "github.com/jabrail059/weather-dashboard/internal/core/repository/redis/client/go-redis"
	core_http_middleware "github.com/jabrail059/weather-dashboard/internal/core/transport/http/middleware"
	core_http_server "github.com/jabrail059/weather-dashboard/internal/core/transport/http/server"
	favorites_postgres_repository "github.com/jabrail059/weather-dashboard/internal/features/favorites/repository/postgres"
	favorites_redis_repository "github.com/jabrail059/weather-dashboard/internal/features/favorites/repository/redis"
	favorites_service "github.com/jabrail059/weather-dashboard/internal/features/favorites/service"
	favorites_transport_http "github.com/jabrail059/weather-dashboard/internal/features/favorites/transport/http"
	forecast_openmeteo_repository "github.com/jabrail059/weather-dashboard/internal/features/forecast/repository/openmeteo"
	forecast_postgres_repository "github.com/jabrail059/weather-dashboard/internal/features/forecast/repository/postgres"
	forecast_redis_repository "github.com/jabrail059/weather-dashboard/internal/features/forecast/repository/redis"
	forecast_service "github.com/jabrail059/weather-dashboard/internal/features/forecast/service"
	forecast_transport_http "github.com/jabrail059/weather-dashboard/internal/features/forecast/transport/http"
	web_api_repository "github.com/jabrail059/weather-dashboard/internal/features/web/repository/api"
	web_redis_repository "github.com/jabrail059/weather-dashboard/internal/features/web/repository/redis"
	web_service "github.com/jabrail059/weather-dashboard/internal/features/web/service"
	web_transport_http "github.com/jabrail059/weather-dashboard/internal/features/web/transport/http"
	"go.uber.org/zap"
)

func main() {
	cfg := core_config.NewConfigMust()
	time.Local = cfg.TimeZone

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	config_log := core_logger.NewConfigMust()
	logger, err := core_logger.NewLogger(&config_log)
	if err != nil {
		fmt.Println("failed to init application logger: %w", err)
		os.Exit(1)
	}
	defer logger.Close()

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	redisClient, err := core_goredis_client.NewClient(
		ctx,
		core_goredis_client.NewConfigMust(),
	)
	if err != nil {
		logger.Fatal("failed to init redis client", zap.Error(err))
	}

	pool, err := core_pgx_pool.NewPool(
		ctx,
		core_pgx_pool.NewConfigMust(),
	)
	if err != nil {
		logger.Fatal("failed to init postgres connection pool", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("application time zone", zap.Any("zone", time.Local))

	logger.Debug("Initializing feature", zap.String("feature", "forecast"))
	forecastOpenmeteoRepository := forecast_openmeteo_repository.NewForecastOpenmeteoRepository(
		client,
		cfg.GeocodingBaseURL,
		cfg.ForecastBaseURL,
	)
	forecastPostgresRepository := forecast_postgres_repository.NewForecastPostgresRepository(
		pool,
	)
	forecastRedisRepository := forecast_redis_repository.NewForecastRedisRepository(redisClient)
	forecastService := forecast_service.NewForecastService(
		forecastOpenmeteoRepository,
		forecastPostgresRepository,
		forecastRedisRepository,
	)
	forecastHTTPHandler := forecast_transport_http.NewForecastHTTPHandler(forecastService)

	logger.Debug("Initializing feature", zap.String("feature", "web"))
	webRedisRepository := web_redis_repository.NewWebRedisRepository(redisClient)
	webApiRepository := web_api_repository.NewWebApiRepository(client, cfg.IpApiBaseURL)
	webService := web_service.NewWebService(webRedisRepository, webApiRepository)
	webHTTPHandler := web_transport_http.NewWebHTTPHadnler(webService)

	logger.Debug("Initializing feature", zap.String("feature", "favorites"))
	favoritesPostgresRepository := favorites_postgres_repository.NewFavoritesRepository(pool)
	favoritesRedisRepository := favorites_redis_repository.NewFavoritesRedisRepository(redisClient)
	favoritesService := favorites_service.NewFavoritesService(
		favoritesPostgresRepository,
		favoritesRedisRepository,
	)
	favoritesHTTPHandler := favorites_transport_http.NewFavoritesHTTPHander(favoritesService)

	logger.Debug("Initializing HTTP server")

	httpConfig := core_http_server.NewConfigMust()
	httpServer := core_http_server.NewHTTPServer(
		httpConfig,
		logger,
		core_http_middleware.CORS(httpConfig.AllowedOrigins),
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)

	apiVersionRouterV1 := core_http_server.NewApiVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouterV1.RegisterRoutes(forecastHTTPHandler.Routes()...)
	apiVersionRouterV1.RegisterRoutes(favoritesHTTPHandler.Routes()...)

	httpServer.RegisterApiRouters(
		apiVersionRouterV1,
	)

	httpServer.RegisterRoutes(webHTTPHandler.Routes()...)

	httpServer.RegisterStaticFiles("/static/", "./public/static")

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}

	if err := redisClient.Close(); err != nil {
		logger.Error("Close Redis client", zap.Error(err))
	}
}
