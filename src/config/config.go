package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port           int16  `default:"8080" split_words:"true"`
	SuperSecretKey string `default:"supersecret" split_words:"true"`
}

func Load() *Config {

	c := &Config{}
	err := envconfig.Process("STARTER", c)

	if err != nil {
		log.Fatal(err.Error())
	}

	return c
}
