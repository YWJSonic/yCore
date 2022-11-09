package grpcconn

import (
	"context"
	"fmt"
	"testing"
	"time"
	"ycore/driver/connect/grpcconn/grpcproto"

	"google.golang.org/grpc"
)

func mockGrpc() (*mockServer, *mockClient) {
	ms := &mockServer{}
	mc := &mockClient{}

	go func() {
		if err := NewGRPCServer(":8081", ms.launchServer); err != nil {
			fmt.Println(err)
		}
	}()

	time.Sleep(time.Second)

	if err := NewGRPCClient(":8081", mc.launchClient); err != nil {
		fmt.Println(err)
	}

	return ms, mc
}

func TestGrpc(t *testing.T) {
	_, client := mockGrpc()
	time.Sleep(time.Second)
	if res, err := client.RpcRequest(context.TODO(), &grpcproto.Req{
		Payload: []byte("1234"),
	}); err != nil {
		fmt.Println(err)

	} else {
		fmt.Println(res)

	}

}

type mockServer struct {
}

func (ms *mockServer) launchServer(grpcServer *grpc.Server) error {
	grpcproto.RegisterGrpcApiServer(grpcServer, ms)
	return nil
}
func (ms *mockServer) RpcRequest(ctx context.Context, req *grpcproto.Req) (*grpcproto.Res, error) {

	return &grpcproto.Res{
		Payload: req.Payload,
	}, nil
}

type mockClient struct {
	grpcproto.GrpcApiClient
}

func (mc *mockClient) launchClient(grpcClient *grpc.ClientConn) error {
	mc.GrpcApiClient = grpcproto.NewGrpcApiClient(grpcClient)
	return nil
}
