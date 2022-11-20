package websocket

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
	"ycore/dao"
	"ycore/driver/connect/websocket/socketclient"
	"ycore/driver/connect/websocket/socketserver"
	"ycore/module/mylog"
)

type ApiCallBack interface {
	OnNewConnect(token string)
	OnClose(token string)
	ReceiveMessage(ctx context.Context, socketClient *socketclient.Handler, message []byte)
}

type WebsocketManager struct {
	apiCallBack ApiCallBack
	server      *socketserver.Handle
	clientMap   dao.FastSyncMap
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

	self.server = socketserver.New()
	self.server.ImportSocketManager(self)

	s := &http.Server{
		Handler:      self.server,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	errc := make(chan error, 1)
	mylog.Infof("[Websocket][Server] at %v", addr)
	go func() {
		errc <- s.Serve(l)
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	select {
	case err := <-errc:
		mylog.Errorf("[Websocket][Server] failed to launch serve: %v", err)
	case f := <-sigs:
		mylog.Errorf("[Websocket][Server] server be terminating: %v", f)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return s.Shutdown(ctx)
}
