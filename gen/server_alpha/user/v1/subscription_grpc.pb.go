// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: server_alpha/user/v1/subscription.proto

package userv1

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
	SubscriptionService_ListSubscriptions_FullMethodName  = "/server_alpha.user.v1.SubscriptionService/ListSubscriptions"
	SubscriptionService_CreateSubscription_FullMethodName = "/server_alpha.user.v1.SubscriptionService/CreateSubscription"
	SubscriptionService_DeleteSubscription_FullMethodName = "/server_alpha.user.v1.SubscriptionService/DeleteSubscription"
)

// SubscriptionServiceClient is the client API for SubscriptionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubscriptionServiceClient interface {
	ListSubscriptions(ctx context.Context, in *ListSubscriptionsRequest, opts ...grpc.CallOption) (*ListSubscriptionsResponse, error)
	CreateSubscription(ctx context.Context, in *CreateSubscriptionRequest, opts ...grpc.CallOption) (*CreateSubscriptionResponse, error)
	DeleteSubscription(ctx context.Context, in *DeleteSubscriptionRequest, opts ...grpc.CallOption) (*DeleteSubscriptionResponse, error)
}

type subscriptionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubscriptionServiceClient(cc grpc.ClientConnInterface) SubscriptionServiceClient {
	return &subscriptionServiceClient{cc}
}

func (c *subscriptionServiceClient) ListSubscriptions(ctx context.Context, in *ListSubscriptionsRequest, opts ...grpc.CallOption) (*ListSubscriptionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListSubscriptionsResponse)
	err := c.cc.Invoke(ctx, SubscriptionService_ListSubscriptions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) CreateSubscription(ctx context.Context, in *CreateSubscriptionRequest, opts ...grpc.CallOption) (*CreateSubscriptionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateSubscriptionResponse)
	err := c.cc.Invoke(ctx, SubscriptionService_CreateSubscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) DeleteSubscription(ctx context.Context, in *DeleteSubscriptionRequest, opts ...grpc.CallOption) (*DeleteSubscriptionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteSubscriptionResponse)
	err := c.cc.Invoke(ctx, SubscriptionService_DeleteSubscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscriptionServiceServer is the server API for SubscriptionService service.
// All implementations must embed UnimplementedSubscriptionServiceServer
// for forward compatibility
type SubscriptionServiceServer interface {
	ListSubscriptions(context.Context, *ListSubscriptionsRequest) (*ListSubscriptionsResponse, error)
	CreateSubscription(context.Context, *CreateSubscriptionRequest) (*CreateSubscriptionResponse, error)
	DeleteSubscription(context.Context, *DeleteSubscriptionRequest) (*DeleteSubscriptionResponse, error)
	mustEmbedUnimplementedSubscriptionServiceServer()
}

// UnimplementedSubscriptionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSubscriptionServiceServer struct {
}

func (UnimplementedSubscriptionServiceServer) ListSubscriptions(context.Context, *ListSubscriptionsRequest) (*ListSubscriptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSubscriptions not implemented")
}
func (UnimplementedSubscriptionServiceServer) CreateSubscription(context.Context, *CreateSubscriptionRequest) (*CreateSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubscription not implemented")
}
func (UnimplementedSubscriptionServiceServer) DeleteSubscription(context.Context, *DeleteSubscriptionRequest) (*DeleteSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSubscription not implemented")
}
func (UnimplementedSubscriptionServiceServer) mustEmbedUnimplementedSubscriptionServiceServer() {}

// UnsafeSubscriptionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubscriptionServiceServer will
// result in compilation errors.
type UnsafeSubscriptionServiceServer interface {
	mustEmbedUnimplementedSubscriptionServiceServer()
}

func RegisterSubscriptionServiceServer(s grpc.ServiceRegistrar, srv SubscriptionServiceServer) {
	s.RegisterService(&SubscriptionService_ServiceDesc, srv)
}

func _SubscriptionService_ListSubscriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSubscriptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).ListSubscriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_ListSubscriptions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).ListSubscriptions(ctx, req.(*ListSubscriptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_CreateSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).CreateSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_CreateSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).CreateSubscription(ctx, req.(*CreateSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_DeleteSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).DeleteSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_DeleteSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).DeleteSubscription(ctx, req.(*DeleteSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SubscriptionService_ServiceDesc is the grpc.ServiceDesc for SubscriptionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubscriptionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server_alpha.user.v1.SubscriptionService",
	HandlerType: (*SubscriptionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSubscriptions",
			Handler:    _SubscriptionService_ListSubscriptions_Handler,
		},
		{
			MethodName: "CreateSubscription",
			Handler:    _SubscriptionService_CreateSubscription_Handler,
		},
		{
			MethodName: "DeleteSubscription",
			Handler:    _SubscriptionService_DeleteSubscription_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server_alpha/user/v1/subscription.proto",
}
