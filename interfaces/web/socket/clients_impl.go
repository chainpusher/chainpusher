package socket

import (
	"errors"
	"github.com/chainpusher/chainpusher/sys"
	"github.com/gorilla/websocket"
	"net/http"
	"sync/atomic"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  102400,
	WriteBufferSize: 102400,
	CheckOrigin: func(r *http.Request) bool {
		return true
	}}

type ClientsImpl struct {
	connections *sys.Map[int64, Client]

	rooms map[string]*Room

	identity int64
}

func (c *ClientsImpl) Upgrade(w http.ResponseWriter, r *http.Request) (Client, error) {
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

func (c *ClientsImpl) Add(client Client) {
	c.connections.Put(client.GetId(), client)
}

func (c *ClientsImpl) Get(clientId int64) (Client, error) {
	client, ok := c.connections.Get(clientId)
	if !ok {
		return nil, errors.New("client not found")
	}
	return client, nil
}

func (c *ClientsImpl) Remove(clientId int64) error {
	client, err := c.Get(clientId)
	if err != nil {
		return err
	}
	c.connections.Remove(clientId)

	err = client.Close()
	return err
}

func (c *ClientsImpl) Join(clientId int64, roomName string) error {
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

func (c *ClientsImpl) Leave(clientId int64, roomName string) error {
	client, err := c.Get(clientId)
	if err != nil {
		return err
	}

	room := c.rooms[roomName]
	if room == nil {
		return errors.New("room not found")
	}

	room.Leave(client)

	return nil
}

func (c *ClientsImpl) Room(name string) *Room {
	room := c.rooms[name]
	if room == nil {
		room = NewRoom(name)
		c.rooms[name] = room
	}

	return room
}

func (c *ClientsImpl) Close(client Client) error {
	return client.Close()
}

func NewClients() Clients {
	rooms := make(map[string]*Room)

	rooms["subscribe"] = NewRoom("subscribe")

	return &ClientsImpl{
		connections: sys.NewMap[int64, Client](),
		identity:    0,
		rooms:       rooms,
	}
}
