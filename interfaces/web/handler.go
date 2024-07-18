package web

import (
	"github.com/chainpusher/chainpusher/interfaces/web/socket"
	"github.com/sirupsen/logrus"
	"net/http"
)

type SocketHandler struct {
	clients *socket.Clients

	processor MessageProcessor
}

func (s *SocketHandler) Handle(w http.ResponseWriter, r *http.Request) {
	logrus.Debugf("New connection from %s", r.RemoteAddr)
	client, err := s.clients.Upgrade(w, r)

	if err != nil {
		logrus.Errorf("Failed to upgrade connection to WebSocket: %s", err)
		return
	}
	logrus.Debugf("Connection upgrade from %s to %d", r.RemoteAddr, client.GetId())

	defer s.close(client)

	s.readMessage(client)

	logrus.Debugf("Connection closed from %s to %d", r.RemoteAddr, client.GetId())
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
