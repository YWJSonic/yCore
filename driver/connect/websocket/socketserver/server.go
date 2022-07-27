package socketserver

import (
	"net/http"
	"ycore/driver/connect/websocket/socketclient"
	"ycore/module/mylog"

	"nhooyr.io/websocket"
)

type SocketManagerCallBack interface {
	OnSocketConnect(socketClient *socketclient.Handler) error
}

type Handle struct {
	socketManagerCallBack SocketManagerCallBack
}

func New() *Handle {
	return &Handle{}
}

func (self *Handle) ImportSocketManager(socketManagerCallBack SocketManagerCallBack) {
	self.socketManagerCallBack = socketManagerCallBack
}

func (self *Handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		mylog.Errorf("[Websocket][ServeHTTP] %v", err)
		return
	}

	callBack := self.socketManagerCallBack.(socketclient.SocketManagerCallBack)
	socketClient := socketclient.New(r.Context(), c, callBack)

	// 開始監聽 Client 訊號後才會通知 api handle 層有新連線
	socketClient.Listen()
}
