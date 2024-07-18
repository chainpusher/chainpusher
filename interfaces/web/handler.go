package web

import (
	"github.com/chainpusher/chainpusher/interfaces/web/socket"
	"net/http"
)

type SocketHandler struct {
	clients *socket.Clients

	processor MessageProcessor
}

func (s *SocketHandler) Handle(w http.ResponseWriter, r *http.Request) {
	client, err := s.clients.Upgrade(w, r)
	if err != nil {
		return
	}

	defer s.close(client)

	s.readMessage(client)
}

func (s *SocketHandler) close(client *socket.Client) {
	err := s.clients.Close(client)
	if err != nil {
		return
	}
}

func (s *SocketHandler) readMessage(client *socket.Client) {
	messages := client.ReadMessage()
	for {

		select {
		case message, ok := <-messages:
			if !ok {
				return
			}
			s.processor.Process(client, message)
		}
	}
}
