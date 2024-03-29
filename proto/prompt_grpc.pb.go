//nolint

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: prompt.proto

package proto

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

const (
	PromptSearch_Echo_FullMethodName                   = "/proto.PromptSearch/Echo"
	PromptSearch_SearchImagesFromPrompt_FullMethodName = "/proto.PromptSearch/SearchImagesFromPrompt"
	PromptSearch_SearchStylesFromPrompt_FullMethodName = "/proto.PromptSearch/SearchStylesFromPrompt"
	PromptSearch_SearchUsersFromText_FullMethodName    = "/proto.PromptSearch/SearchUsersFromText"
	PromptSearch_SearchPromptHistory_FullMethodName    = "/proto.PromptSearch/SearchPromptHistory"
	PromptSearch_HidePromptFromUser_FullMethodName     = "/proto.PromptSearch/HidePromptFromUser"
	PromptSearch_GetFeaturesForPost_FullMethodName     = "/proto.PromptSearch/GetFeaturesForPost"
	PromptSearch_SearchPostsFromPrompt_FullMethodName  = "/proto.PromptSearch/SearchPostsFromPrompt"
)

// PromptSearchClient is the client API for PromptSearch service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PromptSearchClient interface {
	// Echo back the prompt request
	Echo(ctx context.Context, in *PromptRequest, opts ...grpc.CallOption) (*PromptRequest, error)
	// Search images from prompt by vector similarity.
	SearchImagesFromPrompt(ctx context.Context, in *PromptRequest, opts ...grpc.CallOption) (*Prompts, error)
	// Search styles from prompt by text similarity.
	SearchStylesFromPrompt(ctx context.Context, in *PromptRequest, opts ...grpc.CallOption) (*Prompts, error)
	// Search usernames and names
	SearchUsersFromText(ctx context.Context, in *PromptRequest, opts ...grpc.CallOption) (*CreatorProfiles, error)
	// Search prompts history
	SearchPromptHistory(ctx context.Context, in *PromptRequest, opts ...grpc.CallOption) (*Prompts, error)
	// Hide prompt from user's history
	HidePromptFromUser(ctx context.Context, in *PromptRequest, opts ...grpc.CallOption) (*Prompt, error)
	// Get feature array for post
	GetFeaturesForPost(ctx context.Context, in *PromptRequest, opts ...grpc.CallOption) (*Features, error)
	// Get posts from prompt search
	SearchPostsFromPrompt(ctx context.Context, in *PromptRequest, opts ...grpc.CallOption) (*ResponseFeed, error)
}

type promptSearchClient struct {
	cc grpc.ClientConnInterface
}

func NewPromptSearchClient(cc grpc.ClientConnInterface) PromptSearchClient {
	return &promptSearchClient{cc}
}

