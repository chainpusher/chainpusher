package config_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/chainpusher/chainpusher/config"
)

func TestParseConfigFromYamlText(t *testing.T) {
	var text string = `wallets:
  - a
  - b

telegram:
  token: 7041850206:AAEUECtq8CwY-jTJ5SJIdE3Q2hzuWLWcW3s

http:
  - url: https://httpbin.org/post

logger:
  level: debug`
	cfg, err := config.ParseConfigFromYamlText(text)
	if err != nil {
		t.Error(err)
	}

	if len(cfg.Wallets) != 2 {
		t.Error("Expected 2 wallets")
	}

	if cfg.Http[0].Url != "https://httpbin.org/post" {
		t.Error("Expected https://httpbin.org/post")
	}
}

// Test when token is a string
func TestTelegramTokenIsString(t *testing.T) {

	var text string = `telegram:
  tokens:
    - "<token>"
`
	cfg, err := config.ParseConfigFromYamlText(text)
	if err != nil {
		t.Error(err)
	}

	if len(cfg.Telegram.Tokens) != 1 {
		t.Error("Expected 1 token")
	}

	if cfg.Telegram.Tokens[0].(string) != "<token>" {
		t.Error("Expected <token>")
	}
}

// Test when token is an object
func TestTelegramTokenIsObject(t *testing.T) {

	var text string = `telegram:
  tokens:
    - token: "<token>"
      chat_id: 1234
`

	cfg, err := config.ParseConfigFromYamlText(text)
	if err != nil {
		t.Error(err)
	}

	if len(cfg.Telegram.Tokens) != 1 {
		t.Error("Expected 1 token")
	}

	token := cfg.Telegram.Tokens[0].(map[string]interface{})
	if token["token"].(string) != "<token>" {
		t.Error("Expected <token>")
	}

	if token["chat_id"].(int) != 1234 {
		t.Error("Expected 1234")
	}

}

// Test when token is an object but chat_id is empty
func TestTelegramTokenIsObjectChatIdEmpty(t *testing.T) {

	var text string = `telegram:
  tokens:
    - token: "<token>"
`

	cfg, err := config.ParseConfigFromYamlText(text)
	if err != nil {
		t.Error(err)
	}

	if len(cfg.Telegram.Tokens) != 1 {
		t.Error("Expected 1 token")
	}

	token := cfg.Telegram.Tokens[0].(map[string]interface{})
	if token["token"].(string) != "<token>" {
		t.Error("Expected <token>")
	}

	if _, ok := token["chat_id"]; ok {
		t.Error("Expected chat_id to be empty")
	}
}

func TestConfig_KafkaIsValidated(t *testing.T) {
	var text string = `kafka:
  block_topic: block
  raw_block_topic: raw_block
  servers:
    - address: localhost
      port: 9092`
	cfg, err := config.ParseConfigFromYamlText(text)
	assert.Nil(t, err, "Error parsing config")
	assert.True(t, cfg.GetKafka().IsValidated(), "Expected Kafka config to be validated")
}

func TestConfig_KafkaIsInvalidated(t *testing.T) {
	var text string = `kafka:
  block_topic: block
  raw_block_topic: raw_block
  servers: []`
	cfg, err := config.ParseConfigFromYamlText(text)
	assert.Nil(t, err, "Error parsing config")
	assert.False(t, cfg.GetKafka().IsValidated(), "Expected Kafka config to be invalidated")
}
