package application_test

import (
	"github.com/chainpusher/blockchain/model"
	"github.com/chainpusher/chainpusher/application"
	"github.com/chainpusher/chainpusher/interfaces/facade/dto"
	"github.com/chainpusher/chainpusher/interfaces/web/socket"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestTinyBlockServiceImpl_Broadcast(t *testing.T) {

	var clients socket.Clients = socket.NewClients()
	var svc application.TinyBlockService = application.NewTinyBlockService(clients)

	c1 := socket.NewMemoryClient(1)
	c2 := socket.NewMemoryClient(2)

	clients.Add(c1)
	clients.Add(c2)

	_ = clients.Join(1, "subscribe")
	_ = clients.Join(2, "subscribe")

	svc.Broadcast(&model.Block{Height: big.NewInt(1)})

	room := clients.Room("subscribe")
	subscribers := room.GetClients()
	assert.Len(t, subscribers, 2)

	for _, subscriber := range subscribers {
		subscriber := subscriber.(*socket.MemoryClient)

		raw := subscriber.GetEmits()[0]
		event := raw.(*dto.JsonRpcEvent)
		block := event.Data.(*model.Block)
		assert.Equal(t, int64(1), block.Height.Int64())
	}
}
