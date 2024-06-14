// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: post/post-service.proto

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
	PostService_SearchPosts_FullMethodName = "/serveralpha.post.PostService/SearchPosts"
	PostService_ListPosts_FullMethodName   = "/serveralpha.post.PostService/ListPosts"
	PostService_GetPost_FullMethodName     = "/serveralpha.post.PostService/GetPost"
	PostService_CreatePost_FullMethodName  = "/serveralpha.post.PostService/CreatePost"
	PostService_DeletePost_FullMethodName  = "/serveralpha.post.PostService/DeletePost"
)

// PostServiceClient is the client API for PostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostServiceClient interface {
	SearchPosts(ctx context.Context, in *SearchPostsRequest, opts ...grpc.CallOption) (*SearchPostsResponse, error)
	ListPosts(ctx context.Context, in *ListPostsRequest, opts ...grpc.CallOption) (*SearchPostsResponse, error)
	GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*Post, error)
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*Post, error)
	// rpc UpdatePost(UpdatePostRequest) returns (Post);
	DeletePost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*common.Empty, error)
}

type postServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostServiceClient(cc grpc.ClientConnInterface) PostServiceClient {
	return &postServiceClient{cc}
}

func (c *postServiceClient) SearchPosts(ctx context.Context, in *SearchPostsRequest, opts ...grpc.CallOption) (*SearchPostsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchPostsResponse)
	err := c.cc.Invoke(ctx, PostService_SearchPosts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) ListPosts(ctx context.Context, in *ListPostsRequest, opts ...grpc.CallOption) (*SearchPostsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchPostsResponse)
	err := c.cc.Invoke(ctx, PostService_ListPosts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*Post, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Post)
	err := c.cc.Invoke(ctx, PostService_GetPost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*Post, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Post)
	err := c.cc.Invoke(ctx, PostService_CreatePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) DeletePost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*common.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, PostService_DeletePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServiceServer is the server API for PostService service.
// All implementations must embed UnimplementedPostServiceServer
// for forward compatibility
type PostServiceServer interface {
	SearchPosts(context.Context, *SearchPostsRequest) (*SearchPostsResponse, error)
	ListPosts(context.Context, *ListPostsRequest) (*SearchPostsResponse, error)
	GetPost(context.Context, *GetPostRequest) (*Post, error)
	CreatePost(context.Context, *CreatePostRequest) (*Post, error)
	// rpc UpdatePost(UpdatePostRequest) returns (Post);
	DeletePost(context.Context, *GetPostRequest) (*common.Empty, error)
	mustEmbedUnimplementedPostServiceServer()
}

// UnimplementedPostServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPostServiceServer struct {
}

func (UnimplementedPostServiceServer) SearchPosts(context.Context, *SearchPostsRequest) (*SearchPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPosts not implemented")
}
func (UnimplementedPostServiceServer) ListPosts(context.Context, *ListPostsRequest) (*SearchPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPosts not implemented")
}
func (UnimplementedPostServiceServer) GetPost(context.Context, *GetPostRequest) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}
func (UnimplementedPostServiceServer) CreatePost(context.Context, *CreatePostRequest) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedPostServiceServer) DeletePost(context.Context, *GetPostRequest) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedPostServiceServer) mustEmbedUnimplementedPostServiceServer() {}

// UnsafePostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostServiceServer will
// result in compilation errors.
type UnsafePostServiceServer interface {
	mustEmbedUnimplementedPostServiceServer()
}

func RegisterPostServiceServer(s grpc.ServiceRegistrar, srv PostServiceServer) {
	s.RegisterService(&PostService_ServiceDesc, srv)
}

