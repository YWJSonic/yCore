package example

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/YWJSonic/ycore/driver/connect/grpcconn"
	"github.com/YWJSonic/ycore/module/mylog"
	echo "github.com/YWJSonic/ycore/proto"
	"github.com/YWJSonic/ycore/util"

	"google.golang.org/grpc"
)

var port = "127.0.0.1:8080"

func TestDo(t *testing.T) {
	go GoGrpcServer()
	GoGrpcClient()
}

/// Server
func GoGrpcServer() {
	newGrpcServerFun := func(grpcServer *grpc.Server) error {
		ServerProto := &Echo{}
		echo.RegisterEchoEventServer(grpcServer, ServerProto)
		return nil
	}

	_ = grpcconn.NewGRPCServer(port, newGrpcServerFun)
}

type Echo struct{}

func (self *Echo) Echo(ctx context.Context, in *echo.EchoRequest) (*echo.EchoReply, error) {
	ec := util.Sprint(time.Now())
	mylog.Info(ec)
	return &echo.EchoReply{Message: ec}, nil
}

/// Client
func GoGrpcClient() {
	var ecc echo.EchoEventClient
	newGrpcClientFun := func(grpcClient *grpc.ClientConn) error {
		ecc = echo.NewEchoEventClient(grpcClient)
		return nil
	}
	_ = grpcconn.NewGRPCClient(port, newGrpcClientFun)

	r, err := ecc.Echo(context.Background(), &echo.EchoRequest{Message: "123"})
	if err != nil {
		log.Fatalf("無法執行 Plus 函式：%v", err)
	}
	log.Printf("回傳結果：%s , 時間:%d", r.Message, r.Unixtime)
}
