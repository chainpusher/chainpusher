package web

import "github.com/gorilla/websocket"

type ClientContext struct {
	connection *websocket.Conn
}

func (ctx *ClientContext) GetConnection() *websocket.Conn {
	return ctx.connection
}
