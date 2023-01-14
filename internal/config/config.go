package config

import "github.com/caarlos0/env/v6"

func LoadConfig() (*Config, error) {
	config := Config{}
	if err := env.Parse(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

type Config struct {
	AdminID               string `env:"ADMIN_ID"`
	AdminPassword         string `env:"ADMIN_PASSWORD"`
	Port                  string `env:"PORT" envDefault:"8086"`
	PersistantStoragePath string `env:"PERSISTANT_STORAGE_PATH" envDefault:"wfs"`
}
