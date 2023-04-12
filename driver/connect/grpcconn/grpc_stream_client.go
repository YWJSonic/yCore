package grpcconn

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/YWJSonic/ycore/driver/connect/grpcconn/streamproto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientOnReadCallback interface {
	OnNotice(msg []byte)
	OnRequest(msg []byte) ([]byte, error)
}

type IGrpcStreamClient interface {
	Launch() error
	Request(msg []byte) ([]byte, error)
	Notice(msg []byte) error
}

type streamClient struct {
	serverIO ioHandle
	resChan  chan *streamproto.Message
	callback ClientOnReadCallback
}

func NewGrpcStreamClient(callback ClientOnReadCallback) IGrpcStreamClient {
	return &streamClient{
		callback: callback,
		resChan:  make(chan *streamproto.Message),
	}
}

func (client *streamClient) Launch() error {
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()

	stream, err := streamproto.NewStreamClient(conn).OnMessage(context.TODO())
	if err != nil {
		return err
	}
	client.serverIO = newStramHandle(stream, client.onRead)
	return client.serverIO.Read("")
}

func (client *streamClient) Request(msg []byte) ([]byte, error) {
	err := client.serverIO.Write(&streamproto.Message{
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

	case resMsg := <-client.resChan:
		if resMsg.OperationCode == streamproto.Operation_Error {
			return nil, errors.New(string(resMsg.Payload))
		}

		return resMsg.Payload, nil
	}
}

func (client *streamClient) Notice(msg []byte) error {
	return client.serverIO.Write(&streamproto.Message{
		OperationCode: streamproto.Operation_Notice,
		Payload:       msg,
	})
}

func (client *streamClient) onRead(token string, msg *streamproto.Message) {
	switch msg.OperationCode {
	case streamproto.Operation_Request:
		go client.onRequest(msg.Payload)

	case streamproto.Operation_Response, streamproto.Operation_Error:
		client.resChan <- msg

	case streamproto.Operation_Notice:
		client.callback.OnNotice(msg.Payload)

	default:
		fmt.Printf("client onRead token: %v, msg:%v\n", token, msg)
	}
}

func (client *streamClient) onRequest(msg []byte) {
	res, err := client.callback.OnRequest(msg)
	if err != nil {
		if err = client.serverIO.Write(&streamproto.Message{
			OperationCode: streamproto.Operation_Error,
			Payload:       []byte(err.Error()),
		}); err != nil {
			fmt.Printf("client onRequest write err: %v", err)
		}
	} else {
		if err = client.serverIO.Write(&streamproto.Message{
			OperationCode: streamproto.Operation_Response,
			Payload:       res,
		}); err != nil {
			fmt.Printf("client onRequest write err: %v", err)
		}
	}
}
