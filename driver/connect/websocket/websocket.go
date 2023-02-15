package websocket

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/YWJSonic/ycore/dao"
	"github.com/YWJSonic/ycore/driver/connect/websocket/socketclient"
	"github.com/YWJSonic/ycore/driver/connect/websocket/socketserver"
	"github.com/YWJSonic/ycore/module/mylog"
)

type ApiCallBack interface {
	OnNewConnect(socketClient socketclient.IHandle)
	OnClose(token string)
	ReceiveMessage(ctx context.Context, socketClient socketclient.IHandle, message []byte)
}

type WebsocketManager struct {
	apiCallBack  ApiCallBack
	serverHandle *socketserver.Handle
	server       *http.Server
	clientMap    dao.FastSyncMap
}

func New() *WebsocketManager {
	return &WebsocketManager{
		clientMap: dao.FastSyncMap{},
	}
}

func (self *WebsocketManager) ImportApiCallBack(apiCallBack ApiCallBack) {
	self.apiCallBack = apiCallBack
}

func (self *WebsocketManager) Launch(addr string) error {

	if len(addr) == 0 {
		return fmt.Errorf("[Websocket][Launch] addr Error addr: %v", addr)
	}

	l, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("[Websocket][Launch] Listen Error addr: %v", addr)
	}

	self.serverHandle = socketserver.New()
	self.serverHandle.ImportSocketManager(self)

	self.server = &http.Server{
		Handler:      self.serverHandle,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	mylog.Infof("[Websocket][Server] at %v", addr)
	return self.server.Serve(l)
}

func (self *WebsocketManager) Stop() error {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
	defer cancel()
	return self.server.Shutdown(ctx)
}
