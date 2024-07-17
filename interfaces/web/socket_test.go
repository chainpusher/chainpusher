package web_test

import (
	"github.com/chainpusher/chainpusher/interfaces/web"
	"testing"
	"time"
)

func TestSocket_Emit(t *testing.T) {
	processor := web.NewCallbackMessageProcessor(func(client *web.Client, message []byte) {

	})

	go func() {
		server := web.NewServerTask("localhost", 8080, processor)
		server.Start()
	}()

	time.Sleep(100 * time.Millisecond)
}
