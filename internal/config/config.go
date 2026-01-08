package config

import "github.com/caarlos0/env/v6"

type Config struct {
	TelegramBotToken string `env:"TELEGRAM_BOT_TOKEN" envDefault:""`

	OrderServiceURL    string `env:"ORDER_SERVICE_URL" envDefault:"http://localhost:8081"`
	CustomerServiceURL string `env:"CUSTOMER_SERVICE_URL" envDefault:"http://localhost:8082"`
	GoodServiceURL     string `env:"GOOD_SERVICE_URL" envDefault:"http://localhost:8083"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
