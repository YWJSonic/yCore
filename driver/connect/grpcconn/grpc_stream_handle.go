package grpcconn

import "github.com/YWJSonic/ycore/driver/connect/grpcconn/streamproto"

type onRead func(string, *streamproto.Message)

type ioHandle interface {
	Write(msg *streamproto.Message) error
	Read(token string) error
}

type iStram interface {
	Send(*streamproto.Message) error
	Recv() (*streamproto.Message, error)
}

type stramHandle struct {
	stream iStram
	onRead
}

func newStramHandle(stream iStram, callback onRead) ioHandle {
	return &stramHandle{
		stream: stream,
		onRead: callback,
	}
}

func (handle *stramHandle) Write(msg *streamproto.Message) error {
	return handle.stream.Send(msg)
}

func (handle *stramHandle) Read(token string) error {
	for {
		msg, err := handle.stream.Recv()
		if err != nil {
			return err
		}
		handle.onRead(token, msg)
	}
}
