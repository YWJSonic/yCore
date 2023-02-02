package socketclient

import (
	"context"
	"log"
	"testing"

	"nhooyr.io/websocket"
)

func TestClient(t *testing.T) {
	addr := "wss://127.0.0.1:5506"

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	conn, _, err := websocket.Dial(ctx, addr, nil)
	if err != nil {
		log.Fatalf("[Websocket][Launch] Listen Error addr: %v", addr)
	}

	mockTest := &mockTest{}
	socket := New(ctx, conn, mockTest)
	socket.Listen()
}

type mockTest struct{}

func (mockTest) OnClose(token string)                                                      {}
func (mockTest) ReceiveMessage(ctx context.Context, socketClient *Handler, message []byte) {}
