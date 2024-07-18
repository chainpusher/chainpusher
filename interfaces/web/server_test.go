package web_test

import (
	"github.com/chainpusher/chainpusher/interfaces/web"
	"github.com/chainpusher/chainpusher/interfaces/web/socket"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestServer_Start(t *testing.T) {

	processor := web.NewCallbackMessageProcessor(func(client *socket.Client, message []byte) {

	})
	server := web.NewServerTask("127.0.0.1", 8080, processor, nil)

	go func() {
		_ = server.Start()
	}()

	time.Sleep(100 * time.Millisecond)

	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	assert.Nil(t, err)

	rsp, err := client.Do(req)
	assert.Nil(t, err)
	assert.NotNil(t, rsp)

	err = server.Stop()
	assert.Nil(t, err)
}
