package web

import "github.com/gorilla/websocket"

var clients *Clients = NewClients()

type Clients struct {
	connections *map[*websocket.Conn]*ClientContext
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

func NewClients() *Clients {
	connections := make(map[*websocket.Conn]*ClientContext)
	return &Clients{
		connections: &connections,
	}
}

func GetClients() *Clients {
	return clients
}
