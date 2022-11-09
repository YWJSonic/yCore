package grpcconn

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CreateGrpcServerFun func(grpcServer *grpc.Server) error

func NewGRPCServer(address string, newGrpcFun CreateGrpcServerFun, opt ...grpc.ServerOption) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	s := grpc.NewServer(opt...)
	if err := newGrpcFun(s); err != nil {
		return err
	}

	return s.Serve(lis)
}

type CreateGrpcClientFun func(grpcClient *grpc.ClientConn) error

func NewGRPCClient(adderss string, newGrpcFun CreateGrpcClientFun) error {
	conn, err := grpc.Dial(adderss, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	return newGrpcFun(conn)
}
