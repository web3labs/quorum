// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc.proto

package server

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClientClient is the client API for Client service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientClient interface {
	Version(ctx context.Context, in *ApiVersion, opts ...grpc.CallOption) (*ApiVersion, error)
	Upcheck(ctx context.Context, in *UpCheckResponse, opts ...grpc.CallOption) (*UpCheckResponse, error)
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
	Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (*ReceiveResponse, error)
	UpdatePartyInfo(ctx context.Context, in *PartyInfo, opts ...grpc.CallOption) (*PartyInfoResponse, error)
	Push(ctx context.Context, in *PushPayload, opts ...grpc.CallOption) (*PartyInfoResponse, error)
}

type clientClient struct {
	cc *grpc.ClientConn
}

func NewClientClient(cc *grpc.ClientConn) ClientClient {
	return &clientClient{cc}
}

func (c *clientClient) Version(ctx context.Context, in *ApiVersion, opts ...grpc.CallOption) (*ApiVersion, error) {
	out := new(ApiVersion)
	err := c.cc.Invoke(ctx, "/server.Client/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) Upcheck(ctx context.Context, in *UpCheckResponse, opts ...grpc.CallOption) (*UpCheckResponse, error) {
	out := new(UpCheckResponse)
	err := c.cc.Invoke(ctx, "/server.Client/Upcheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/server.Client/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (*ReceiveResponse, error) {
	out := new(ReceiveResponse)
	err := c.cc.Invoke(ctx, "/server.Client/Receive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) UpdatePartyInfo(ctx context.Context, in *PartyInfo, opts ...grpc.CallOption) (*PartyInfoResponse, error) {
	out := new(PartyInfoResponse)
	err := c.cc.Invoke(ctx, "/server.Client/UpdatePartyInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) Push(ctx context.Context, in *PushPayload, opts ...grpc.CallOption) (*PartyInfoResponse, error) {
	out := new(PartyInfoResponse)
	err := c.cc.Invoke(ctx, "/server.Client/Push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServer is the server API for Client service.
type ClientServer interface {
	Version(context.Context, *ApiVersion) (*ApiVersion, error)
	Upcheck(context.Context, *UpCheckResponse) (*UpCheckResponse, error)
	Send(context.Context, *SendRequest) (*SendResponse, error)
	Receive(context.Context, *ReceiveRequest) (*ReceiveResponse, error)
	UpdatePartyInfo(context.Context, *PartyInfo) (*PartyInfoResponse, error)
	Push(context.Context, *PushPayload) (*PartyInfoResponse, error)
}

func RegisterClientServer(s *grpc.Server, srv ClientServer) {
	s.RegisterService(&_Client_serviceDesc, srv)
}

func _Client_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiVersion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Client/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).Version(ctx, req.(*ApiVersion))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_Upcheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpCheckResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).Upcheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Client/Upcheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).Upcheck(ctx, req.(*UpCheckResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Client/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_Receive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).Receive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Client/Receive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).Receive(ctx, req.(*ReceiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_UpdatePartyInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PartyInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).UpdatePartyInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Client/UpdatePartyInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).UpdatePartyInfo(ctx, req.(*PartyInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Client/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).Push(ctx, req.(*PushPayload))
	}
	return interceptor(ctx, in, info, handler)
}

var _Client_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.Client",
	HandlerType: (*ClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _Client_Version_Handler,
		},
		{
			MethodName: "Upcheck",
			Handler:    _Client_Upcheck_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _Client_Send_Handler,
		},
		{
			MethodName: "Receive",
			Handler:    _Client_Receive_Handler,
		},
		{
			MethodName: "UpdatePartyInfo",
			Handler:    _Client_UpdatePartyInfo_Handler,
		},
		{
			MethodName: "Push",
			Handler:    _Client_Push_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}

func init() { proto.RegisterFile("grpc.proto", fileDescriptor_grpc_65f146c8ea2192a0) }

var fileDescriptor_grpc_65f146c8ea2192a0 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcf, 0x4a, 0xc3, 0x30,
	0x1c, 0xc7, 0x9d, 0xce, 0x56, 0x02, 0x3a, 0xcd, 0xd4, 0x61, 0xf0, 0xd4, 0xe3, 0x0e, 0x2d, 0xb8,
	0xdb, 0xf0, 0x22, 0x03, 0xd1, 0x5b, 0xa9, 0x74, 0xf7, 0xd8, 0xfe, 0xd6, 0x06, 0x6b, 0x12, 0x93,
	0xb4, 0xb0, 0xab, 0xaf, 0xe0, 0xc3, 0xf8, 0x20, 0xbe, 0x82, 0x0f, 0x22, 0xe9, 0x3f, 0x95, 0xe9,
	0x8e, 0xbf, 0x4f, 0xbe, 0xdf, 0x0f, 0x7c, 0x09, 0x42, 0x99, 0x92, 0x89, 0x2f, 0x95, 0x30, 0x02,
	0x3b, 0x1a, 0x54, 0x05, 0x8a, 0x5c, 0x66, 0x42, 0x64, 0x05, 0x04, 0x54, 0xb2, 0x80, 0x72, 0x2e,
	0x0c, 0x35, 0x4c, 0x70, 0xdd, 0xa4, 0xc8, 0xd1, 0x33, 0x68, 0x4d, 0x33, 0x68, 0xef, 0xab, 0xf7,
	0x3d, 0xe4, 0x2c, 0x0a, 0x06, 0xdc, 0xe0, 0x5b, 0xe4, 0x2e, 0x41, 0x69, 0x26, 0x38, 0xc6, 0x7e,
	0x23, 0xf3, 0x6f, 0x24, 0x6b, 0x19, 0xf9, 0x83, 0x79, 0xe3, 0xd7, 0x8f, 0xcf, 0xb7, 0xdd, 0x43,
	0xef, 0x20, 0xa8, 0x1a, 0x32, 0x1f, 0x4c, 0x71, 0x88, 0xdc, 0x58, 0x26, 0x39, 0x24, 0x4f, 0x78,
	0xd2, 0x75, 0x62, 0xb9, 0xb0, 0x20, 0x02, 0x2d, 0x05, 0xd7, 0x40, 0xfe, 0x7b, 0xf8, 0x61, 0x2c,
	0x1b, 0x87, 0x35, 0xce, 0xd0, 0xf0, 0x01, 0x78, 0x8a, 0xc7, 0x5d, 0xcb, 0x5e, 0x11, 0xbc, 0x94,
	0xa0, 0x0d, 0x39, 0xfd, 0x0d, 0x5b, 0xcf, 0x0e, 0xbe, 0x46, 0x6e, 0x04, 0x09, 0xb0, 0x0a, 0xf0,
	0x79, 0x17, 0x69, 0x41, 0x57, 0x9d, 0x6c, 0xf0, 0xbe, 0xbd, 0x44, 0xa3, 0x58, 0xa6, 0xd4, 0x40,
	0x48, 0x95, 0x59, 0xdf, 0xf3, 0x95, 0xc0, 0x27, 0x5d, 0xba, 0x47, 0xe4, 0x62, 0x03, 0xf5, 0x8a,
	0xb3, 0x7a, 0xc8, 0xc8, 0x43, 0x81, 0xb4, 0x6f, 0x8c, 0xaf, 0x84, 0x9d, 0x72, 0x87, 0x86, 0x61,
	0xa9, 0xf3, 0xef, 0x29, 0xf6, 0x0a, 0xe9, 0xba, 0x10, 0x34, 0xdd, 0xa6, 0x3b, 0xae, 0x75, 0xc8,
	0xdb, 0x0f, 0x64, 0xa9, 0xf3, 0xf9, 0x60, 0xfa, 0xe8, 0xd4, 0x1f, 0x38, 0xfb, 0x0a, 0x00, 0x00,
	0xff, 0xff, 0xe7, 0x59, 0x1e, 0x07, 0x04, 0x02, 0x00, 0x00,
}
