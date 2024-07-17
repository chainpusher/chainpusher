package web

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clients *Clients = NewClients()

type Clients struct {
	connections *map[*websocket.Conn]*ClientContext

	room *Room
}

func (c *Clients) Upgrade(w http.ResponseWriter, r *http.Request) (*Client, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	return NewClient(conn), err
}

func (c *Clients) Add(conn *websocket.Conn) *ClientContext {
	ctx := &ClientContext{
		connection: conn,
	}
	(*c.connections)[conn] = ctx
	return ctx
}

func (c *Clients) Get(conn *websocket.Conn) *ClientContext {
	return (*c.connections)[conn]
}

func (c *Clients) Remove(conn *websocket.Conn) {
	delete(*c.connections, conn)
	err := conn.Close()
	if err != nil {
		return
	}
}

func (c *Clients) SendAll(message interface{}) error {
	for conn := range *c.connections {
		_ = conn.WriteJSON(message)
	}

	return nil
}

func (c *Clients) Close(client *Client) error {
	err := client.connection.Close()
	return err
}

func NewClients() *Clients {
	connections := make(map[*websocket.Conn]*ClientContext)
	return &Clients{
		connections: &connections,
	}
}

func GetClients() *Clients {
	return clients
}
