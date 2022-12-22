// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: cosmos/crisis/v1/tx.proto

package crisisv1

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

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// AuthorizeCircuitBreaker allows a super-admin to grant (or revoke) another
	// account's circuit breaker permissions.
	AuthorizeCircuitBreaker(ctx context.Context, in *MsgAuthorizeCircuitBreaker, opts ...grpc.CallOption) (*MsgAuthorizeCircuitBreakerResponse, error)
	// TripCircuitBreaker pauses processing of Msg's in the state machine.
	TripCircuitBreaker(ctx context.Context, in *MsgTripCircuitBreaker, opts ...grpc.CallOption) (*MsgTripCircuitBreakerResponse, error)
	// ResetCircuitBreaker resumes processing of Msg's in the state machine that
	// have been been paused using TripCircuitBreaker.
	ResetCircuitBreaker(ctx context.Context, in *MsgResetCircuitBreaker, opts ...grpc.CallOption) (*MsgResetCircuitBreakerResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) AuthorizeCircuitBreaker(ctx context.Context, in *MsgAuthorizeCircuitBreaker, opts ...grpc.CallOption) (*MsgAuthorizeCircuitBreakerResponse, error) {
	out := new(MsgAuthorizeCircuitBreakerResponse)
	err := c.cc.Invoke(ctx, "/cosmos.crisis.v1.Msg/AuthorizeCircuitBreaker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) TripCircuitBreaker(ctx context.Context, in *MsgTripCircuitBreaker, opts ...grpc.CallOption) (*MsgTripCircuitBreakerResponse, error) {
	out := new(MsgTripCircuitBreakerResponse)
	err := c.cc.Invoke(ctx, "/cosmos.crisis.v1.Msg/TripCircuitBreaker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ResetCircuitBreaker(ctx context.Context, in *MsgResetCircuitBreaker, opts ...grpc.CallOption) (*MsgResetCircuitBreakerResponse, error) {
	out := new(MsgResetCircuitBreakerResponse)
	err := c.cc.Invoke(ctx, "/cosmos.crisis.v1.Msg/ResetCircuitBreaker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// AuthorizeCircuitBreaker allows a super-admin to grant (or revoke) another
	// account's circuit breaker permissions.
	AuthorizeCircuitBreaker(context.Context, *MsgAuthorizeCircuitBreaker) (*MsgAuthorizeCircuitBreakerResponse, error)
	// TripCircuitBreaker pauses processing of Msg's in the state machine.
	TripCircuitBreaker(context.Context, *MsgTripCircuitBreaker) (*MsgTripCircuitBreakerResponse, error)
	// ResetCircuitBreaker resumes processing of Msg's in the state machine that
	// have been been paused using TripCircuitBreaker.
	ResetCircuitBreaker(context.Context, *MsgResetCircuitBreaker) (*MsgResetCircuitBreakerResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) AuthorizeCircuitBreaker(context.Context, *MsgAuthorizeCircuitBreaker) (*MsgAuthorizeCircuitBreakerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizeCircuitBreaker not implemented")
}
func (UnimplementedMsgServer) TripCircuitBreaker(context.Context, *MsgTripCircuitBreaker) (*MsgTripCircuitBreakerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TripCircuitBreaker not implemented")
}
func (UnimplementedMsgServer) ResetCircuitBreaker(context.Context, *MsgResetCircuitBreaker) (*MsgResetCircuitBreakerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetCircuitBreaker not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_AuthorizeCircuitBreaker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAuthorizeCircuitBreaker)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AuthorizeCircuitBreaker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.crisis.v1.Msg/AuthorizeCircuitBreaker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AuthorizeCircuitBreaker(ctx, req.(*MsgAuthorizeCircuitBreaker))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_TripCircuitBreaker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTripCircuitBreaker)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).TripCircuitBreaker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.crisis.v1.Msg/TripCircuitBreaker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).TripCircuitBreaker(ctx, req.(*MsgTripCircuitBreaker))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ResetCircuitBreaker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgResetCircuitBreaker)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ResetCircuitBreaker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.crisis.v1.Msg/ResetCircuitBreaker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ResetCircuitBreaker(ctx, req.(*MsgResetCircuitBreaker))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.crisis.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthorizeCircuitBreaker",
			Handler:    _Msg_AuthorizeCircuitBreaker_Handler,
		},
		{
			MethodName: "TripCircuitBreaker",
			Handler:    _Msg_TripCircuitBreaker_Handler,
		},
		{
			MethodName: "ResetCircuitBreaker",
			Handler:    _Msg_ResetCircuitBreaker_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/crisis/v1/tx.proto",
}
