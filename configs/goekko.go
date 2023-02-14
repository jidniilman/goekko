package configs

import (
	"os"
)

type GoekkoConfig struct {
	AppName string
	// HasuraEndpoint string
	// HasuraSecret   string
}

func LoadGoekkoConfig() GoekkoConfig {
	// hasuraEndpoint := os.Getenv("HASURA_ENDPOINT")
	// hasuraSecret := os.Getenv("HASURA_SECRET")

	return GoekkoConfig{
		AppName: os.Getenv("APP_NAME"),
		// HasuraEndpoint: hasuraEndpoint,
		// HasuraSecret:   hasuraSecret,
	}
}
