package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	S3Config
}

func (c *Config) LoadFromEnv() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("can`t load from env: %v", err)
	}

	cfg := &Config{}

	if err := env.Parse(&cfg.S3Config); err != nil {
		log.Fatalf("can`t load S3 from env: %v", err)
	}
	
	return cfg
}
