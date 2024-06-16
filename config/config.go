package config

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
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
		Tokens []interface{} `json:"token"`
	}
	Http []HttpConfig `json:"http"`

	InfuraKey string `yaml:"infura_key"`

	BlockLoggingFile string `yaml:"logging_file"`

	TransactionLoggingFile string `yaml:"transaction_file"`

	IsTesting bool

	Kafka struct {
		// The block will serialize the message into the topic
		BlockTopic string `yaml:"block_topic"`

		RawBlockTopic string `yaml:"raw_block_topic"`

		Servers []struct {
			Address string `yaml:"address"`
			Port    int    `yaml:"port"`
		} `yaml:"servers"`
	} `yaml:"kafka"`
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

	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			repo := strings.Split(f.File, "/")
			fun := strings.Split(f.Function, "/")
			return fmt.Sprintf("%s()", fun[len(fun)-1]), fmt.Sprintf("%s:%d", repo[len(repo)-1], f.Line)
		},
	})

	return ParseConfigFromYamlText(text)
}
