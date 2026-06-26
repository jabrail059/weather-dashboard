package main

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jabrail059/weather-dashboard/internal/app"
	"github.com/jabrail059/weather-dashboard/internal/config"
	"github.com/jabrail059/weather-dashboard/internal/server"
	"github.com/jabrail059/weather-dashboard/internal/storage/redis"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("no .env file found")
	}
}

func main() {
	cfg := config.New()

	if err := app.Connect(cfg); err != nil {
		log.Fatal(err.Error())
	}

	r := server.NewRouter()

	srv := http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	slog.Info("Server started, addr " + srv.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	slog.Info("Server shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("server shutdown error: %v", err)
	}

	if storage := app.Storage(); storage != nil {
		if err := storage.Close(); err != nil {
			log.Printf("sqlite close error: %v", err)
		}
	}

	if err := redis.Close(); err != nil {
		log.Printf("redis close error: %v", err)
	}
}
