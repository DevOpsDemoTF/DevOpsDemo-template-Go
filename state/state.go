package state

import "github.com/butzist/DevOpsDemo-template-Go/config"

type State struct {
	Config  *config.Config
	Healthy bool
	// TODO put your database connections and other stateful components here
}

func Create(config *config.Config) (*State, error) {
	return &State{
		Config:  config,
		Healthy: true,
	}, nil
}
