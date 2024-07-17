package web

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
