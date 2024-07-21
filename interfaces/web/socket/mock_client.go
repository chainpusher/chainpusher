package socket

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type MockClient struct {
	id     int64
	client *websocket.Conn
	logger *logrus.Entry
}

func (m MockClient) GetId() int64 {
	return m.id
}

func (m MockClient) Emit(message interface{}) error {
	m.logger.Tracef("emit message %v", message)
	//err := m.client.WriteJSON(message)

	encoded, err := json.Marshal(message)
	if err != nil {
		return err
	}
	err = m.client.WriteMessage(websocket.TextMessage, encoded)
	return err
}

func (m MockClient) Read() chan []byte {
	r := make(chan []byte)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				m.logger.Errorf("panic: %v", r)
			}
		}()
		for {
			t, message, err := m.client.ReadMessage()
			m.logger.Debugf("read message type is (%v) and size is %d", t, len(message))
			m.logger.Debug(string(message))
			if err != nil {
				m.logger.Error(err)
				close(r)
				break
			}
			m.logger.Tracef("read message %v", string(message))
			r <- message
			m.logger.Debugf("write message to the chan")

		}
	}()

	return r
}

func (m MockClient) Close() error {
	return m.client.Close()
}

var i int64 = 0

func NewMockClient() Client {

	defer func() {
		i++
	}()

	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		panic(err)
	}
	return &MockClient{id: i, client: c, logger: logrus.WithField("clientId", i)}
}
