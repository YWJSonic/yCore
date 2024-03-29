package websocket

import (
	"github.com/YWJSonic/ycore/constant"
	"github.com/YWJSonic/ycore/driver/connect/websocket/socketclient"
)

// 通知 Websocket Manager 有連線被中斷
// 斷線訊號轉拋至 GsApi
func (self *WebsocketManager) OnClose(token string) {
	self.clientMap.Delete(token)
	self.apiCallBack.OnClose(token)
}

func (self *WebsocketManager) StoryClient(socketClient socketclient.IHandle) {
	self.clientMap.Store(socketClient.GetToken(), socketClient)
}

func (self *WebsocketManager) GetClient(token string) socketclient.IHandle {

	if client, ok := self.clientMap.Load(token); ok {
		return client.(socketclient.IHandle)
	} else {
		return nil
	}
}

func (self *WebsocketManager) GetBalanceClient() socketclient.IHandle {

	// 取得最低權重 client
	var weight int64 = int64(constant.MaxInt)
	var client *socketclient.Handler
	self.clientMap.Range(func(key string, value interface{}) bool {
		if value.(*socketclient.Handler).GetWeight() < weight {
			weight = value.(*socketclient.Handler).GetWeight()
			client = value.(*socketclient.Handler)
		}
		return true
	})

	return client
}

func (self *WebsocketManager) GetAllClient() map[string]socketclient.IHandle {
	clients := map[string]socketclient.IHandle{}
	self.clientMap.Range(func(key string, value interface{}) bool {
		clients[key] = value.(*socketclient.Handler)
		return true
	})

	return clients
}
