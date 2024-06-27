// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: server_alpha/mail/v1/mail.proto

package mailv1

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
	MailService_SendTokenMail_FullMethodName        = "/server_alpha.mail.v1.MailService/SendTokenMail"
	MailService_SendConfirmationMail_FullMethodName = "/server_alpha.mail.v1.MailService/SendConfirmationMail"
)

// MailServiceClient is the client API for MailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MailServiceClient interface {
	SendTokenMail(ctx context.Context, in *SendTokenMailRequest, opts ...grpc.CallOption) (*SendTokenMailResponse, error)
	SendConfirmationMail(ctx context.Context, in *SendConfirmationMailRequest, opts ...grpc.CallOption) (*SendConfirmationMailResponse, error)
}

type mailServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMailServiceClient(cc grpc.ClientConnInterface) MailServiceClient {
	return &mailServiceClient{cc}
}

func (c *mailServiceClient) SendTokenMail(ctx context.Context, in *SendTokenMailRequest, opts ...grpc.CallOption) (*SendTokenMailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendTokenMailResponse)
	err := c.cc.Invoke(ctx, MailService_SendTokenMail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailServiceClient) SendConfirmationMail(ctx context.Context, in *SendConfirmationMailRequest, opts ...grpc.CallOption) (*SendConfirmationMailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendConfirmationMailResponse)
	err := c.cc.Invoke(ctx, MailService_SendConfirmationMail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MailServiceServer is the server API for MailService service.
// All implementations must embed UnimplementedMailServiceServer
// for forward compatibility
type MailServiceServer interface {
	SendTokenMail(context.Context, *SendTokenMailRequest) (*SendTokenMailResponse, error)
	SendConfirmationMail(context.Context, *SendConfirmationMailRequest) (*SendConfirmationMailResponse, error)
	mustEmbedUnimplementedMailServiceServer()
}

// UnimplementedMailServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMailServiceServer struct {
}

func (UnimplementedMailServiceServer) SendTokenMail(context.Context, *SendTokenMailRequest) (*SendTokenMailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTokenMail not implemented")
}
func (UnimplementedMailServiceServer) SendConfirmationMail(context.Context, *SendConfirmationMailRequest) (*SendConfirmationMailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendConfirmationMail not implemented")
}
func (UnimplementedMailServiceServer) mustEmbedUnimplementedMailServiceServer() {}

// UnsafeMailServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MailServiceServer will
// result in compilation errors.
type UnsafeMailServiceServer interface {
	mustEmbedUnimplementedMailServiceServer()
}

func RegisterMailServiceServer(s grpc.ServiceRegistrar, srv MailServiceServer) {
	s.RegisterService(&MailService_ServiceDesc, srv)
}

func _MailService_SendTokenMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTokenMailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailServiceServer).SendTokenMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MailService_SendTokenMail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailServiceServer).SendTokenMail(ctx, req.(*SendTokenMailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MailService_SendConfirmationMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendConfirmationMailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailServiceServer).SendConfirmationMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MailService_SendConfirmationMail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailServiceServer).SendConfirmationMail(ctx, req.(*SendConfirmationMailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MailService_ServiceDesc is the grpc.ServiceDesc for MailService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MailService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server_alpha.mail.v1.MailService",
	HandlerType: (*MailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendTokenMail",
			Handler:    _MailService_SendTokenMail_Handler,
		},
		{
			MethodName: "SendConfirmationMail",
			Handler:    _MailService_SendConfirmationMail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server_alpha/mail/v1/mail.proto",
}