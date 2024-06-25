// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: server_alpha/chat/v1/chat.proto

package chatv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ChatService_CreateChat_FullMethodName        = "/server_alpha.chat.v1.ChatService/CreateChat"
	ChatService_GetChat_FullMethodName           = "/server_alpha.chat.v1.ChatService/GetChat"
	ChatService_ListChats_FullMethodName         = "/server_alpha.chat.v1.ChatService/ListChats"
	ChatService_PrepareChatStream_FullMethodName = "/server_alpha.chat.v1.ChatService/PrepareChatStream"
	ChatService_ChatMessage_FullMethodName       = "/server_alpha.chat.v1.ChatService/ChatMessage"
)

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	CreateChat(ctx context.Context, in *CreateChatRequest, opts ...grpc.CallOption) (*CreateChatResponse, error)
	GetChat(ctx context.Context, in *GetChatRequest, opts ...grpc.CallOption) (*GetChatResponse, error)
	ListChats(ctx context.Context, in *ListChatsRequest, opts ...grpc.CallOption) (*ListChatsResponse, error)
	PrepareChatStream(ctx context.Context, in *PrepareChatStreamRequest, opts ...grpc.CallOption) (*PrepareChatStreamResponse, error)
	ChatMessage(ctx context.Context, opts ...grpc.CallOption) (ChatService_ChatMessageClient, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) CreateChat(ctx context.Context, in *CreateChatRequest, opts ...grpc.CallOption) (*CreateChatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateChatResponse)
	err := c.cc.Invoke(ctx, ChatService_CreateChat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetChat(ctx context.Context, in *GetChatRequest, opts ...grpc.CallOption) (*GetChatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetChatResponse)
	err := c.cc.Invoke(ctx, ChatService_GetChat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) ListChats(ctx context.Context, in *ListChatsRequest, opts ...grpc.CallOption) (*ListChatsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListChatsResponse)
	err := c.cc.Invoke(ctx, ChatService_ListChats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) PrepareChatStream(ctx context.Context, in *PrepareChatStreamRequest, opts ...grpc.CallOption) (*PrepareChatStreamResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PrepareChatStreamResponse)
	err := c.cc.Invoke(ctx, ChatService_PrepareChatStream_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) ChatMessage(ctx context.Context, opts ...grpc.CallOption) (ChatService_ChatMessageClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], ChatService_ChatMessage_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceChatMessageClient{ClientStream: stream}
	return x, nil
}

type ChatService_ChatMessageClient interface {
	Send(*ChatMessageRequest) error
	Recv() (*ChatMessageResponse, error)
	grpc.ClientStream
}

type chatServiceChatMessageClient struct {
	grpc.ClientStream
}

func (x *chatServiceChatMessageClient) Send(m *ChatMessageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceChatMessageClient) Recv() (*ChatMessageResponse, error) {
	m := new(ChatMessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	CreateChat(context.Context, *CreateChatRequest) (*CreateChatResponse, error)
	GetChat(context.Context, *GetChatRequest) (*GetChatResponse, error)
	ListChats(context.Context, *ListChatsRequest) (*ListChatsResponse, error)
	PrepareChatStream(context.Context, *PrepareChatStreamRequest) (*PrepareChatStreamResponse, error)
	ChatMessage(ChatService_ChatMessageServer) error
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) CreateChat(context.Context, *CreateChatRequest) (*CreateChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChat not implemented")
}
func (UnimplementedChatServiceServer) GetChat(context.Context, *GetChatRequest) (*GetChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChat not implemented")
}
func (UnimplementedChatServiceServer) ListChats(context.Context, *ListChatsRequest) (*ListChatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListChats not implemented")
}
func (UnimplementedChatServiceServer) PrepareChatStream(context.Context, *PrepareChatStreamRequest) (*PrepareChatStreamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareChatStream not implemented")
}
func (UnimplementedChatServiceServer) ChatMessage(ChatService_ChatMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method ChatMessage not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_CreateChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CreateChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_CreateChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CreateChat(ctx, req.(*CreateChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetChat(ctx, req.(*GetChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_ListChats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListChatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).ListChats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_ListChats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).ListChats(ctx, req.(*ListChatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_PrepareChatStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareChatStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).PrepareChatStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_PrepareChatStream_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).PrepareChatStream(ctx, req.(*PrepareChatStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_ChatMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).ChatMessage(&chatServiceChatMessageServer{ServerStream: stream})
}

type ChatService_ChatMessageServer interface {
	Send(*ChatMessageResponse) error
	Recv() (*ChatMessageRequest, error)
	grpc.ServerStream
}

type chatServiceChatMessageServer struct {
	grpc.ServerStream
}

func (x *chatServiceChatMessageServer) Send(m *ChatMessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceChatMessageServer) Recv() (*ChatMessageRequest, error) {
	m := new(ChatMessageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server_alpha.chat.v1.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChat",
			Handler:    _ChatService_CreateChat_Handler,
		},
		{
			MethodName: "GetChat",
			Handler:    _ChatService_GetChat_Handler,
		},
		{
			MethodName: "ListChats",
			Handler:    _ChatService_ListChats_Handler,
		},
		{
			MethodName: "PrepareChatStream",
			Handler:    _ChatService_PrepareChatStream_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ChatMessage",
			Handler:       _ChatService_ChatMessage_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "server_alpha/chat/v1/chat.proto",
}
