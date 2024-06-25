// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: server_alpha/image/v1/image.proto

package imagev1

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
	ImageService_UploadImage_FullMethodName = "/server_alpha.image.v1.ImageService/UploadImage"
	ImageService_GetImage_FullMethodName    = "/server_alpha.image.v1.ImageService/GetImage"
)

// ImageServiceClient is the client API for ImageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageServiceClient interface {
	UploadImage(ctx context.Context, in *UploadImageRequest, opts ...grpc.CallOption) (*UploadImageResponse, error)
	GetImage(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (*GetImageResponse, error)
}

type imageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImageServiceClient(cc grpc.ClientConnInterface) ImageServiceClient {
	return &imageServiceClient{cc}
}

func (c *imageServiceClient) UploadImage(ctx context.Context, in *UploadImageRequest, opts ...grpc.CallOption) (*UploadImageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadImageResponse)
	err := c.cc.Invoke(ctx, ImageService_UploadImage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageServiceClient) GetImage(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (*GetImageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetImageResponse)
	err := c.cc.Invoke(ctx, ImageService_GetImage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageServiceServer is the server API for ImageService service.
// All implementations must embed UnimplementedImageServiceServer
// for forward compatibility
type ImageServiceServer interface {
	UploadImage(context.Context, *UploadImageRequest) (*UploadImageResponse, error)
	GetImage(context.Context, *GetImageRequest) (*GetImageResponse, error)
	mustEmbedUnimplementedImageServiceServer()
}

// UnimplementedImageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedImageServiceServer struct {
}

func (UnimplementedImageServiceServer) UploadImage(context.Context, *UploadImageRequest) (*UploadImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadImage not implemented")
}
func (UnimplementedImageServiceServer) GetImage(context.Context, *GetImageRequest) (*GetImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImage not implemented")
}
func (UnimplementedImageServiceServer) mustEmbedUnimplementedImageServiceServer() {}

// UnsafeImageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageServiceServer will
// result in compilation errors.
type UnsafeImageServiceServer interface {
	mustEmbedUnimplementedImageServiceServer()
}

func RegisterImageServiceServer(s grpc.ServiceRegistrar, srv ImageServiceServer) {
	s.RegisterService(&ImageService_ServiceDesc, srv)
}

func _ImageService_UploadImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).UploadImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageService_UploadImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).UploadImage(ctx, req.(*UploadImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageService_GetImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).GetImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageService_GetImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).GetImage(ctx, req.(*GetImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImageService_ServiceDesc is the grpc.ServiceDesc for ImageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server_alpha.image.v1.ImageService",
	HandlerType: (*ImageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadImage",
			Handler:    _ImageService_UploadImage_Handler,
		},
		{
			MethodName: "GetImage",
			Handler:    _ImageService_GetImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server_alpha/image/v1/image.proto",
}
