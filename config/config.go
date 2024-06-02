package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Logger struct {
		Level string `json:"level"`
	}
	Wallets []string `json:"wallets"`
}

func ParseConfigFromYaml(file string) (*Config, error) {
	var config Config

	fd, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(fd)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	if config.Logger.Level == "" {
		config.Logger.Level = "INFO"
	}

	return &config, nil
}
