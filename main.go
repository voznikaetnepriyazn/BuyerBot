package main

import (
	"fmt"
	"os"
	"log/slog"
	".\order\internal\config"
)

const (
	envLocal = "local"
	envDev = "dev"
	envProd = "prod"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg) //это надо хранить в секретах

	log := setupLogger(cfg.Env)

	log.Info("starting url-shortener", slog.String("env", cfg.Env))

	log.Debug("debug messages are enabled")

	storage, err := postgresql.New(cfg.StoragePath)
	if err != nil{
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}
}

func setUpLogger(env string) {//конфигурация логгера
	var log *slog.Logger

	switch env{
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
