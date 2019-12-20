package config

import (
	env "github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	DebugLevel string `env:"DEBUG_LEVEL" envDefault:"DEBUG"`
}

func ParseConfig() *Config {
	var cfg = Config{}
	log.SetFormatter(&log.JSONFormatter{})

	if err := env.Parse(&cfg); err != nil {
		log.Fatal("error parsing configuration", err)
	}

	if err := configLogLevel(&cfg); err != nil {
		log.Fatal("error configuring logger", err)
	}

	return &cfg
}
