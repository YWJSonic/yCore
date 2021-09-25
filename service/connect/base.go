package connect

import (
	"context"
	"log"
	"net/url"
	"yangServer/constant"

	"github.com/gorilla/websocket"
)

type SocketClientSetting struct {
	Ctx  context.Context
	Addr constant.Addr
	Path constant.Path
}

func NewSocket(setting SocketClientSetting) {
	u := url.URL{Scheme: "ws", Host: setting.Addr, Path: setting.Path}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	scoket := newSocket(setting.Ctx, c)

	go scoket.readMessage()
	go scoket.writeMessage()

}

type SocketServerSetting struct {
	Ctx  context.Context
	Addr constant.Addr
	Path constant.Path
}
