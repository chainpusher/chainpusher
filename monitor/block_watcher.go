package monitor

import (
	"encoding/json"
	"os"
	"path"

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
	go b.Forever()
}

func (b *BlcokLoggingWatcher) Forever() {
	for {
		select {
		case block, ok := <-b.GetChannel():
			if !ok {
				logrus.Warn("The channel had been closed")
				return
			}

			b.WriteBlock(block)
		}
	}
}

func (b *BlcokLoggingWatcher) WriteBlock(block interface{}) {
	serialized, err := json.Marshal(block)
	if err != nil {
		logrus.Errorf("Error marshalling block: %v", err)
		return
	}

	logrus.Debugf("Write block to file, that size is %d", len(serialized))

	serializationWritten, err := b.Descriptor.Write(serialized)
	if err != nil {
		logrus.Errorf("Error writing block: %v", err)
		logrus.Debugf("Written %d bytes, data is %s", serializationWritten, serialized)
		return
	}

	b.Descriptor.Write(NewLineByte)
}

func (b *BlcokLoggingWatcher) Close() {
	b.Descriptor.Close()
}

func (b *BlcokLoggingWatcher) GetChannel() chan interface{} {
	return b.Channel
}

func NewBlockLoggingWatcher(channel chan interface{}, rawFilePath string) BlockWatcher {

	if len(rawFilePath) == 0 {
		logrus.Debug("Block logging file path is empty")
		return nil
	}

	// is absolute path
	if !path.IsAbs(rawFilePath) {
		wd, err := os.Getwd()
		if err != nil {
			logrus.Errorf("Error getting working directory: %v", err)
			return nil
		}
		rawFilePath = path.Join(wd, rawFilePath)
	}
	logrus.Debugf("Block logging file path: %s", rawFilePath)

	fd, err := os.OpenFile(rawFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logrus.Errorf("Error opening file: %v", err)
	}

	return &BlcokLoggingWatcher{
		Descriptor: fd,
		Channel:    channel,
	}
}
