package web

import (
	"github.com/chainpusher/chainpusher/interfaces/facade/impl"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func GetTransactionHandler(writer http.ResponseWriter, request *http.Request) {
	serviceFacade := impl.NewTransactionServiceFacadeImpl()

	_, err := serviceFacade.GetTransactions(nil)

	if err != nil {

	}

	_, err = writer.Write([]byte("Hello"))
	if err != nil {
		return
	}
}

func WSHandler(writer http.ResponseWriter, request *http.Request) {
	connection, err := upgrader.Upgrade(writer, request, nil)

	if err != nil {
		return
	}

	ctx := clients.Add(connection)

	go func(ctx *ClientContext) {
		for {
			_, _, err := ctx.GetConnection().ReadMessage()
			if err != nil {
				return
			}
		}
	}(ctx)

	defer clients.Remove(connection)
}
