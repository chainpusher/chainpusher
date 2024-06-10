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

	InfuraKey string `yaml:"infura_key"`

	BlockLoggingFile string `yaml:"logging_file"`

	TransactionLoggingFile string `yaml:"transaction_file"`

	IsTesting bool
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

	text := string(bytes)

	// logrus.SetReportCaller(true)
	// logrus.SetFormatter(&logrus.TextFormatter{
	// 	FullTimestamp: true,
	// 	CallerPrettyfier: func(f *runtime.Frame) (string, string) {
	// 		repo := strings.Split(f.File, "/")
	// 		return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", repo[len(repo)-1], f.Line)
	// 	},
	// })

	return ParseConfigFromYamlText(text)
}
