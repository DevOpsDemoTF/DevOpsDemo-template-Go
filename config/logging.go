package config

import log "github.com/sirupsen/logrus"

func configLogLevel(cfg *Config) error {
	if level, err := log.ParseLevel(cfg.DebugLevel); err != nil {
		return err
	} else {
		log.SetLevel(level)
		return nil
	}
}
