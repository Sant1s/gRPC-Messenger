// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: service.proto

package messenger

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MessengerClient is the client API for Messenger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessengerClient interface {
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*MessageNotification, error)
	ReceiveMessage(ctx context.Context, in *ReceiveMessageRequest, opts ...grpc.CallOption) (Messenger_ReceiveMessageClient, error)
}

type messengerClient struct {
	cc grpc.ClientConnInterface
}

func NewMessengerClient(cc grpc.ClientConnInterface) MessengerClient {
	return &messengerClient{cc}
}

func (c *messengerClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*MessageNotification, error) {
	out := new(MessageNotification)
	err := c.cc.Invoke(ctx, "/messenger.Messenger/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messengerClient) ReceiveMessage(ctx context.Context, in *ReceiveMessageRequest, opts ...grpc.CallOption) (Messenger_ReceiveMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &Messenger_ServiceDesc.Streams[0], "/messenger.Messenger/ReceiveMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &messengerReceiveMessageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Messenger_ReceiveMessageClient interface {
	Recv() (*MessageNotification, error)
	grpc.ClientStream
}

type messengerReceiveMessageClient struct {
	grpc.ClientStream
}

func (x *messengerReceiveMessageClient) Recv() (*MessageNotification, error) {
	m := new(MessageNotification)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessengerServer is the server API for Messenger service.
// All implementations must embed UnimplementedMessengerServer
// for forward compatibility
type MessengerServer interface {
	SendMessage(context.Context, *SendMessageRequest) (*MessageNotification, error)
	ReceiveMessage(*ReceiveMessageRequest, Messenger_ReceiveMessageServer) error
	mustEmbedUnimplementedMessengerServer()
}

// UnimplementedMessengerServer must be embedded to have forward compatible implementations.
type UnimplementedMessengerServer struct {
}

func (UnimplementedMessengerServer) SendMessage(context.Context, *SendMessageRequest) (*MessageNotification, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedMessengerServer) ReceiveMessage(*ReceiveMessageRequest, Messenger_ReceiveMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveMessage not implemented")
}
func (UnimplementedMessengerServer) mustEmbedUnimplementedMessengerServer() {}

// UnsafeMessengerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessengerServer will
// result in compilation errors.
type UnsafeMessengerServer interface {
	mustEmbedUnimplementedMessengerServer()
}

func RegisterMessengerServer(s grpc.ServiceRegistrar, srv MessengerServer) {
	s.RegisterService(&Messenger_ServiceDesc, srv)
}

func _Messenger_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messenger.Messenger/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messenger_ReceiveMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReceiveMessageRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessengerServer).ReceiveMessage(m, &messengerReceiveMessageServer{stream})
}

type Messenger_ReceiveMessageServer interface {
	Send(*MessageNotification) error
	grpc.ServerStream
}

type messengerReceiveMessageServer struct {
	grpc.ServerStream
}

func (x *messengerReceiveMessageServer) Send(m *MessageNotification) error {
	return x.ServerStream.SendMsg(m)
}

// Messenger_ServiceDesc is the grpc.ServiceDesc for Messenger service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Messenger_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "messenger.Messenger",
	HandlerType: (*MessengerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _Messenger_SendMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReceiveMessage",
			Handler:       _Messenger_ReceiveMessage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service.proto",
}