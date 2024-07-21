package socket

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type ClientImpl struct {
	connection *websocket.Conn

	id int64

	logger *logrus.Entry
}

func (c *ClientImpl) Emit(message interface{}) error {
	return c.connection.WriteJSON(message)
}

func (c *ClientImpl) Close() error {
	return c.connection.Close()
}

func (c *ClientImpl) Read() chan []byte {
	messages := make(chan []byte)
	go func() {
		for {
			_, message, err := c.connection.ReadMessage()
			if err != nil {
				close(messages)
				break
			}
			messages <- message
		}
	}()
	return messages

}

func (c *ClientImpl) GetId() int64 {
	return c.id
}

func NewClient(id int64, c *websocket.Conn) Client {
	return &ClientImpl{
		id:         id,
		connection: c,
		logger:     logrus.WithFields(logrus.Fields{"id": id}),
	}
}
