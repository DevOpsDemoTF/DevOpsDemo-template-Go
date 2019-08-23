package main

import (
	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"
)

type config struct {
	DebugLevel string `env:"DEBUG_LEVEL" envDefault:"DEBUG"`
}

func parseConfig() *config {
	var cfg = config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal("error parsing configuration", err)
	}

	if err := configLog(&cfg); err != nil {
		log.Fatal("error configuring logger", err)
	}

	return &cfg
}

func configLog(cfg *config) error {
	log.SetFormatter(&log.JSONFormatter{})
	if level, err := log.ParseLevel(cfg.DebugLevel); err != nil {
		return err
	} else {
		log.SetLevel(level)
		return nil
	}
}
