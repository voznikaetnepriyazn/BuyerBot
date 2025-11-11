package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

//парсер конфига

type Config struct {
	Env        string `yaml:"env" env-default:"local" env-required:"true"`
	HttpServer `yaml:"http_server"`
	//db
}

type HttpServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"10s"`
	IdleTimeout time.Duration `yaml:"iddle-timeout" env-default:"60s"`
}

// инициализация - конструктор
func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH") //переменная окружения
	if configPath == " " {
		log.Fatal("CONFIG_PATH is not set")
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