func (c *promptSearchClient) Echo(ctx context.Context, in *PromptRequest, opts ...grpc.CallOption) (*PromptRequest, error) {
	out := new(PromptRequest)
	err := c.cc.Invoke(ctx, PromptSearch_Echo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promptSearchClient) SearchImagesFromPrompt(ctx context.Context, in *PromptRequest, opts ...grpc.CallOption) (*Prompts, error) {
	out := new(Prompts)
	err := c.cc.Invoke(ctx, PromptSearch_SearchImagesFromPrompt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promptSearchClient) SearchStylesFromPrompt(ctx context.Context, in *PromptRequest, opts ...grpc.CallOption) (*Prompts, error) {
	out := new(Prompts)
	err := c.cc.Invoke(ctx, PromptSearch_SearchStylesFromPrompt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promptSearchClient) SearchUsersFromText(ctx context.Context, in *PromptRequest, opts ...grpc.CallOption) (*CreatorProfiles, error) {
	out := new(CreatorProfiles)
	err := c.cc.Invoke(ctx, PromptSearch_SearchUsersFromText_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promptSearchClient) SearchPromptHistory(ctx context.Context, in *PromptRequest, opts ...grpc.CallOption) (*Prompts, error) {
	out := new(Prompts)
	err := c.cc.Invoke(ctx, PromptSearch_SearchPromptHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promptSearchClient) HidePromptFromUser(ctx context.Context, in *PromptRequest, opts ...grpc.CallOption) (*Prompt, error) {
	out := new(Prompt)
	err := c.cc.Invoke(ctx, PromptSearch_HidePromptFromUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promptSearchClient) GetFeaturesForPost(ctx context.Context, in *PromptRequest, opts ...grpc.CallOption) (*Features, error) {
	out := new(Features)
	err := c.cc.Invoke(ctx, PromptSearch_GetFeaturesForPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promptSearchClient) SearchPostsFromPrompt(ctx context.Context, in *PromptRequest, opts ...grpc.CallOption) (*ResponseFeed, error) {
	out := new(ResponseFeed)
	err := c.cc.Invoke(ctx, PromptSearch_SearchPostsFromPrompt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PromptSearchServer is the server API for PromptSearch service.
// All implementations must embed UnimplementedPromptSearchServer
// for forward compatibility
type PromptSearchServer interface {
	// Echo back the prompt request
	Echo(context.Context, *PromptRequest) (*PromptRequest, error)
	// Search images from prompt by vector similarity.
	SearchImagesFromPrompt(context.Context, *PromptRequest) (*Prompts, error)
	// Search styles from prompt by text similarity.
	SearchStylesFromPrompt(context.Context, *PromptRequest) (*Prompts, error)
	// Search usernames and names
	SearchUsersFromText(context.Context, *PromptRequest) (*CreatorProfiles, error)
	// Search prompts history
	SearchPromptHistory(context.Context, *PromptRequest) (*Prompts, error)
	// Hide prompt from user's history
	HidePromptFromUser(context.Context, *PromptRequest) (*Prompt, error)
	// Get feature array for post
	GetFeaturesForPost(context.Context, *PromptRequest) (*Features, error)
	// Get posts from prompt search
	SearchPostsFromPrompt(context.Context, *PromptRequest) (*ResponseFeed, error)
	mustEmbedUnimplementedPromptSearchServer()
}

// UnimplementedPromptSearchServer must be embedded to have forward compatible implementations.
type UnimplementedPromptSearchServer struct {
}

func (UnimplementedPromptSearchServer) Echo(context.Context, *PromptRequest) (*PromptRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedPromptSearchServer) SearchImagesFromPrompt(context.Context, *PromptRequest) (*Prompts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchImagesFromPrompt not implemented")
}
func (UnimplementedPromptSearchServer) SearchStylesFromPrompt(context.Context, *PromptRequest) (*Prompts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchStylesFromPrompt not implemented")
}
func (UnimplementedPromptSearchServer) SearchUsersFromText(context.Context, *PromptRequest) (*CreatorProfiles, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUsersFromText not implemented")
}
func (UnimplementedPromptSearchServer) SearchPromptHistory(context.Context, *PromptRequest) (*Prompts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPromptHistory not implemented")
}
func (UnimplementedPromptSearchServer) HidePromptFromUser(context.Context, *PromptRequest) (*Prompt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HidePromptFromUser not implemented")
}
func (UnimplementedPromptSearchServer) GetFeaturesForPost(context.Context, *PromptRequest) (*Features, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeaturesForPost not implemented")
}
func (UnimplementedPromptSearchServer) SearchPostsFromPrompt(context.Context, *PromptRequest) (*ResponseFeed, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPostsFromPrompt not implemented")
}
func (UnimplementedPromptSearchServer) mustEmbedUnimplementedPromptSearchServer() {}

// UnsafePromptSearchServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PromptSearchServer will
// result in compilation errors.
type UnsafePromptSearchServer interface {
	mustEmbedUnimplementedPromptSearchServer()
}

func RegisterPromptSearchServer(s grpc.ServiceRegistrar, srv PromptSearchServer) {
	s.RegisterService(&PromptSearch_ServiceDesc, srv)
}

func _PromptSearch_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromptSearchServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PromptSearch_Echo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromptSearchServer).Echo(ctx, req.(*PromptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromptSearch_SearchImagesFromPrompt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromptSearchServer).SearchImagesFromPrompt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PromptSearch_SearchImagesFromPrompt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromptSearchServer).SearchImagesFromPrompt(ctx, req.(*PromptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromptSearch_SearchStylesFromPrompt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromptSearchServer).SearchStylesFromPrompt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PromptSearch_SearchStylesFromPrompt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromptSearchServer).SearchStylesFromPrompt(ctx, req.(*PromptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromptSearch_SearchUsersFromText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromptSearchServer).SearchUsersFromText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PromptSearch_SearchUsersFromText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromptSearchServer).SearchUsersFromText(ctx, req.(*PromptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromptSearch_SearchPromptHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromptSearchServer).SearchPromptHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PromptSearch_SearchPromptHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromptSearchServer).SearchPromptHistory(ctx, req.(*PromptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromptSearch_HidePromptFromUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromptSearchServer).HidePromptFromUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PromptSearch_HidePromptFromUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromptSearchServer).HidePromptFromUser(ctx, req.(*PromptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromptSearch_GetFeaturesForPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromptSearchServer).GetFeaturesForPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PromptSearch_GetFeaturesForPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromptSearchServer).GetFeaturesForPost(ctx, req.(*PromptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromptSearch_SearchPostsFromPrompt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromptSearchServer).SearchPostsFromPrompt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PromptSearch_SearchPostsFromPrompt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromptSearchServer).SearchPostsFromPrompt(ctx, req.(*PromptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PromptSearch_ServiceDesc is the grpc.ServiceDesc for PromptSearch service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PromptSearch_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PromptSearch",
	HandlerType: (*PromptSearchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _PromptSearch_Echo_Handler,
		},
		{
			MethodName: "SearchImagesFromPrompt",
			Handler:    _PromptSearch_SearchImagesFromPrompt_Handler,
		},
		{
			MethodName: "SearchStylesFromPrompt",
			Handler:    _PromptSearch_SearchStylesFromPrompt_Handler,
		},
		{
			MethodName: "SearchUsersFromText",
			Handler:    _PromptSearch_SearchUsersFromText_Handler,
		},
		{
			MethodName: "SearchPromptHistory",
			Handler:    _PromptSearch_SearchPromptHistory_Handler,
		},
		{
			MethodName: "HidePromptFromUser",
			Handler:    _PromptSearch_HidePromptFromUser_Handler,
		},
		{
			MethodName: "GetFeaturesForPost",
			Handler:    _PromptSearch_GetFeaturesForPost_Handler,
		},
		{
			MethodName: "SearchPostsFromPrompt",
			Handler:    _PromptSearch_SearchPostsFromPrompt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "prompt.proto",
}