func _PostService_SearchPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).SearchPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_SearchPosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).SearchPosts(ctx, req.(*SearchPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_ListPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).ListPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_ListPosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).ListPosts(ctx, req.(*ListPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_GetPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetPost(ctx, req.(*GetPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_CreatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_DeletePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).DeletePost(ctx, req.(*GetPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PostService_ServiceDesc is the grpc.ServiceDesc for PostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "serveralpha.post.PostService",
	HandlerType: (*PostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchPosts",
			Handler:    _PostService_SearchPosts_Handler,
		},
		{
			MethodName: "ListPosts",
			Handler:    _PostService_ListPosts_Handler,
		},
		{
			MethodName: "GetPost",
			Handler:    _PostService_GetPost_Handler,
		},
		{
			MethodName: "CreatePost",
			Handler:    _PostService_CreatePost_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _PostService_DeletePost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "post/post-service.proto",
}

const (
	FeedService_GetGlobalFeed_FullMethodName   = "/serveralpha.post.FeedService/GetGlobalFeed"
	FeedService_GetPersonalFeed_FullMethodName = "/serveralpha.post.FeedService/GetPersonalFeed"
)

// FeedServiceClient is the client API for FeedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeedServiceClient interface {
	GetGlobalFeed(ctx context.Context, in *GetFeedRequest, opts ...grpc.CallOption) (*SearchPostsResponse, error)
	GetPersonalFeed(ctx context.Context, in *GetFeedRequest, opts ...grpc.CallOption) (*SearchPostsResponse, error)
}

type feedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedServiceClient(cc grpc.ClientConnInterface) FeedServiceClient {
	return &feedServiceClient{cc}
}

func (c *feedServiceClient) GetGlobalFeed(ctx context.Context, in *GetFeedRequest, opts ...grpc.CallOption) (*SearchPostsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchPostsResponse)
	err := c.cc.Invoke(ctx, FeedService_GetGlobalFeed_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedServiceClient) GetPersonalFeed(ctx context.Context, in *GetFeedRequest, opts ...grpc.CallOption) (*SearchPostsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchPostsResponse)
	err := c.cc.Invoke(ctx, FeedService_GetPersonalFeed_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedServiceServer is the server API for FeedService service.
// All implementations must embed UnimplementedFeedServiceServer
// for forward compatibility
type FeedServiceServer interface {
	GetGlobalFeed(context.Context, *GetFeedRequest) (*SearchPostsResponse, error)
	GetPersonalFeed(context.Context, *GetFeedRequest) (*SearchPostsResponse, error)
	mustEmbedUnimplementedFeedServiceServer()
}

// UnimplementedFeedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFeedServiceServer struct {
}

func (UnimplementedFeedServiceServer) GetGlobalFeed(context.Context, *GetFeedRequest) (*SearchPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGlobalFeed not implemented")
}
func (UnimplementedFeedServiceServer) GetPersonalFeed(context.Context, *GetFeedRequest) (*SearchPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPersonalFeed not implemented")
}
func (UnimplementedFeedServiceServer) mustEmbedUnimplementedFeedServiceServer() {}

// UnsafeFeedServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeedServiceServer will
// result in compilation errors.
type UnsafeFeedServiceServer interface {
	mustEmbedUnimplementedFeedServiceServer()
}

func RegisterFeedServiceServer(s grpc.ServiceRegistrar, srv FeedServiceServer) {
	s.RegisterService(&FeedService_ServiceDesc, srv)
}

func _FeedService_GetGlobalFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).GetGlobalFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeedService_GetGlobalFeed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).GetGlobalFeed(ctx, req.(*GetFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedService_GetPersonalFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).GetPersonalFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeedService_GetPersonalFeed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).GetPersonalFeed(ctx, req.(*GetFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FeedService_ServiceDesc is the grpc.ServiceDesc for FeedService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeedService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "serveralpha.post.FeedService",
	HandlerType: (*FeedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGlobalFeed",
			Handler:    _FeedService_GetGlobalFeed_Handler,
		},
		{
			MethodName: "GetPersonalFeed",
			Handler:    _FeedService_GetPersonalFeed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "post/post-service.proto",
}

const (
	InteractionService_LikePost_FullMethodName      = "/serveralpha.post.InteractionService/LikePost"
	InteractionService_UnlikePost_FullMethodName    = "/serveralpha.post.InteractionService/UnlikePost"
	InteractionService_CreateComment_FullMethodName = "/serveralpha.post.InteractionService/CreateComment"
	InteractionService_ListComments_FullMethodName  = "/serveralpha.post.InteractionService/ListComments"
)

// InteractionServiceClient is the client API for InteractionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InteractionServiceClient interface {
	LikePost(ctx context.Context, in *LikePostRequest, opts ...grpc.CallOption) (*common.Empty, error)
	UnlikePost(ctx context.Context, in *UnlikePostRequest, opts ...grpc.CallOption) (*common.Empty, error)
	CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*Comment, error)
	ListComments(ctx context.Context, in *ListCommentsRequest, opts ...grpc.CallOption) (*ListCommentsResponse, error)
}

type interactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInteractionServiceClient(cc grpc.ClientConnInterface) InteractionServiceClient {
	return &interactionServiceClient{cc}
}

func (c *interactionServiceClient) LikePost(ctx context.Context, in *LikePostRequest, opts ...grpc.CallOption) (*common.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, InteractionService_LikePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionServiceClient) UnlikePost(ctx context.Context, in *UnlikePostRequest, opts ...grpc.CallOption) (*common.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, InteractionService_UnlikePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionServiceClient) CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*Comment, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Comment)
	err := c.cc.Invoke(ctx, InteractionService_CreateComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionServiceClient) ListComments(ctx context.Context, in *ListCommentsRequest, opts ...grpc.CallOption) (*ListCommentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCommentsResponse)
	err := c.cc.Invoke(ctx, InteractionService_ListComments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InteractionServiceServer is the server API for InteractionService service.
// All implementations must embed UnimplementedInteractionServiceServer
// for forward compatibility
type InteractionServiceServer interface {
	LikePost(context.Context, *LikePostRequest) (*common.Empty, error)
	UnlikePost(context.Context, *UnlikePostRequest) (*common.Empty, error)
	CreateComment(context.Context, *CreateCommentRequest) (*Comment, error)
	ListComments(context.Context, *ListCommentsRequest) (*ListCommentsResponse, error)
	mustEmbedUnimplementedInteractionServiceServer()
}

// UnimplementedInteractionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInteractionServiceServer struct {
}

func (UnimplementedInteractionServiceServer) LikePost(context.Context, *LikePostRequest) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikePost not implemented")
}
func (UnimplementedInteractionServiceServer) UnlikePost(context.Context, *UnlikePostRequest) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlikePost not implemented")
}
func (UnimplementedInteractionServiceServer) CreateComment(context.Context, *CreateCommentRequest) (*Comment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedInteractionServiceServer) ListComments(context.Context, *ListCommentsRequest) (*ListCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListComments not implemented")
}
func (UnimplementedInteractionServiceServer) mustEmbedUnimplementedInteractionServiceServer() {}

// UnsafeInteractionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InteractionServiceServer will
// result in compilation errors.
type UnsafeInteractionServiceServer interface {
	mustEmbedUnimplementedInteractionServiceServer()
}

func RegisterInteractionServiceServer(s grpc.ServiceRegistrar, srv InteractionServiceServer) {
	s.RegisterService(&InteractionService_ServiceDesc, srv)
}

func _InteractionService_LikePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionServiceServer).LikePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractionService_LikePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionServiceServer).LikePost(ctx, req.(*LikePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionService_UnlikePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnlikePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionServiceServer).UnlikePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractionService_UnlikePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionServiceServer).UnlikePost(ctx, req.(*UnlikePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionService_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionServiceServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractionService_CreateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionServiceServer).CreateComment(ctx, req.(*CreateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionService_ListComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionServiceServer).ListComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractionService_ListComments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionServiceServer).ListComments(ctx, req.(*ListCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InteractionService_ServiceDesc is the grpc.ServiceDesc for InteractionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InteractionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "serveralpha.post.InteractionService",
	HandlerType: (*InteractionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LikePost",
			Handler:    _InteractionService_LikePost_Handler,
		},
		{
			MethodName: "UnlikePost",
			Handler:    _InteractionService_UnlikePost_Handler,
		},
		{
			MethodName: "CreateComment",
			Handler:    _InteractionService_CreateComment_Handler,
		},
		{
			MethodName: "ListComments",
			Handler:    _InteractionService_ListComments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "post/post-service.proto",
}
