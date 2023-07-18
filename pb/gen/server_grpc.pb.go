// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: server.proto

package gen

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

// ChainServiceClient is the client API for ChainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChainServiceClient interface {
	CreateChainAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountRespon, error)
	SendToken(ctx context.Context, in *SendTokenRequest, opts ...grpc.CallOption) (*SendTokenRespon, error)
	CreateTransfer(ctx context.Context, in *CreateTransferRequest, opts ...grpc.CallOption) (*CreateTransferRespon, error)
	SignTransfer(ctx context.Context, in *SignTransferRequest, opts ...grpc.CallOption) (*SignTransferRespon, error)
	Broadcast(ctx context.Context, in *BroadcastTransferRequest, opts ...grpc.CallOption) (*BroadcastTransferRespon, error)
	BalanceByUid(ctx context.Context, in *BalanceByUidRequest, opts ...grpc.CallOption) (*BalanceByUidRespon, error)
	BalanceByAddress(ctx context.Context, in *BalanceByAddressRequest, opts ...grpc.CallOption) (*BalanceByAddressRespon, error)
	SetNotifyUrl(ctx context.Context, in *SetNotifyUrlRequest, opts ...grpc.CallOption) (*SetNotifyUrlRespon, error)
}

type chainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChainServiceClient(cc grpc.ClientConnInterface) ChainServiceClient {
	return &chainServiceClient{cc}
}

