package configs

import (
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	GoekkoAppConfig GoekkoConfig
}

func GetConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	return &Config{
		GoekkoAppConfig: LoadGoekkoConfig(),
	}
}
