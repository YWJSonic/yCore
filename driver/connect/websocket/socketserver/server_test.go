package socketserver

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"testing"
	"time"

	"github.com/YWJSonic/ycore/driver/connect/websocket/socketclient"
	"github.com/YWJSonic/ycore/module/mylog"
)

func TestServer(t *testing.T) {
	addr := ":5506"
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("[Websocket][Launch] Listen Error addr: %v", addr)
	}

	server := New()
	mockTest := &mockTest{}
	server.ImportSocketManager(mockTest)

	s := &http.Server{
		Handler:      server,
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
}

type mockTest struct{}

func (mockTest) OnSocketConnect(socketClient *socketclient.Handler) {
	fmt.Println("New Client:", socketClient.GetToken())
}
func (mockTest) OnClose(token string) {
	fmt.Println("----Server OnClose----")
}
func (mockTest) ReceiveMessage(ctx context.Context, socketClient *socketclient.Handler, message []byte) {
}
