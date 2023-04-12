package grpcconn

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/YWJSonic/ycore/driver/connect/grpcconn/streamproto"
)

func mockStreamServerClient() (IGrpcStreamServer, IGrpcStreamClient) {

	ms := &mockServerReceiver{}
	mc := &mockClientReceiver{}

	token := "mock"
	mockServerIO := &mockioHandle{
		token: token,
	}
	server := &streamServer{
		clientIOMap: map[string]ioHandle{
			token: mockServerIO,
		},
		resChanMap: map[string]chan *streamproto.Message{
			token: make(chan *streamproto.Message),
		},
		callback: ms,
	}
	mockServerIO.callback = server.onRead

	mockClientIO := &mockioHandle{
		token: "",
	}
	client := &streamClient{
		serverIO: mockClientIO,
		callback: mc,
		resChan:  make(chan *streamproto.Message),
	}
	mockClientIO.callback = client.onRead
	return server, client
}

func TestGrpcStream(t *testing.T) {
	server, client := mockStreamServerClient()

	res, err := client.Request([]byte("client send request to servr"))
	fmt.Println(string(res), err)

	err = client.Notice([]byte("client send notice to servr"))
	fmt.Println(err)

	tokens := server.GetAllToken()
	for _, token := range tokens {
		res, err = server.Request(token, []byte("server send request to client"))
		fmt.Println(string(res), err)

		err = server.Notice(token, []byte("servr send notice to client"))
		fmt.Println(err)
	}
	time.Sleep(time.Second)
}

func TestGrpcConcurrency(t *testing.T) {
	_, client := mockStreamServerClient()

	wg := sync.WaitGroup{}
	count := 20000000
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(i int) {
			if err := client.Notice([]byte(strconv.Itoa(i))); err != nil {
				fmt.Println(err)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()

	count = 20000000
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(i int) {
			if _, err := client.Request([]byte(strconv.Itoa(i))); err != nil {
				fmt.Println(err)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
}

// BenchmarkClientNotice-8   	28089295	        44.48 ns/op	      64 B/op	       1 allocs/op
func BenchmarkClientNotice(b *testing.B) {
	_, client := mockStreamServerClient()

	mockPayload := []byte("123456789012345678901234567890")
	for i := 0; i < b.N; i++ {
		client.Notice(mockPayload)
	}
}

// BenchmarkClientReq-8   	  621322	      1736 ns/op	     496 B/op	       9 allocs/op
func BenchmarkClientReq(b *testing.B) {
	_, client := mockStreamServerClient()

	mockPayload := []byte("123456789012345678901234567890")
	for i := 0; i < b.N; i++ {
		client.Request(mockPayload)
	}
}

type mockServerReceiver struct {
}

func (ms *mockServerReceiver) OnNotice(token string, msg []byte) {
	// fmt.Println("server OnNotice", token, string(msg))
}
func (ms *mockServerReceiver) OnRequest(token string, msg []byte) ([]byte, error) {
	// fmt.Println("server OnRequest", token, string(msg))

	return []byte("to client res ok"), nil
}

type mockClientReceiver struct {
}

func (mc *mockClientReceiver) OnNotice(msg []byte) {
	// fmt.Println("client OnNotice", string(msg))
}

func (mc *mockClientReceiver) OnRequest(msg []byte) ([]byte, error) {
	// fmt.Println("client OnRequest", string(msg))

	return []byte("to server res ok"), nil
}

type mockioHandle struct {
	callback onRead
	token    string
}

func (mock *mockioHandle) Write(msg *streamproto.Message) error {
	mock.callback(mock.token, msg)
	return nil
}

func (mock *mockioHandle) Read(token string) error {
	return nil
}
