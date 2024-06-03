// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.26.1
// source: notification/notfication-service.proto


package serveralpha

import (
	context "context"
	common "github.com/wwi21seb-projekt/alpha-shared/proto/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	PushService_GetPublicKey_FullMethodName           = "/serveralpha.notification.PushService/GetPublicKey"
	PushService_CreatePushSubscription_FullMethodName = "/serveralpha.notification.PushService/CreatePushSubscription"
)

// PushServiceClient is the client API for PushService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PushServiceClient interface {
	GetPublicKey(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*PublicKeyResponse, error)
	CreatePushSubscription(ctx context.Context, in *CreatePushSubscriptionRequest, opts ...grpc.CallOption) (*CreatePushSubscriptionResponse, error)
}

type pushServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPushServiceClient(cc grpc.ClientConnInterface) PushServiceClient {
	return &pushServiceClient{cc}
}

func (c *pushServiceClient) GetPublicKey(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*PublicKeyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublicKeyResponse)
	err := c.cc.Invoke(ctx, PushService_GetPublicKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushServiceClient) CreatePushSubscription(ctx context.Context, in *CreatePushSubscriptionRequest, opts ...grpc.CallOption) (*CreatePushSubscriptionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePushSubscriptionResponse)
	err := c.cc.Invoke(ctx, PushService_CreatePushSubscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PushServiceServer is the server API for PushService service.
// All implementations must embed UnimplementedPushServiceServer
// for forward compatibility
type PushServiceServer interface {
	GetPublicKey(context.Context, *common.Empty) (*PublicKeyResponse, error)
	CreatePushSubscription(context.Context, *CreatePushSubscriptionRequest) (*CreatePushSubscriptionResponse, error)
	mustEmbedUnimplementedPushServiceServer()
}

// UnimplementedPushServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPushServiceServer struct {
}

func (UnimplementedPushServiceServer) GetPublicKey(context.Context, *common.Empty) (*PublicKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicKey not implemented")
}
func (UnimplementedPushServiceServer) CreatePushSubscription(context.Context, *CreatePushSubscriptionRequest) (*CreatePushSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePushSubscription not implemented")
}
func (UnimplementedPushServiceServer) mustEmbedUnimplementedPushServiceServer() {}

// UnsafePushServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PushServiceServer will
// result in compilation errors.
type UnsafePushServiceServer interface {
	mustEmbedUnimplementedPushServiceServer()
}

func RegisterPushServiceServer(s grpc.ServiceRegistrar, srv PushServiceServer) {
	s.RegisterService(&PushService_ServiceDesc, srv)
}

func _PushService_GetPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServiceServer).GetPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PushService_GetPublicKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServiceServer).GetPublicKey(ctx, req.(*common.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PushService_CreatePushSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePushSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServiceServer).CreatePushSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PushService_CreatePushSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServiceServer).CreatePushSubscription(ctx, req.(*CreatePushSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PushService_ServiceDesc is the grpc.ServiceDesc for PushService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PushService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "serveralpha.notification.PushService",
	HandlerType: (*PushServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPublicKey",
			Handler:    _PushService_GetPublicKey_Handler,
		},
		{
			MethodName: "CreatePushSubscription",
			Handler:    _PushService_CreatePushSubscription_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notification/notfication-service.proto",

}

const (
	NotificationService_GetNotifications_FullMethodName   = "/serveralpha.notification.NotificationService/GetNotifications"
	NotificationService_DeleteNotification_FullMethodName = "/serveralpha.notification.NotificationService/DeleteNotification"
)

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationServiceClient interface {
	GetNotifications(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*GetNotificationsResponse, error)
	DeleteNotification(ctx context.Context, in *DeleteNotificationRequest, opts ...grpc.CallOption) (*common.Empty, error)
}

type notificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationServiceClient(cc grpc.ClientConnInterface) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) GetNotifications(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*GetNotificationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNotificationsResponse)
	err := c.cc.Invoke(ctx, NotificationService_GetNotifications_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) DeleteNotification(ctx context.Context, in *DeleteNotificationRequest, opts ...grpc.CallOption) (*common.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, NotificationService_DeleteNotification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
// All implementations must embed UnimplementedNotificationServiceServer
// for forward compatibility
type NotificationServiceServer interface {
	GetNotifications(context.Context, *common.Empty) (*GetNotificationsResponse, error)
	DeleteNotification(context.Context, *DeleteNotificationRequest) (*common.Empty, error)
	mustEmbedUnimplementedNotificationServiceServer()
}

// UnimplementedNotificationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationServiceServer struct {
}

func (UnimplementedNotificationServiceServer) GetNotifications(context.Context, *common.Empty) (*GetNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotifications not implemented")
}
func (UnimplementedNotificationServiceServer) DeleteNotification(context.Context, *DeleteNotificationRequest) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNotification not implemented")
}
func (UnimplementedNotificationServiceServer) mustEmbedUnimplementedNotificationServiceServer() {}

// UnsafeNotificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationServiceServer will
// result in compilation errors.
type UnsafeNotificationServiceServer interface {
	mustEmbedUnimplementedNotificationServiceServer()
}

func RegisterNotificationServiceServer(s grpc.ServiceRegistrar, srv NotificationServiceServer) {
	s.RegisterService(&NotificationService_ServiceDesc, srv)
}

func _NotificationService_GetNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).GetNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_GetNotifications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).GetNotifications(ctx, req.(*common.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_DeleteNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).DeleteNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_DeleteNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).DeleteNotification(ctx, req.(*DeleteNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationService_ServiceDesc is the grpc.ServiceDesc for NotificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "serveralpha.notification.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNotifications",
			Handler:    _NotificationService_GetNotifications_Handler,
		},
		{
			MethodName: "DeleteNotification",
			Handler:    _NotificationService_DeleteNotification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notification/notfication-service.proto",

}
