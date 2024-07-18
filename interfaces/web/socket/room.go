package socket

import (
	"sync"
)

type Room struct {
	name    string
	clients []*Client
	mutex   sync.Mutex
}

func (g *Room) Join(client *Client) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	g.clients = append(g.clients, client)
}

func (g *Room) Leave(client *Client) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	for i, c := range g.clients {
		if c == client {
			g.clients = append(g.clients[:i], g.clients[i+1:]...)
			break
		}
	}
}

func (g *Room) Emit(message interface{}) {
	for _, c := range g.clients {
		c.Emit(message)
	}
}

func NewRoom(name string) *Room {
	return &Room{
		name: name,
	}
}
