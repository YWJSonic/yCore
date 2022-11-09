package grpcconn

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type IStreamServer interface {
}

type IStreamClient interface {
}

type StreamServer struct {
}

func NewGrpcStreamServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 3000))
	if err != nil {
		log.Fatalf("failed to listed: %v", err)
	}

	// STEP 2-2：使用 gRPC 的 NewServer 方法來建立 gRPC Server 的實例
	grpcServer := grpc.NewServer()

	// STEP 2-3：在 gRPC Server 中註冊 service 的實作
	// 使用 proto 提供的 RegisterRouteGuideServer 方法，並將 routeGuideServer 作為參數傳入
	
	
	// STEP 2-4：啟動 grpcServer，並阻塞在這裡直到該程序被 kill 或 stop
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type StreamClient struct {
}

func NewGrpcStreamClient() {

}

type stramHandle struct {
}
