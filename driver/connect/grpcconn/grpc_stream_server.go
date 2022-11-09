package grpcconn

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"
	"ycore/driver/connect/grpcconn/streamproto"

	"github.com/rs/xid"
	"google.golang.org/grpc"
)

type ServerOnReadCallback interface {
	OnNotice(token string, msg []byte)
	OnRequest(token string, msg []byte) ([]byte, error)
}

type IGrpcStreamServer interface {
	Launch() error
	GetAllToken() []string
	Request(token string, msg []byte) ([]byte, error)
	Notice(token string, msg []byte) error
}

type streamServer struct {
	clientSync  sync.RWMutex
	clientIOMap map[string]ioHandle
	callback    ServerOnReadCallback

	resChanMapSync sync.RWMutex
	resChanMap     map[string]chan *streamproto.Message
}

func NewGrpcStreamServer(callback ServerOnReadCallback) IGrpcStreamServer {
	return &streamServer{
		clientIOMap: make(map[string]ioHandle),
		resChanMap:  make(map[string]chan *streamproto.Message),
		callback:    callback,
	}
}

func (server *streamServer) OnMessage(stream streamproto.Stream_OnMessageServer) error {
	// TODO: auth login conn 第一筆訊息可定為驗證資訊
	// msg, err := stream.Recv()
	// if err != nil {
	// 	return err
	// }

	handle := newStramHandle(stream, server.onRead)

	token := xid.New().String()

	server.clientSync.Lock()
	server.clientIOMap[token] = handle
	server.clientSync.Unlock()

	server.resChanMapSync.Lock()
	server.resChanMap[token] = make(chan *streamproto.Message)
	server.resChanMapSync.Unlock()

	err := handle.Read(token)

	server.clientSync.Lock()
	delete(server.clientIOMap, token)
	server.clientSync.Unlock()

	server.resChanMapSync.Lock()
	delete(server.resChanMap, token)
	server.resChanMapSync.Unlock()

	return err
}

func (server *streamServer) Launch() error {
	lis, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	streamproto.RegisterStreamServer(grpcServer, server)
	return grpcServer.Serve(lis)
}

func (server *streamServer) GetAllToken() []string {
	tokens := []string{}
	server.clientSync.Lock()
	defer server.clientSync.Unlock()
	for token := range server.clientIOMap {
		tokens = append(tokens, token)
	}
	return tokens
}

func (server *streamServer) Request(token string, msg []byte) ([]byte, error) {

	server.clientSync.Lock()
	handle, ok := server.clientIOMap[token]
	server.clientSync.Unlock()
	if !ok {
		return nil, fmt.Errorf("client not exist token: %v", token)
	}

	server.resChanMapSync.Lock()
	reschan, ok := server.resChanMap[token]
	server.resChanMapSync.Unlock()
	if !ok {
		return nil, fmt.Errorf("client res chan not exist token: %v", token)
	}

	err := handle.Write(&streamproto.Message{
		OperationCode: streamproto.Operation_Request,
		Payload:       msg,
	})
	if err != nil {
		return nil, err
	}

	reqCtx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancel()

	select {
	case <-reqCtx.Done():
		return nil, errors.New("req time out")

	case resMsg := <-reschan:
		if resMsg.OperationCode == streamproto.Operation_Error {
			return nil, errors.New(string(resMsg.Payload))
		}

		return resMsg.Payload, nil
	}
}

func (server *streamServer) Notice(token string, msg []byte) error {
	server.clientSync.Lock()
	handle, ok := server.clientIOMap[token]
	server.clientSync.Unlock()
	if !ok {
		return fmt.Errorf("client not exist token: %v", token)
	}

	return handle.Write(&streamproto.Message{
		OperationCode: streamproto.Operation_Notice,
		Payload:       msg,
	})
}

func (server *streamServer) onRead(token string, msg *streamproto.Message) {
	switch msg.OperationCode {
	case streamproto.Operation_Request:
		go server.onRequest(token, msg.Payload)

	case streamproto.Operation_Response, streamproto.Operation_Error:
		server.resChanMapSync.Lock()
		defer server.resChanMapSync.Unlock()

		reschan, ok := server.resChanMap[token]
		if !ok {
			fmt.Printf("client res chan not exist token: %v", token)
			return
		}

		reschan <- msg

	case streamproto.Operation_Notice:
		server.callback.OnNotice(token, msg.Payload)

	default:
		fmt.Printf("msg opcode error token: %v", token)
	}
}

func (server *streamServer) onRequest(token string, msg []byte) {
	server.clientSync.Lock()
	handle, ok := server.clientIOMap[token]
	server.clientSync.Unlock()
	if !ok {
		err := errors.New("server no client connect ref")
		if err = handle.Write(&streamproto.Message{
			OperationCode: streamproto.Operation_Error,
			Payload:       []byte(err.Error()),
		}); err != nil {
			fmt.Printf("client onRequest write err: %v", err)
		}

		return
	}

	res, err := server.callback.OnRequest(token, msg)
	if err != nil {
		if err = handle.Write(&streamproto.Message{
			OperationCode: streamproto.Operation_Error,
			Payload:       []byte(err.Error()),
		}); err != nil {
			fmt.Printf("client onRequest write err: %v", err)
		}
	} else {
		if err = handle.Write(&streamproto.Message{
			OperationCode: streamproto.Operation_Response,
			Payload:       res,
		}); err != nil {
			fmt.Printf("client onRequest write err: %v", err)
		}
	}
}
