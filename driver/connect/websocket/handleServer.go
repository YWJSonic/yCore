package websocket

import (
	"ycore/driver/connect/websocket/socketclient"
)

// Websocker Server 有新連線
// @params socketclient.Handler Websocket Client 物件
func (self *WebsocketManager) OnSocketConnect(socketClient *socketclient.Handler) error {
	return nil
}
