// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.example

/*
Package example is a generated protocol buffer package.

It is generated from these files:
	example.example

It has these top-level messages:
	H3Req
	H3Resp
*/
package example

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the example package it is being compiled against.
// A compilation error at this line likely means your copy of the
// example package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the example package

type H3Req struct {
	MsgType string            `protobuf:"bytes,1,opt,name=msgType" json:"msgType,omitempty"`
	Headers map[string]string `protobuf:"bytes,2,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body    []byte            `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Fin     bool              `protobuf:"varint,4,opt,name=fin" json:"fin,omitempty"`
}

func (m *H3Req) Reset()                    { *m = H3Req{} }
func (m *H3Req) String() string            { return proto.CompactTextString(m) }
func (*H3Req) ProtoMessage()               {}
func (*H3Req) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *H3Req) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *H3Req) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *H3Req) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *H3Req) GetFin() bool {
	if m != nil {
		return m.Fin
	}
	return false
}

type H3Resp struct {
	MsgType string            `protobuf:"bytes,1,opt,name=msgType" json:"msgType,omitempty"`
	Headers map[string]string `protobuf:"bytes,2,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body    []byte            `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Fin     bool              `protobuf:"varint,4,opt,name=fin" json:"fin,omitempty"`
}

func (m *H3Resp) Reset()                    { *m = H3Resp{} }
func (m *H3Resp) String() string            { return proto.CompactTextString(m) }
func (*H3Resp) ProtoMessage()               {}
func (*H3Resp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *H3Resp) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *H3Resp) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *H3Resp) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *H3Resp) GetFin() bool {
	if m != nil {
		return m.Fin
	}
	return false
}

func init() {
	proto.RegisterType((*H3Req)(nil), "example.H3Req")
	proto.RegisterType((*H3Resp)(nil), "example.H3Resp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for H3Wrapper service

type H3WrapperClient interface {
	Stream(ctx context.Context, opts ...grpc.CallOption) (H3Wrapper_StreamClient, error)
	OneWay(ctx context.Context, in *H3Req, opts ...grpc.CallOption) (*H3Resp, error)
}

type h3WrapperClient struct {
	cc *grpc.ClientConn
}

func NewH3WrapperClient(cc *grpc.ClientConn) H3WrapperClient {
	return &h3WrapperClient{cc}
}

func (c *h3WrapperClient) Stream(ctx context.Context, opts ...grpc.CallOption) (H3Wrapper_StreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_H3Wrapper_serviceDesc.Streams[0], c.cc, "/example.H3Wrapper/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &h3WrapperStreamClient{stream}
	return x, nil
}

type H3Wrapper_StreamClient interface {
	Send(*H3Req) error
	Recv() (*H3Resp, error)
	grpc.ClientStream
}

type h3WrapperStreamClient struct {
	grpc.ClientStream
}

func (x *h3WrapperStreamClient) Send(m *H3Req) error {
	return x.ClientStream.SendMsg(m)
}

func (x *h3WrapperStreamClient) Recv() (*H3Resp, error) {
	m := new(H3Resp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *h3WrapperClient) OneWay(ctx context.Context, in *H3Req, opts ...grpc.CallOption) (*H3Resp, error) {
	out := new(H3Resp)
	err := grpc.Invoke(ctx, "/example.H3Wrapper/OneWay", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for H3Wrapper service

type H3WrapperServer interface {
	Stream(H3Wrapper_StreamServer) error
	OneWay(context.Context, *H3Req) (*H3Resp, error)
}

func RegisterH3WrapperServer(s *grpc.Server, srv H3WrapperServer) {
	s.RegisterService(&_H3Wrapper_serviceDesc, srv)
}

func _H3Wrapper_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(H3WrapperServer).Stream(&h3WrapperStreamServer{stream})
}

type H3Wrapper_StreamServer interface {
	Send(*H3Resp) error
	Recv() (*H3Req, error)
	grpc.ServerStream
}

type h3WrapperStreamServer struct {
	grpc.ServerStream
}

func (x *h3WrapperStreamServer) Send(m *H3Resp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *h3WrapperStreamServer) Recv() (*H3Req, error) {
	m := new(H3Req)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _H3Wrapper_OneWay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(H3Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(H3WrapperServer).OneWay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.H3Wrapper/OneWay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(H3WrapperServer).OneWay(ctx, req.(*H3Req))
	}
	return interceptor(ctx, in, info, handler)
}

var _H3Wrapper_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.H3Wrapper",
	HandlerType: (*H3WrapperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OneWay",
			Handler:    _H3Wrapper_OneWay_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _H3Wrapper_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "example.example",
}

func init() { proto.RegisterFile("example.example", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x76, 0x31,
	0x72, 0xb1, 0x7a, 0x18, 0x07, 0xa5, 0x16, 0x0a, 0x49, 0x70, 0xb1, 0xe7, 0x16, 0xa7, 0x87, 0x54,
	0x16, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x42, 0xa6, 0x5c, 0xec, 0x19,
	0xa9, 0x89, 0x29, 0xa9, 0x45, 0xc5, 0x12, 0x4c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0xd2, 0x7a, 0x30,
	0xd3, 0xc0, 0x5a, 0xf5, 0x3c, 0x20, 0xb2, 0xae, 0x79, 0x25, 0x45, 0x95, 0x41, 0x30, 0xb5, 0x42,
	0x42, 0x5c, 0x2c, 0x49, 0xf9, 0x29, 0x95, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x60, 0xb6,
	0x90, 0x00, 0x17, 0x73, 0x5a, 0x66, 0x9e, 0x04, 0x8b, 0x02, 0xa3, 0x06, 0x47, 0x10, 0x88, 0x29,
	0x65, 0xc5, 0xc5, 0x83, 0xac, 0x1d, 0xa4, 0x22, 0x3b, 0xb5, 0x12, 0xea, 0x04, 0x10, 0x53, 0x48,
	0x84, 0x8b, 0xb5, 0x2c, 0x31, 0xa7, 0x34, 0x55, 0x82, 0x09, 0x2c, 0x06, 0xe1, 0x58, 0x31, 0x59,
	0x30, 0x2a, 0xed, 0x61, 0xe4, 0x62, 0x03, 0xb9, 0xa0, 0xb8, 0x00, 0x8f, 0xeb, 0xcd, 0xd0, 0x5d,
	0x2f, 0x83, 0xe2, 0xfa, 0xe2, 0x82, 0x81, 0x72, 0xbe, 0x51, 0x2a, 0x17, 0xa7, 0x87, 0x71, 0x78,
	0x51, 0x62, 0x41, 0x41, 0x6a, 0x91, 0x90, 0x2e, 0x17, 0x5b, 0x70, 0x49, 0x51, 0x6a, 0x62, 0xae,
	0x10, 0x1f, 0x6a, 0xe8, 0x4a, 0xf1, 0xa3, 0xb9, 0x57, 0x83, 0xd1, 0x80, 0x51, 0x48, 0x93, 0x8b,
	0xcd, 0x3f, 0x2f, 0x35, 0x3c, 0xb1, 0x92, 0xa0, 0xf2, 0x24, 0x36, 0x70, 0x94, 0x1b, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x17, 0x65, 0x07, 0x6b, 0x03, 0x02, 0x00, 0x00,
}