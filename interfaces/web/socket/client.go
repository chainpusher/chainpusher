package socket

type Client interface {
	GetId() int64

	Emit(message interface{}) error

	Read() chan []byte

	Close() error
}
