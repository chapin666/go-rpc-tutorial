// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello_http.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	hello_http.proto

It has these top-level messages:
	HelloHttpRequest
	HelloHttpReply
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "git.vodjk.com/go-grpc/example/proto/google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type HelloHttpRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloHttpRequest) Reset()                    { *m = HelloHttpRequest{} }
func (m *HelloHttpRequest) String() string            { return proto1.CompactTextString(m) }
func (*HelloHttpRequest) ProtoMessage()               {}
func (*HelloHttpRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloHttpRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloHttpReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloHttpReply) Reset()                    { *m = HelloHttpReply{} }
func (m *HelloHttpReply) String() string            { return proto1.CompactTextString(m) }
func (*HelloHttpReply) ProtoMessage()               {}
func (*HelloHttpReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloHttpReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto1.RegisterType((*HelloHttpRequest)(nil), "proto.HelloHttpRequest")
	proto1.RegisterType((*HelloHttpReply)(nil), "proto.HelloHttpReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for HelloHttp service

type HelloHttpClient interface {
	SayHello(ctx context.Context, in *HelloHttpRequest, opts ...grpc.CallOption) (*HelloHttpReply, error)
}

type helloHttpClient struct {
	cc *grpc.ClientConn
}

func NewHelloHttpClient(cc *grpc.ClientConn) HelloHttpClient {
	return &helloHttpClient{cc}
}

func (c *helloHttpClient) SayHello(ctx context.Context, in *HelloHttpRequest, opts ...grpc.CallOption) (*HelloHttpReply, error) {
	out := new(HelloHttpReply)
	err := grpc.Invoke(ctx, "/proto.HelloHttp/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HelloHttp service

type HelloHttpServer interface {
	SayHello(context.Context, *HelloHttpRequest) (*HelloHttpReply, error)
}

func RegisterHelloHttpServer(s *grpc.Server, srv HelloHttpServer) {
	s.RegisterService(&_HelloHttp_serviceDesc, srv)
}

func _HelloHttp_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloHttpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloHttpServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HelloHttp/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloHttpServer).SayHello(ctx, req.(*HelloHttpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloHttp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.HelloHttp",
	HandlerType: (*HelloHttpServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloHttp_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello_http.proto",
}

func init() { proto1.RegisterFile("hello_http.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x48, 0xcd, 0xc9,
	0xc9, 0x8f, 0xcf, 0x28, 0x29, 0x29, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53,
	0x52, 0x32, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x89, 0x05, 0x99, 0xfa, 0x89, 0x79, 0x79,
	0xf9, 0x25, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0x10, 0x45, 0x4a, 0x6a, 0x5c, 0x02, 0x1e, 0x20,
	0x8d, 0x1e, 0x25, 0x25, 0x05, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c,
	0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92, 0x16, 0x17,
	0x1f, 0x92, 0xba, 0x82, 0x9c, 0x4a, 0x21, 0x09, 0x2e, 0xf6, 0xdc, 0xd4, 0xe2, 0xe2, 0xc4, 0x74,
	0x98, 0x42, 0x18, 0xd7, 0x28, 0x91, 0x8b, 0x13, 0xae, 0x56, 0x28, 0x84, 0x8b, 0x23, 0x38, 0xb1,
	0x12, 0xcc, 0x17, 0x12, 0x87, 0x58, 0xaa, 0x87, 0x6e, 0xa3, 0x94, 0x28, 0xa6, 0x44, 0x41, 0x4e,
	0xa5, 0x92, 0x44, 0xd3, 0xe5, 0x27, 0x93, 0x99, 0x84, 0x94, 0x78, 0xf5, 0x53, 0x2b, 0x12, 0x73,
	0x0b, 0x72, 0x52, 0xf5, 0x53, 0x93, 0x33, 0xf2, 0xad, 0x18, 0xb5, 0x92, 0xd8, 0xc0, 0xea, 0x8d,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x33, 0xe5, 0x25, 0x7c, 0xf6, 0x00, 0x00, 0x00,
}
