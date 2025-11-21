package main

//TODO: add interface dependencies. подгрузить скуль драйвер, уникальный тип данных для ид

import (
	"Order/internal/config"
	"Order/internal/http-server/handlers/order/add"
	"Order/internal/http-server/handlers/order/delete"
	getall "Order/internal/http-server/handlers/order/getAll"
	getbyid "Order/internal/http-server/handlers/order/getById"
	isordercreated "Order/internal/http-server/handlers/order/isOrderCreated"
	"Order/internal/http-server/handlers/order/update"
	"Order/internal/http-server/middleware"
	"Order/internal/http-server/middleware/logger"
	services "Order/internal/services/order-service"
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

	dbStorage, err := postgresql.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", slog.Any("error", err))
		os.Exit(1)
	}

	urlService := services.NewService(dbStorage)

	router := gin.Default()

	router.Use(middleware.RequestID())
	router.Use(logger.New(log))
	router.Use(middleware.Recoverer(log))

	registerHandlers(router, log, urlService)

	addr := cfg.HttpServer.Address
	if addr == "" {
		addr = ":8080"
	}

	log.Info("starting server", slog.String("address", addr))
	if err := router.Run(addr); err != nil {
		log.Error("failed to start server", slog.Any("error", err))
		os.Exit(1)
	}
}

func registerHandlers(router *gin.Engine, log *slog.Logger, urlService *services.OrderStruct) {
	router.POST("url/add", add.New(log, service))

	router.GET("url/getById", getall.New(log, service))

	router.GET("url/getAll", getbyid.New(log, service))

	router.PUT("url/update", update.New(log, service))

	router.DELETE("url/delete", isordercreated.New(log, service))

	router.GET("url/isOrderCreated", delete.New(log, service))
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
