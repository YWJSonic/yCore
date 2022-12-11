package socketserver

import (
	"net/http"

	"github.com/YWJSonic/ycore/driver/connect/websocket/socketclient"
	"github.com/YWJSonic/ycore/module/mylog"

	"github.com/rs/xid"
	"nhooyr.io/websocket"
)

type SocketManagerCallBack interface {
	OnSocketConnect(socketClient *socketclient.Handler)
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
	socketClient.SetToken(xid.New().String())
	if self.socketManagerCallBack != nil {
		self.socketManagerCallBack.OnSocketConnect(socketClient)
	}
	// 開始監聽 Client 訊號後才會通知 api handle 層有新連線
	_ = socketClient.Listen()
}
