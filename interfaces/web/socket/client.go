package socket

import "github.com/gorilla/websocket"

type Client struct {
	connection *websocket.Conn

	id int64
}

func (c *Client) Emit(message interface{}) error {
	return c.connection.WriteJSON(message)
}

func (c *Client) Close() error {
	return c.connection.Close()
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

func (c *Client) GetId() int64 {
	return c.id
}

func NewClient(id int64, c *websocket.Conn) *Client {
	return &Client{
		id:         id,
		connection: c,
	}
}
