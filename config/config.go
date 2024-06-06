package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type HttpConfig struct {
	Url           string `json:"url"`
	EncryptionKey string `json:"encryption_key"`
}

type Config struct {
	Logger struct {
		Level string `json:"level"`
	}
	Wallets  []string `json:"wallets"`
	Telegram struct {
		Tokens []string `json:"token"`
	}
	Http []HttpConfig `json:"http"`
}

func ParseConfigFromYamlText(text string) (*Config, error) {
	var config Config

	err := yaml.Unmarshal([]byte(text), &config)
	if err != nil {
		return nil, err
	}

	if config.Logger.Level == "" {
		config.Logger.Level = "INFO"
	}

	return &config, nil
}

func ParseConfigFromYaml(file string) (*Config, error) {

	fd, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(fd)
	if err != nil {
		return nil, err
	}

	return ParseConfigFromYamlText(string(bytes))
}
