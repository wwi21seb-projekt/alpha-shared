// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: server_alpha/user/v1/authentication.proto

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
	AuthenticationService_RegisterUser_FullMethodName          = "/server_alpha.user.v1.AuthenticationService/RegisterUser"
	AuthenticationService_ResendActivationEmail_FullMethodName = "/server_alpha.user.v1.AuthenticationService/ResendActivationEmail"
	AuthenticationService_ActivateUser_FullMethodName          = "/server_alpha.user.v1.AuthenticationService/ActivateUser"
	AuthenticationService_LoginUser_FullMethodName             = "/server_alpha.user.v1.AuthenticationService/LoginUser"
	AuthenticationService_UpdatePassword_FullMethodName        = "/server_alpha.user.v1.AuthenticationService/UpdatePassword"
	AuthenticationService_ResetPassword_FullMethodName         = "/server_alpha.user.v1.AuthenticationService/ResetPassword"
	AuthenticationService_SetPassword_FullMethodName           = "/server_alpha.user.v1.AuthenticationService/SetPassword"
)

// AuthenticationServiceClient is the client API for AuthenticationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticationServiceClient interface {
	RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error)
	ResendActivationEmail(ctx context.Context, in *ResendActivationEmailRequest, opts ...grpc.CallOption) (*ResendActivationEmailResponse, error)
	ActivateUser(ctx context.Context, in *ActivateUserRequest, opts ...grpc.CallOption) (*ActivateUserResponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
	UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*UpdatePasswordResponse, error)
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error)
	SetPassword(ctx context.Context, in *SetPasswordRequest, opts ...grpc.CallOption) (*SetPasswordResponse, error)
}

type authenticationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationServiceClient(cc grpc.ClientConnInterface) AuthenticationServiceClient {
	return &authenticationServiceClient{cc}
}

func (c *authenticationServiceClient) RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterUserResponse)
	err := c.cc.Invoke(ctx, AuthenticationService_RegisterUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) ResendActivationEmail(ctx context.Context, in *ResendActivationEmailRequest, opts ...grpc.CallOption) (*ResendActivationEmailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResendActivationEmailResponse)
	err := c.cc.Invoke(ctx, AuthenticationService_ResendActivationEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) ActivateUser(ctx context.Context, in *ActivateUserRequest, opts ...grpc.CallOption) (*ActivateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ActivateUserResponse)
	err := c.cc.Invoke(ctx, AuthenticationService_ActivateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, AuthenticationService_LoginUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*UpdatePasswordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePasswordResponse)
	err := c.cc.Invoke(ctx, AuthenticationService_UpdatePassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResetPasswordResponse)
	err := c.cc.Invoke(ctx, AuthenticationService_ResetPassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) SetPassword(ctx context.Context, in *SetPasswordRequest, opts ...grpc.CallOption) (*SetPasswordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetPasswordResponse)
	err := c.cc.Invoke(ctx, AuthenticationService_SetPassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServiceServer is the server API for AuthenticationService service.
// All implementations must embed UnimplementedAuthenticationServiceServer
// for forward compatibility
type AuthenticationServiceServer interface {
	RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserResponse, error)
	ResendActivationEmail(context.Context, *ResendActivationEmailRequest) (*ResendActivationEmailResponse, error)
	ActivateUser(context.Context, *ActivateUserRequest) (*ActivateUserResponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*UpdatePasswordResponse, error)
	ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error)
	SetPassword(context.Context, *SetPasswordRequest) (*SetPasswordResponse, error)
	mustEmbedUnimplementedAuthenticationServiceServer()
}

// UnimplementedAuthenticationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServiceServer struct {
}

func (UnimplementedAuthenticationServiceServer) RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedAuthenticationServiceServer) ResendActivationEmail(context.Context, *ResendActivationEmailRequest) (*ResendActivationEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResendActivationEmail not implemented")
}
func (UnimplementedAuthenticationServiceServer) ActivateUser(context.Context, *ActivateUserRequest) (*ActivateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateUser not implemented")
}
func (UnimplementedAuthenticationServiceServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedAuthenticationServiceServer) UpdatePassword(context.Context, *UpdatePasswordRequest) (*UpdatePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedAuthenticationServiceServer) ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedAuthenticationServiceServer) SetPassword(context.Context, *SetPasswordRequest) (*SetPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPassword not implemented")
}
func (UnimplementedAuthenticationServiceServer) mustEmbedUnimplementedAuthenticationServiceServer() {}

// UnsafeAuthenticationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticationServiceServer will
// result in compilation errors.
type UnsafeAuthenticationServiceServer interface {
	mustEmbedUnimplementedAuthenticationServiceServer()
}

func RegisterAuthenticationServiceServer(s grpc.ServiceRegistrar, srv AuthenticationServiceServer) {
	s.RegisterService(&AuthenticationService_ServiceDesc, srv)
}

func _AuthenticationService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationService_RegisterUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).RegisterUser(ctx, req.(*RegisterUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_ResendActivationEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResendActivationEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).ResendActivationEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationService_ResendActivationEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).ResendActivationEmail(ctx, req.(*ResendActivationEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_ActivateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).ActivateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationService_ActivateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).ActivateUser(ctx, req.(*ActivateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationService_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationService_UpdatePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).UpdatePassword(ctx, req.(*UpdatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationService_ResetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).ResetPassword(ctx, req.(*ResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_SetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).SetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationService_SetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).SetPassword(ctx, req.(*SetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthenticationService_ServiceDesc is the grpc.ServiceDesc for AuthenticationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthenticationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server_alpha.user.v1.AuthenticationService",
	HandlerType: (*AuthenticationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _AuthenticationService_RegisterUser_Handler,
		},
		{
			MethodName: "ResendActivationEmail",
			Handler:    _AuthenticationService_ResendActivationEmail_Handler,
		},
		{
			MethodName: "ActivateUser",
			Handler:    _AuthenticationService_ActivateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _AuthenticationService_LoginUser_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _AuthenticationService_UpdatePassword_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _AuthenticationService_ResetPassword_Handler,
		},
		{
			MethodName: "SetPassword",
			Handler:    _AuthenticationService_SetPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server_alpha/user/v1/authentication.proto",
}
