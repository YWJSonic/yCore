package grpcconn

import (
	"net"

	"google.golang.org/grpc"
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

	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil
}

type CreateGrpcClientFun func(grpcClient *grpc.ClientConn) error

func NewGRPCClient(adderss string, newGrpcFun CreateGrpcClientFun) error {
	conn, err := grpc.Dial(adderss, grpc.WithInsecure())
	if err != nil {
		return err
	}

	if err := newGrpcFun(conn); err != nil {
		return err
	}
	return nil
}
