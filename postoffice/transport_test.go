package postoffice_test

import (
	"testing"

	"github.com/chainpusher/chainpusher/config"
	"github.com/chainpusher/chainpusher/postoffice"
)

func TestNewTransportHttpFromConfig(t *testing.T) {
	urls := []string{"http://localhost:8080"}
	transport := postoffice.NewTransportHttp(urls)
	if transport == nil {
		t.Error("Expected transport to be created")
	}

	if len(transport.(*postoffice.TransportHttp).Urls) != 1 {
		t.Error("Expected 1 URL")
	}

	if transport.(*postoffice.TransportHttp).Urls[0] != "http://localhost:8080" {
		t.Error("Expected http://localhost:8080")
	}
}

func TestNewTransportsFromConfig(t *testing.T) {
	cfg := &config.Config{
		Http: []config.HttpConfig{ // Change the type here
			{
				Url:           "http://localhost:8080",
				EncryptionKey: "key1",
			},
		},
	}

	transports := postoffice.NewTransportFromConfig(cfg)
	if len(transports) != 1 {
		t.Error("Expected 1 transports")
	}

	if _, ok := transports[0].(*postoffice.TransportHttp); !ok {
		t.Error("Expected HTTP transport")
	}
}
