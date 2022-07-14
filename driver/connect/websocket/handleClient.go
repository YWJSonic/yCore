package websocket

import (
	"context"
	"yangServer/driver/connect/websocket/socketclient"
)

func (self *WebsocketManager) ReceiveMessage(ctx context.Context, socketClient *socketclient.Handler, message []byte) {
	self.apiCallBack.ReceiveMessage(ctx, socketClient, message)
}
