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
	core_http_middleware "github.com/jabrail059/weather-dashboard/internal/core/transport/http/middleware"
	core_http_server "github.com/jabrail059/weather-dashboard/internal/core/transport/http/server"
	forecast_openmeteo_repository "github.com/jabrail059/weather-dashboard/internal/features/forecast/repository/openmeteo"
	forecast_service "github.com/jabrail059/weather-dashboard/internal/features/forecast/service"
	forecast_transport_http "github.com/jabrail059/weather-dashboard/internal/features/forecast/transport/http"
	"go.uber.org/zap"
)

// func init() {
// 	if err := godotenv.Load(); err != nil {
// 		log.Print("no .env file found")
// 	}
// }

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

	logger.Debug("application time zone", zap.Any("zone", time.Local))

	logger.Debug("Initializing feature", zap.String("feature", "forecast"))
	forecastRepository := forecast_openmeteo_repository.NewForecastRepository(
		client,
		cfg.GeocodingBaseURL,
		cfg.ForecastBaseURL,
	)
	forecastService := forecast_service.NewForecastService(forecastRepository)
	forecastHTTPHandler := forecast_transport_http.NewForecastHTTPHandler(forecastService)

	logger.Debug("Initializing HTTP server")

	httpConfig := core_http_server.NewConfigMust()
	httpServer := core_http_server.NewHTTPServer(
		httpConfig,
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)

	apiVersionRouterV1 := core_http_server.NewApiVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouterV1.RegisterRoutes(forecastHTTPHandler.Routes()...)

	httpServer.RegisterApiRouters(
		apiVersionRouterV1,
	)
	//httpServer.RegisterRoutes()

	httpServer.RegisterStaticFiles("/static/", "./static")

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}

	// cfg := config.New()

	// if err := app.Connect(cfg); err != nil {
	// 	log.Fatal(err.Error())
	// }

	// r := server.NewRouter()

	// srv := http.Server{
	// 	Addr:    ":" + cfg.AppPort,
	// 	Handler: r,
	// }

	// go func() {
	// 	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
	// 		log.Fatal(err)
	// 	}
	// }()

	// slog.Info("Server started, addr " + srv.Addr)

	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	// <-quit

	// slog.Info("Server shutting down")

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// if err := srv.Shutdown(ctx); err != nil {
	// 	log.Printf("server shutdown error: %v", err)
	// }

	// if storage := app.Storage(); storage != nil {
	// 	if err := storage.Close(); err != nil {
	// 		log.Printf("sqlite close error: %v", err)
	// 	}
	// }

	// if err := redis.Close(); err != nil {
	// 	log.Printf("redis close error: %v", err)
	// }
}
