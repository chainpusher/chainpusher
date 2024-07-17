package web

import "net/http"

type SocketHandler struct {
	clients *Clients

	processor MessageProcessor
}

func (s *SocketHandler) Handle(w http.ResponseWriter, r *http.Request) {
	client, err := clients.Upgrade(w, r)
	if err != nil {
		return
	}

	defer s.close(client)

	s.readMessage(client)
}

func (s *SocketHandler) close(client *Client) {
	err := s.clients.Close(client)
	if err != nil {
		return
	}
}

func (s *SocketHandler) readMessage(client *Client) {
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
