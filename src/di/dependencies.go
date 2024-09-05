package di

import "github.com/tobyscott25/gin-air-docker-boilerplate/src/config"

type Dependencies struct {
	Config *config.Config
}

func InitDependencies(config *config.Config) (*Dependencies, error) {
	return &Dependencies{
		Config: config,
	}, nil
}
