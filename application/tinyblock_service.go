package application

import (
	"github.com/chainpusher/chainpusher/interfaces/web/socket"
)

type TinyBlockService interface {
	Subscribe(clientId int64) (*socket.Client, error)
}
