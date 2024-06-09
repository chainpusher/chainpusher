package monitor

import (
	"encoding/json"
	"os"

	"github.com/sirupsen/logrus"
)

var NewLineByte []byte = []byte("\n")

type BlockWatcher interface {
	GetChannel() chan interface{}

	Start()
}

type BlcokLoggingWatcher struct {
	Descriptor *os.File

	Channel chan interface{}
}

func (b *BlcokLoggingWatcher) Start() {
	go func() {
		for {
			select {
			case block, ok := <-b.GetChannel():
				if !ok {
					logrus.Warn("Channel closed")
					return
				}

				serialized, err := json.Marshal(block)
				if err != nil {
					logrus.Errorf("Error marshalling block: %v", err)
				}

				b.Descriptor.Write(serialized)
				b.Descriptor.Write(NewLineByte)
			}
		}
	}()

}

func (b *BlcokLoggingWatcher) GetChannel() chan interface{} {
	return b.Channel
}

func NewBlockLoggingWatcher(channel chan interface{}) BlockWatcher {

	fd, err := os.OpenFile("block.json", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logrus.Errorf("Error opening file: %v", err)
	}

	return &BlcokLoggingWatcher{
		Descriptor: fd,
		Channel:    channel,
	}
}
