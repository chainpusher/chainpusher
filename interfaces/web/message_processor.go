package web

import (
	"encoding/json"
	"github.com/chainpusher/chainpusher/interfaces/facade"
	"github.com/chainpusher/chainpusher/interfaces/facade/dto"
	"github.com/chainpusher/chainpusher/interfaces/web/socket"
	"github.com/sirupsen/logrus"
)

type MessageProcessor interface {
	Process(client socket.Client, message []byte)
}

type CallbackMessageProcessor struct {
	callback func(client socket.Client, message []byte)
}

func (c *CallbackMessageProcessor) Process(client socket.Client, message []byte) {
	c.callback(client, message)
}

func NewCallbackMessageProcessor(callback func(client socket.Client, message []byte)) *CallbackMessageProcessor {
	return &CallbackMessageProcessor{
		callback: callback,
	}
}

type JsonRpcMessageProcessor struct {
	facade facade.TinyBlockServiceFacade
}

func (p *JsonRpcMessageProcessor) Process(client socket.Client, message []byte) {
	var call *dto.JsonRpcDto
	err := json.Unmarshal(message, &call)
	if err != nil {
		logrus.Warnf("Error unmarshalling json rpc message: %v", err)
		return
	}

	if call.Method == "subscribe" {
		p.facade.Subscribe(client.GetId())
	}
}

func NewJsonRpcMessageProcess(aFacade facade.TinyBlockServiceFacade) MessageProcessor {
	return &JsonRpcMessageProcessor{facade: aFacade}
}
