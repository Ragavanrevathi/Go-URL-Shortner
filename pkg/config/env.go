package config

import (
	"log"

	"github.com/caarlos0/env"
)

type shortenConfig struct {
	Domain string `env:"APP_DOMAIN" envDefault:"localhost:8080"`
}

var Env shortenConfig

// Initializing env variables
func Init() {
	if err := env.Parse(&Env); err != nil {
		log.Fatalf("Failed to parse env: %+v", err)
	}
}
