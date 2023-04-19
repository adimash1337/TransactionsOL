package config

import (
	"github.com/caarlos0/env/v6"
	"log"
)

type Config struct {
	Port     string `env:"PORT" envDefault:":8000"`
	Database string `env:"DATABASE" envDefault:"user=dimash password=dimash dbname=OneLab sslmode=disable host=localhost port=5432"`
	DbDriver string `envconfig:"DB_DRIVER" default:"postgres"`
}

func New() (*Config, error) {
	config := Config{}
	if err := env.Parse(&config); err != nil {
		log.Println(err)

	}
	return &config, nil
}
