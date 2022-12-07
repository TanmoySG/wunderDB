package config

import "github.com/caarlos0/env/v6"

func LoadConfig() (*Config, error) {
	config := &Config{}
	if err := env.Parse(&config); err != nil {
		return nil, err
	}
	return config, nil
}

type Config struct {
	AdminID       string `env:"ADMIN_ID" envDefault:"admin"`
	AdminPassword string `env:"ADMIN_PASSWORD" envDefault:""` 
}
