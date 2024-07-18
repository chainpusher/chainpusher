package web_test

import (
	"encoding/json"
	"github.com/chainpusher/chainpusher/application"
	"github.com/chainpusher/chainpusher/interfaces/facade/dto"
	"github.com/chainpusher/chainpusher/interfaces/facade/impl"
	"github.com/chainpusher/chainpusher/interfaces/web"
	"github.com/chainpusher/chainpusher/interfaces/web/socket"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSocket_Emit(t *testing.T) {
	done := make(chan struct{})

	processor := web.NewCallbackMessageProcessor(func(client *socket.Client, message []byte) {
		var rpc *dto.JsonRpcDto
		_ = json.Unmarshal(message, &rpc)

		assert.Equal(t, "subscribe", rpc.Method)
		close(done)
	})

	go func() {
		server := web.NewServerTask("localhost", 8080, processor, nil)
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

func TestSocket_FunctionCall(t *testing.T) {
	done := make(chan *dto.JsonRpcResponseDto)

	processor := web.NewCallbackMessageProcessor(func(client *socket.Client, message []byte) {
		rsp := dto.JsonRpcResponseDto{
			Call: &dto.JsonRpcDto{
				Method: "ping",
			},
			Result: nil,
		}
		err := client.Emit(rsp)
		assert.Nil(t, err)
	})

	go func() {
		server := web.NewServerTask("localhost", 8080, processor, nil)
		_ = server.Start()
	}()

	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	assert.Nil(t, err)
	defer func(c *websocket.Conn) {
		err := c.Close()
		assert.Nil(t, err)
	}(c)

	go func() {
		var response *dto.JsonRpcResponseDto
		_, bytes, err := c.ReadMessage()

		assert.Nil(t, err)
		err = json.Unmarshal(bytes, &response)
		assert.Nil(t, err)

		done <- response
	}()

	_ = c.WriteJSON(&dto.JsonRpcDto{Method: "subscribe"})

	select {
	case rsp := <-done:
		assert.Equal(t, rsp.Call.Method, "ping")
		assert.Nil(t, rsp.Result)
		close(done)
		return
	}

}

func TestSocket_Subscribe(t *testing.T) {
	done := make(chan *dto.JsonRpcResponseDto)

	clients := socket.NewClients()
	applicationService := application.NewTinyBlockService(clients)
	processor := web.NewJsonRpcMessageProcess(impl.NewTinyBlockServiceFacade(applicationService))

	go func() {
		server := web.NewServerTask("localhost", 8080, processor, clients)
		_ = server.Start()
	}()

	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	assert.Nil(t, err)
	defer func(c *websocket.Conn) {
		err := c.Close()
		assert.Nil(t, err)
	}(c)

	go func() {
		var response *dto.JsonRpcResponseDto
		_, bytes, err := c.ReadMessage()

		assert.Nil(t, err)
		err = json.Unmarshal(bytes, &response)
		assert.Nil(t, err)

		done <- response
	}()

	_ = c.WriteJSON(&dto.JsonRpcDto{Method: "subscribe"})

	select {
	case rsp := <-done:
		assert.Equal(t, rsp.Call.Method, "subscribe")
		assert.Nil(t, rsp.Result)
		close(done)
		return
	}

}
