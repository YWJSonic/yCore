package connect

import (
	"context"
	"log"
	"yangServer/dao"

	"github.com/gorilla/websocket"
)

type SocketManager struct {
}

func (self *SocketManager) Send() {

}

type Socket struct {
	ctx  context.Context
	conn *websocket.Conn
	read chan dao.MessageData
	send chan dao.MessageData
	// Read  func(messageType int, p []byte, err error) // 提供外部接收資料
	// Write func(messageType int, data []byte) error   // 提供外部發送資料
}

func newSocket(ctx context.Context, conn *websocket.Conn) *Socket {
	return &Socket{
		ctx:  ctx,
		conn: conn,
	}
}

func (self *Socket) readMessage() {
	for {
		_, message, err := self.conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", message)
	}
}

func (self *Socket) writeMessage() {
	for {
		select {
		case messageData := <-self.send:
			err := self.conn.WriteMessage(websocket.TextMessage, messageData.Payload)
			if err != nil {
				log.Println("write fail:", err)
				return
			}
		}
	}
}
