package web_test

import (
	"encoding/json"
	"github.com/chainpusher/chainpusher/interfaces/facade/dto"
	"github.com/chainpusher/chainpusher/interfaces/web"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSocket_Emit(t *testing.T) {
	done := make(chan struct{})

	processor := web.NewCallbackMessageProcessor(func(client *web.Client, message []byte) {
		var rpc *dto.JsonRpcDto
		_ = json.Unmarshal(message, &rpc)

		assert.Equal(t, "subscribe", rpc.Method)
		close(done)
	})

	go func() {
		server := web.NewServerTask("localhost", 8080, processor)
		_ = server.Start()
	}()

	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	assert.Nil(t, err)
	defer func(c *websocket.Conn) {
		err := c.Close()
		assert.Nil(t, err)
	}(c)

	_ = c.WriteJSON(&dto.JsonRpcDto{Method: "subscribe"})

	<-done
}
