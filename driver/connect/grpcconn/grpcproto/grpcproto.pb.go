// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpcproto.proto

package grpcproto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Req struct {
	Payload              []byte   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Req) Reset()         { *m = Req{} }
func (m *Req) String() string { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()    {}
func (*Req) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4c2577fa46b17ad, []int{0}
}

func (m *Req) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Req.Unmarshal(m, b)
}
func (m *Req) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Req.Marshal(b, m, deterministic)
}
func (m *Req) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Req.Merge(m, src)
}
func (m *Req) XXX_Size() int {
	return xxx_messageInfo_Req.Size(m)
}
func (m *Req) XXX_DiscardUnknown() {
	xxx_messageInfo_Req.DiscardUnknown(m)
}

var xxx_messageInfo_Req proto.InternalMessageInfo

func (m *Req) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Res struct {
	Payload              []byte   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Res) Reset()         { *m = Res{} }
func (m *Res) String() string { return proto.CompactTextString(m) }
func (*Res) ProtoMessage()    {}
func (*Res) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4c2577fa46b17ad, []int{1}
}

func (m *Res) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Res.Unmarshal(m, b)
}
func (m *Res) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Res.Marshal(b, m, deterministic)
}
func (m *Res) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Res.Merge(m, src)
}
func (m *Res) XXX_Size() int {
	return xxx_messageInfo_Res.Size(m)
}
func (m *Res) XXX_DiscardUnknown() {
	xxx_messageInfo_Res.DiscardUnknown(m)
}

var xxx_messageInfo_Res proto.InternalMessageInfo

func (m *Res) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*Req)(nil), "Req")
	proto.RegisterType((*Res)(nil), "Res")
}

func init() {
	proto.RegisterFile("grpcproto.proto", fileDescriptor_a4c2577fa46b17ad)
}

var fileDescriptor_a4c2577fa46b17ad = []byte{
	// 111 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x2f, 0x2a, 0x48,
	0x2e, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x03, 0x93, 0x4a, 0xf2, 0x5c, 0xcc, 0x41, 0xa9, 0x85, 0x42,
	0x12, 0x5c, 0xec, 0x05, 0x89, 0x95, 0x39, 0xf9, 0x89, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x3c,
	0x41, 0x30, 0x2e, 0x44, 0x41, 0x31, 0x6e, 0x05, 0x46, 0xaa, 0x5c, 0xec, 0xee, 0x45, 0x05, 0xc9,
	0x8e, 0x05, 0x99, 0x42, 0x52, 0x5c, 0x5c, 0x41, 0x05, 0xc9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5,
	0x25, 0x42, 0x2c, 0x7a, 0x41, 0xa9, 0x85, 0x52, 0x20, 0xb2, 0x58, 0x89, 0x21, 0x89, 0x0d, 0x6c,
	0x9f, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x2d, 0x35, 0x03, 0x82, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GrpcApiClient is the client API for GrpcApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GrpcApiClient interface {
	RpcRequest(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error)
}

type grpcApiClient struct {
	cc grpc.ClientConnInterface
}

func NewGrpcApiClient(cc grpc.ClientConnInterface) GrpcApiClient {
	return &grpcApiClient{cc}
}

func (c *grpcApiClient) RpcRequest(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/GrpcApi/RpcRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GrpcApiServer is the server API for GrpcApi service.
type GrpcApiServer interface {
	RpcRequest(context.Context, *Req) (*Res, error)
}

// UnimplementedGrpcApiServer can be embedded to have forward compatible implementations.
type UnimplementedGrpcApiServer struct {
}

func (*UnimplementedGrpcApiServer) RpcRequest(ctx context.Context, req *Req) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RpcRequest not implemented")
}

func RegisterGrpcApiServer(s *grpc.Server, srv GrpcApiServer) {
	s.RegisterService(&_GrpcApi_serviceDesc, srv)
}

func _GrpcApi_RpcRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcApiServer).RpcRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GrpcApi/RpcRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcApiServer).RpcRequest(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

var _GrpcApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "GrpcApi",
	HandlerType: (*GrpcApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RpcRequest",
			Handler:    _GrpcApi_RpcRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpcproto.proto",
}
