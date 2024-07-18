package socket

import (
	"errors"
	"github.com/chainpusher/chainpusher/sys"
	"github.com/gorilla/websocket"
	"net/http"
	"sync/atomic"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	}}

type Clients struct {
	connections *sys.Map[int64, *Client]

	rooms map[string]*Room

	identity int64
}

func (c *Clients) Upgrade(w http.ResponseWriter, r *http.Request) (*Client, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	id := c.identity
	atomic.AddInt64(&c.identity, 1)

	client := NewClient(id, conn)
	c.connections.Put(id, client)

	return client, err
}

func (c *Clients) Get(clientId int64) (*Client, error) {
	client, ok := c.connections.Get(clientId)
	if !ok {
		return nil, errors.New("client not found")
	}
	return client, nil
}

func (c *Clients) Remove(clientId int64) error {
	client, err := c.Get(clientId)
	if err != nil {
		return err
	}
	c.connections.Remove(clientId)

	err = client.Close()
	return err
}

func (c *Clients) Join(clientId int64, roomName string) error {
	client, err := c.Get(clientId)
	if err != nil {
		return err
	}

	room := c.rooms[roomName]
	if room == nil {
		return errors.New("room not found")
	}

	room.Join(client)

	return nil
}

//func (c *Clients) SendAll(message interface{}) error {
//	for conn := range *c.connections {
//		_ = conn.WriteJSON(message)
//	}
//
//	return nil
//}

func (c *Clients) Close(client *Client) error {
	err := client.connection.Close()
	return err
}

func NewClients() *Clients {
	rooms := make(map[string]*Room)

	rooms["subscribe"] = NewRoom("subscribe")

	return &Clients{
		connections: sys.NewMap[int64, *Client](),
		identity:    0,
		rooms:       rooms,
	}
}
