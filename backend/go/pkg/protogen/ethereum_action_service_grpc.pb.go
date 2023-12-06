// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: ethereum_action_service.proto

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
	EthereumServiceAction_RegisterAddresseWatcher_FullMethodName = "/EthereumServiceAction/RegisterAddresseWatcher"
	EthereumServiceAction_RegisterEventWatcher_FullMethodName    = "/EthereumServiceAction/RegisterEventWatcher"
)

// EthereumServiceActionClient is the client API for EthereumServiceAction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EthereumServiceActionClient interface {
	RegisterAddresseWatcher(ctx context.Context, in *AddressWatcher_Request, opts ...grpc.CallOption) (*Empty, error)
	RegisterEventWatcher(ctx context.Context, in *EventWatcher_Request, opts ...grpc.CallOption) (*Empty, error)
}

type ethereumServiceActionClient struct {
	cc grpc.ClientConnInterface
}

func NewEthereumServiceActionClient(cc grpc.ClientConnInterface) EthereumServiceActionClient {
	return &ethereumServiceActionClient{cc}
}

func (c *ethereumServiceActionClient) RegisterAddresseWatcher(ctx context.Context, in *AddressWatcher_Request, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, EthereumServiceAction_RegisterAddresseWatcher_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ethereumServiceActionClient) RegisterEventWatcher(ctx context.Context, in *EventWatcher_Request, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, EthereumServiceAction_RegisterEventWatcher_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EthereumServiceActionServer is the server API for EthereumServiceAction service.
// All implementations must embed UnimplementedEthereumServiceActionServer
// for forward compatibility
type EthereumServiceActionServer interface {
	RegisterAddresseWatcher(context.Context, *AddressWatcher_Request) (*Empty, error)
	RegisterEventWatcher(context.Context, *EventWatcher_Request) (*Empty, error)
	mustEmbedUnimplementedEthereumServiceActionServer()
}

// UnimplementedEthereumServiceActionServer must be embedded to have forward compatible implementations.
type UnimplementedEthereumServiceActionServer struct {
}

func (UnimplementedEthereumServiceActionServer) RegisterAddresseWatcher(context.Context, *AddressWatcher_Request) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAddresseWatcher not implemented")
}
func (UnimplementedEthereumServiceActionServer) RegisterEventWatcher(context.Context, *EventWatcher_Request) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterEventWatcher not implemented")
}
func (UnimplementedEthereumServiceActionServer) mustEmbedUnimplementedEthereumServiceActionServer() {}

// UnsafeEthereumServiceActionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EthereumServiceActionServer will
// result in compilation errors.
type UnsafeEthereumServiceActionServer interface {
	mustEmbedUnimplementedEthereumServiceActionServer()
}

func RegisterEthereumServiceActionServer(s grpc.ServiceRegistrar, srv EthereumServiceActionServer) {
	s.RegisterService(&EthereumServiceAction_ServiceDesc, srv)
}

func _EthereumServiceAction_RegisterAddresseWatcher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressWatcher_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EthereumServiceActionServer).RegisterAddresseWatcher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EthereumServiceAction_RegisterAddresseWatcher_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EthereumServiceActionServer).RegisterAddresseWatcher(ctx, req.(*AddressWatcher_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _EthereumServiceAction_RegisterEventWatcher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventWatcher_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EthereumServiceActionServer).RegisterEventWatcher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EthereumServiceAction_RegisterEventWatcher_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EthereumServiceActionServer).RegisterEventWatcher(ctx, req.(*EventWatcher_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// EthereumServiceAction_ServiceDesc is the grpc.ServiceDesc for EthereumServiceAction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EthereumServiceAction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "EthereumServiceAction",
	HandlerType: (*EthereumServiceActionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterAddresseWatcher",
			Handler:    _EthereumServiceAction_RegisterAddresseWatcher_Handler,
		},
		{
			MethodName: "RegisterEventWatcher",
			Handler:    _EthereumServiceAction_RegisterEventWatcher_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ethereum_action_service.proto",
}
