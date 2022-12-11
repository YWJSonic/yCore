package websocket

import (
	"github.com/YWJSonic/ycore/driver/connect/websocket/socketclient"
)

// Websocker Server 有新連線
// @params socketclient.Handler Websocket Client 物件
func (self *WebsocketManager) OnSocketConnect(socketClient *socketclient.Handler) {
	self.apiCallBack.OnNewConnect(socketClient)
}
