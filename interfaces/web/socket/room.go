package socket

import (
	"github.com/sirupsen/logrus"
	"sync"
)

type Room struct {
	name    string
	clients []Client
	mutex   sync.Mutex
	logger  *logrus.Entry
}

func (r *Room) Join(client Client) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.clients = append(r.clients, client)
}

func (r *Room) Leave(client *ClientImpl) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i, c := range r.clients {
		if c == client {
			r.clients = append(r.clients[:i], r.clients[i+1:]...)
			break
		}
	}
}

func (r *Room) Emit(message interface{}) {
	for _, c := range r.clients {
		_ = c.Emit(message)
	}
}

func (r *Room) GetClients() []Client {
	return r.clients
}

func NewRoom(name string) *Room {
	return &Room{
		name:   name,
		logger: logrus.WithFields(logrus.Fields{name: name}),
	}
}
