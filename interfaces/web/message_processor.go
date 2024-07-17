package web

import (
	"encoding/json"
	"github.com/chainpusher/chainpusher/interfaces/facade/dto"
	"github.com/sirupsen/logrus"
)

type MessageProcessor interface {
	Process(client *Client, message []byte)
}

type CallbackMessageProcessor struct {
	callback func(client *Client, message []byte)
}

func (c *CallbackMessageProcessor) Process(client *Client, message []byte) {
	c.callback(client, message)
}

func NewCallbackMessageProcessor(callback func(client *Client, message []byte)) *CallbackMessageProcessor {
	return &CallbackMessageProcessor{
		callback: callback,
	}
}

type JsonRpcMessageProcessor struct {
}

func (p *JsonRpcMessageProcessor) Process(client *Client, message []byte) {
	var call *dto.JsonRpcDto
	err := json.Unmarshal(message, &call)
	if err != nil {
		logrus.Warnf("Error unmarshalling json rpc message: %v", err)
		return
	}

}

func NewJsonRpcMessageProcess() MessageProcessor {
	return &JsonRpcMessageProcessor{}
}
