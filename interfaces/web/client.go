package web

import "github.com/gorilla/websocket"

type Client struct {
	connection *websocket.Conn
}

func (c *Client) Emit(message interface{}) error {
	return c.connection.WriteJSON(message)
}

func (c *Client) ReadMessage() chan []byte {
	messages := make(chan []byte)
	go func() {
		for {
			_, message, err := c.connection.ReadMessage()
			if err != nil {
				close(messages)
				return
			}
			messages <- message
		}
	}()
	return messages

}

func NewClient(c *websocket.Conn) *Client {
	return &Client{
		connection: c,
	}
}
