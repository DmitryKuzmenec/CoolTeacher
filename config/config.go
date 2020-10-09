package config

import (
	"github.com/joho/godotenv"
	"github.com/vrischmann/envconfig"
)

// Config - struct of config
type Config struct {
}

// Init - get config
func Init() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	var conf Config
	if err := envconfig.Init(&conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
