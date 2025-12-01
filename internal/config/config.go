package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

//парсер конфига

type Config struct {
	Env         string `json:"env" env-default:"local" env-required:"true"`
	HttpServer  `json:"http_server"`
	StoragePath string `json:"storage_path" env-required:"true"`
}

type HttpServer struct {
	Address     string        `json:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `json:"timeout" env-default:"10s"`
	IdleTimeout time.Duration `json:"idle-timeout" env-default:"60s"`
}

// инициализация - конструктор
func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH") //переменная окружения
	if configPath == "" {
		configPath = "config/local.json"
		log.Printf("CONFIG_PATH not set, using default: %s", configPath)
	}

	//check is file exist
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file %s does not exist", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("can not read config: %s", err)
	}

	return &cfg
}