func (c *chainServiceClient) CreateChainAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountRespon, error) {
	out := new(CreateAccountRespon)
	err := c.cc.Invoke(ctx, "/proto.ChainService/CreateChainAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainServiceClient) SendToken(ctx context.Context, in *SendTokenRequest, opts ...grpc.CallOption) (*SendTokenRespon, error) {
	out := new(SendTokenRespon)
	err := c.cc.Invoke(ctx, "/proto.ChainService/SendToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainServiceClient) CreateTransfer(ctx context.Context, in *CreateTransferRequest, opts ...grpc.CallOption) (*CreateTransferRespon, error) {
	out := new(CreateTransferRespon)
	err := c.cc.Invoke(ctx, "/proto.ChainService/CreateTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainServiceClient) SignTransfer(ctx context.Context, in *SignTransferRequest, opts ...grpc.CallOption) (*SignTransferRespon, error) {
	out := new(SignTransferRespon)
	err := c.cc.Invoke(ctx, "/proto.ChainService/SignTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainServiceClient) Broadcast(ctx context.Context, in *BroadcastTransferRequest, opts ...grpc.CallOption) (*BroadcastTransferRespon, error) {
	out := new(BroadcastTransferRespon)
	err := c.cc.Invoke(ctx, "/proto.ChainService/Broadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainServiceClient) BalanceByUid(ctx context.Context, in *BalanceByUidRequest, opts ...grpc.CallOption) (*BalanceByUidRespon, error) {
	out := new(BalanceByUidRespon)
	err := c.cc.Invoke(ctx, "/proto.ChainService/BalanceByUid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainServiceClient) BalanceByAddress(ctx context.Context, in *BalanceByAddressRequest, opts ...grpc.CallOption) (*BalanceByAddressRespon, error) {
	out := new(BalanceByAddressRespon)
	err := c.cc.Invoke(ctx, "/proto.ChainService/BalanceByAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainServiceClient) SetNotifyUrl(ctx context.Context, in *SetNotifyUrlRequest, opts ...grpc.CallOption) (*SetNotifyUrlRespon, error) {
	out := new(SetNotifyUrlRespon)
	err := c.cc.Invoke(ctx, "/proto.ChainService/SetNotifyUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChainServiceServer is the server API for ChainService service.
// All implementations must embed UnimplementedChainServiceServer
// for forward compatibility
type ChainServiceServer interface {
	CreateChainAccount(context.Context, *CreateAccountRequest) (*CreateAccountRespon, error)
	SendToken(context.Context, *SendTokenRequest) (*SendTokenRespon, error)
	CreateTransfer(context.Context, *CreateTransferRequest) (*CreateTransferRespon, error)
	SignTransfer(context.Context, *SignTransferRequest) (*SignTransferRespon, error)
	Broadcast(context.Context, *BroadcastTransferRequest) (*BroadcastTransferRespon, error)
	BalanceByUid(context.Context, *BalanceByUidRequest) (*BalanceByUidRespon, error)
	BalanceByAddress(context.Context, *BalanceByAddressRequest) (*BalanceByAddressRespon, error)
	SetNotifyUrl(context.Context, *SetNotifyUrlRequest) (*SetNotifyUrlRespon, error)
	mustEmbedUnimplementedChainServiceServer()
}

// UnimplementedChainServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChainServiceServer struct {
}

func (UnimplementedChainServiceServer) CreateChainAccount(context.Context, *CreateAccountRequest) (*CreateAccountRespon, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChainAccount not implemented")
}
func (UnimplementedChainServiceServer) SendToken(context.Context, *SendTokenRequest) (*SendTokenRespon, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendToken not implemented")
}
func (UnimplementedChainServiceServer) CreateTransfer(context.Context, *CreateTransferRequest) (*CreateTransferRespon, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransfer not implemented")
}
func (UnimplementedChainServiceServer) SignTransfer(context.Context, *SignTransferRequest) (*SignTransferRespon, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignTransfer not implemented")
}
func (UnimplementedChainServiceServer) Broadcast(context.Context, *BroadcastTransferRequest) (*BroadcastTransferRespon, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (UnimplementedChainServiceServer) BalanceByUid(context.Context, *BalanceByUidRequest) (*BalanceByUidRespon, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BalanceByUid not implemented")
}
func (UnimplementedChainServiceServer) BalanceByAddress(context.Context, *BalanceByAddressRequest) (*BalanceByAddressRespon, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BalanceByAddress not implemented")
}
func (UnimplementedChainServiceServer) SetNotifyUrl(context.Context, *SetNotifyUrlRequest) (*SetNotifyUrlRespon, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetNotifyUrl not implemented")
}
func (UnimplementedChainServiceServer) mustEmbedUnimplementedChainServiceServer() {}

// UnsafeChainServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChainServiceServer will
// result in compilation errors.
type UnsafeChainServiceServer interface {
	mustEmbedUnimplementedChainServiceServer()
}

func RegisterChainServiceServer(s grpc.ServiceRegistrar, srv ChainServiceServer) {
	s.RegisterService(&ChainService_ServiceDesc, srv)
}

func _ChainService_CreateChainAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServiceServer).CreateChainAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ChainService/CreateChainAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServiceServer).CreateChainAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainService_SendToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServiceServer).SendToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ChainService/SendToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServiceServer).SendToken(ctx, req.(*SendTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainService_CreateTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServiceServer).CreateTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ChainService/CreateTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServiceServer).CreateTransfer(ctx, req.(*CreateTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainService_SignTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServiceServer).SignTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ChainService/SignTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServiceServer).SignTransfer(ctx, req.(*SignTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainService_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServiceServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ChainService/Broadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServiceServer).Broadcast(ctx, req.(*BroadcastTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainService_BalanceByUid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceByUidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServiceServer).BalanceByUid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ChainService/BalanceByUid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServiceServer).BalanceByUid(ctx, req.(*BalanceByUidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainService_BalanceByAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceByAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServiceServer).BalanceByAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ChainService/BalanceByAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServiceServer).BalanceByAddress(ctx, req.(*BalanceByAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainService_SetNotifyUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetNotifyUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServiceServer).SetNotifyUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ChainService/SetNotifyUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServiceServer).SetNotifyUrl(ctx, req.(*SetNotifyUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChainService_ServiceDesc is the grpc.ServiceDesc for ChainService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChainService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ChainService",
	HandlerType: (*ChainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChainAccount",
			Handler:    _ChainService_CreateChainAccount_Handler,
		},
		{
			MethodName: "SendToken",
			Handler:    _ChainService_SendToken_Handler,
		},
		{
			MethodName: "CreateTransfer",
			Handler:    _ChainService_CreateTransfer_Handler,
		},
		{
			MethodName: "SignTransfer",
			Handler:    _ChainService_SignTransfer_Handler,
		},
		{
			MethodName: "Broadcast",
			Handler:    _ChainService_Broadcast_Handler,
		},
		{
			MethodName: "BalanceByUid",
			Handler:    _ChainService_BalanceByUid_Handler,
		},
		{
			MethodName: "BalanceByAddress",
			Handler:    _ChainService_BalanceByAddress_Handler,
		},
		{
			MethodName: "SetNotifyUrl",
			Handler:    _ChainService_SetNotifyUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}
