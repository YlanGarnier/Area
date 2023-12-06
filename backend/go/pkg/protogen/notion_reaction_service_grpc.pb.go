// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: notion_reaction_service.proto

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
	NotionServiceReaction_CreateDefaultComment_FullMethodName = "/NotionServiceReaction/createDefaultComment"
	NotionServiceReaction_CreateComment_FullMethodName        = "/NotionServiceReaction/createComment"
	NotionServiceReaction_CreateDefaultPage_FullMethodName    = "/NotionServiceReaction/createDefaultPage"
	NotionServiceReaction_CreatePage_FullMethodName           = "/NotionServiceReaction/createPage"
	NotionServiceReaction_CreateBlock_FullMethodName          = "/NotionServiceReaction/createBlock"
)

// NotionServiceReactionClient is the client API for NotionServiceReaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotionServiceReactionClient interface {
	CreateDefaultComment(ctx context.Context, in *Format_OnlyTitle, opts ...grpc.CallOption) (*Empty, error)
	CreateComment(ctx context.Context, in *Format_GHIncidentReport, opts ...grpc.CallOption) (*Empty, error)
	CreateDefaultPage(ctx context.Context, in *Format_OnlyTitle, opts ...grpc.CallOption) (*Empty, error)
	CreatePage(ctx context.Context, in *Format_GHIncidentReport, opts ...grpc.CallOption) (*Empty, error)
	CreateBlock(ctx context.Context, in *Format_GHIncidentReport, opts ...grpc.CallOption) (*Empty, error)
}

type notionServiceReactionClient struct {
	cc grpc.ClientConnInterface
}

func NewNotionServiceReactionClient(cc grpc.ClientConnInterface) NotionServiceReactionClient {
	return &notionServiceReactionClient{cc}
}

func (c *notionServiceReactionClient) CreateDefaultComment(ctx context.Context, in *Format_OnlyTitle, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, NotionServiceReaction_CreateDefaultComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notionServiceReactionClient) CreateComment(ctx context.Context, in *Format_GHIncidentReport, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, NotionServiceReaction_CreateComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notionServiceReactionClient) CreateDefaultPage(ctx context.Context, in *Format_OnlyTitle, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, NotionServiceReaction_CreateDefaultPage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notionServiceReactionClient) CreatePage(ctx context.Context, in *Format_GHIncidentReport, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, NotionServiceReaction_CreatePage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notionServiceReactionClient) CreateBlock(ctx context.Context, in *Format_GHIncidentReport, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, NotionServiceReaction_CreateBlock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotionServiceReactionServer is the server API for NotionServiceReaction service.
// All implementations must embed UnimplementedNotionServiceReactionServer
// for forward compatibility
type NotionServiceReactionServer interface {
	CreateDefaultComment(context.Context, *Format_OnlyTitle) (*Empty, error)
	CreateComment(context.Context, *Format_GHIncidentReport) (*Empty, error)
	CreateDefaultPage(context.Context, *Format_OnlyTitle) (*Empty, error)
	CreatePage(context.Context, *Format_GHIncidentReport) (*Empty, error)
	CreateBlock(context.Context, *Format_GHIncidentReport) (*Empty, error)
	mustEmbedUnimplementedNotionServiceReactionServer()
}

// UnimplementedNotionServiceReactionServer must be embedded to have forward compatible implementations.
type UnimplementedNotionServiceReactionServer struct {
}

func (UnimplementedNotionServiceReactionServer) CreateDefaultComment(context.Context, *Format_OnlyTitle) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDefaultComment not implemented")
}
func (UnimplementedNotionServiceReactionServer) CreateComment(context.Context, *Format_GHIncidentReport) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedNotionServiceReactionServer) CreateDefaultPage(context.Context, *Format_OnlyTitle) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDefaultPage not implemented")
}
func (UnimplementedNotionServiceReactionServer) CreatePage(context.Context, *Format_GHIncidentReport) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePage not implemented")
}
func (UnimplementedNotionServiceReactionServer) CreateBlock(context.Context, *Format_GHIncidentReport) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlock not implemented")
}
func (UnimplementedNotionServiceReactionServer) mustEmbedUnimplementedNotionServiceReactionServer() {}

// UnsafeNotionServiceReactionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotionServiceReactionServer will
// result in compilation errors.
type UnsafeNotionServiceReactionServer interface {
	mustEmbedUnimplementedNotionServiceReactionServer()
}

func RegisterNotionServiceReactionServer(s grpc.ServiceRegistrar, srv NotionServiceReactionServer) {
	s.RegisterService(&NotionServiceReaction_ServiceDesc, srv)
}

func _NotionServiceReaction_CreateDefaultComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Format_OnlyTitle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotionServiceReactionServer).CreateDefaultComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotionServiceReaction_CreateDefaultComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotionServiceReactionServer).CreateDefaultComment(ctx, req.(*Format_OnlyTitle))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotionServiceReaction_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Format_GHIncidentReport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotionServiceReactionServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotionServiceReaction_CreateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotionServiceReactionServer).CreateComment(ctx, req.(*Format_GHIncidentReport))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotionServiceReaction_CreateDefaultPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Format_OnlyTitle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotionServiceReactionServer).CreateDefaultPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotionServiceReaction_CreateDefaultPage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotionServiceReactionServer).CreateDefaultPage(ctx, req.(*Format_OnlyTitle))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotionServiceReaction_CreatePage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Format_GHIncidentReport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotionServiceReactionServer).CreatePage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotionServiceReaction_CreatePage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotionServiceReactionServer).CreatePage(ctx, req.(*Format_GHIncidentReport))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotionServiceReaction_CreateBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Format_GHIncidentReport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotionServiceReactionServer).CreateBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotionServiceReaction_CreateBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotionServiceReactionServer).CreateBlock(ctx, req.(*Format_GHIncidentReport))
	}
	return interceptor(ctx, in, info, handler)
}

// NotionServiceReaction_ServiceDesc is the grpc.ServiceDesc for NotionServiceReaction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotionServiceReaction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "NotionServiceReaction",
	HandlerType: (*NotionServiceReactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createDefaultComment",
			Handler:    _NotionServiceReaction_CreateDefaultComment_Handler,
		},
		{
			MethodName: "createComment",
			Handler:    _NotionServiceReaction_CreateComment_Handler,
		},
		{
			MethodName: "createDefaultPage",
			Handler:    _NotionServiceReaction_CreateDefaultPage_Handler,
		},
		{
			MethodName: "createPage",
			Handler:    _NotionServiceReaction_CreatePage_Handler,
		},
		{
			MethodName: "createBlock",
			Handler:    _NotionServiceReaction_CreateBlock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notion_reaction_service.proto",
}