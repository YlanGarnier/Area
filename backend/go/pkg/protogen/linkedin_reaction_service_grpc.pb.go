// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: linkedin_reaction_service.proto

package protogen

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
	LinkedinServiceReaction_CreateDefaultPost_FullMethodName = "/LinkedinServiceReaction/createDefaultPost"
	LinkedinServiceReaction_CreatePost_FullMethodName        = "/LinkedinServiceReaction/createPost"
)

// LinkedinServiceReactionClient is the client API for LinkedinServiceReaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LinkedinServiceReactionClient interface {
	CreateDefaultPost(ctx context.Context, in *Format_OnlyTitle, opts ...grpc.CallOption) (*Empty, error)
	CreatePost(ctx context.Context, in *Format_OnlyTitle, opts ...grpc.CallOption) (*Empty, error)
}

type linkedinServiceReactionClient struct {
	cc grpc.ClientConnInterface
}

func NewLinkedinServiceReactionClient(cc grpc.ClientConnInterface) LinkedinServiceReactionClient {
	return &linkedinServiceReactionClient{cc}
}

func (c *linkedinServiceReactionClient) CreateDefaultPost(ctx context.Context, in *Format_OnlyTitle, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, LinkedinServiceReaction_CreateDefaultPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkedinServiceReactionClient) CreatePost(ctx context.Context, in *Format_OnlyTitle, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, LinkedinServiceReaction_CreatePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LinkedinServiceReactionServer is the server API for LinkedinServiceReaction service.
// All implementations must embed UnimplementedLinkedinServiceReactionServer
// for forward compatibility
type LinkedinServiceReactionServer interface {
	CreateDefaultPost(context.Context, *Format_OnlyTitle) (*Empty, error)
	CreatePost(context.Context, *Format_OnlyTitle) (*Empty, error)
	mustEmbedUnimplementedLinkedinServiceReactionServer()
}

// UnimplementedLinkedinServiceReactionServer must be embedded to have forward compatible implementations.
type UnimplementedLinkedinServiceReactionServer struct {
}

func (UnimplementedLinkedinServiceReactionServer) CreateDefaultPost(context.Context, *Format_OnlyTitle) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDefaultPost not implemented")
}
func (UnimplementedLinkedinServiceReactionServer) CreatePost(context.Context, *Format_OnlyTitle) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedLinkedinServiceReactionServer) mustEmbedUnimplementedLinkedinServiceReactionServer() {
}

// UnsafeLinkedinServiceReactionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LinkedinServiceReactionServer will
// result in compilation errors.
type UnsafeLinkedinServiceReactionServer interface {
	mustEmbedUnimplementedLinkedinServiceReactionServer()
}

func RegisterLinkedinServiceReactionServer(s grpc.ServiceRegistrar, srv LinkedinServiceReactionServer) {
	s.RegisterService(&LinkedinServiceReaction_ServiceDesc, srv)
}

func _LinkedinServiceReaction_CreateDefaultPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Format_OnlyTitle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkedinServiceReactionServer).CreateDefaultPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkedinServiceReaction_CreateDefaultPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkedinServiceReactionServer).CreateDefaultPost(ctx, req.(*Format_OnlyTitle))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkedinServiceReaction_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Format_OnlyTitle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkedinServiceReactionServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkedinServiceReaction_CreatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkedinServiceReactionServer).CreatePost(ctx, req.(*Format_OnlyTitle))
	}
	return interceptor(ctx, in, info, handler)
}

// LinkedinServiceReaction_ServiceDesc is the grpc.ServiceDesc for LinkedinServiceReaction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LinkedinServiceReaction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "LinkedinServiceReaction",
	HandlerType: (*LinkedinServiceReactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createDefaultPost",
			Handler:    _LinkedinServiceReaction_CreateDefaultPost_Handler,
		},
		{
			MethodName: "createPost",
			Handler:    _LinkedinServiceReaction_CreatePost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "linkedin_reaction_service.proto",
}
