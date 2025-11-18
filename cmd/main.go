package main

//TODO: add interface dependencies. подгрузить скуль драйвер

import (
	"Order/internal/config"
	"Order/internal/http-server/middleware"
	"Order/internal/http-server/middleware/logger"
	"Order/internal/storage/postgresql"
	"fmt"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg) //это надо хранить в секретах

	log := setUpLogger(cfg.Env)

	log.Info("starting url-shortener", slog.String("env", cfg.Env))

	log.Debug("debug messages are enabled")

	_, err := postgresql.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", slog.Any("error", err))
		os.Exit(1)
	}

	router := gin.Default()

	router.Use(middleware.RequestID())
	router.Use(logger.New(log))
	router.Use(middleware.Recoverer(log))

	/*router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	*/router.Run(":8080")
}

func setUpLogger(env string) *slog.Logger { //конфигурация логгера
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
