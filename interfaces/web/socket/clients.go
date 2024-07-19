package socket

import "net/http"

type Clients interface {
	Upgrade(w http.ResponseWriter, r *http.Request) (Client, error)

	Add(client Client)

	Close(client Client) error

	Get(clientId int64) (Client, error)

	Room(name string) *Room

	Join(clientId int64, roomName string) error
}
