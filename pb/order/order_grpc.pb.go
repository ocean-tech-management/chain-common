// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: order.proto

// 定义包名

package order

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
	OrderServer_Transfer_FullMethodName             = "/order.OrderServer/Transfer"
	OrderServer_GetAssetsByAccountID_FullMethodName = "/order.OrderServer/GetAssetsByAccountID"
	OrderServer_GetAssetsHistory_FullMethodName     = "/order.OrderServer/GetAssetsHistory"
	OrderServer_GetOrderHistory_FullMethodName      = "/order.OrderServer/GetOrderHistory"
	OrderServer_WithDraw_FullMethodName             = "/order.OrderServer/WithDraw"
)

// OrderServerClient is the client API for OrderServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServerClient interface {
	// 执行转账
	Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferReply, error)
	// 1. 账户资产查询
	GetAssetsByAccountID(ctx context.Context, in *GetAssetsByAccountIDRequest, opts ...grpc.CallOption) (*GetAssetsByAccountIDReply, error)
	// 资产变动记录
	GetAssetsHistory(ctx context.Context, in *GetAssetsHistoryRequest, opts ...grpc.CallOption) (*GetAssetsHistoryReply, error)
	// 订单历史（充值-提现-内部转账）
	GetOrderHistory(ctx context.Context, in *GetOrderHistoryRequest, opts ...grpc.CallOption) (*GetOrderHistoryReply, error)
	// 执行提现api
	WithDraw(ctx context.Context, in *WithDrawRequest, opts ...grpc.CallOption) (*WithDrawReply, error)
}

type orderServerClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServerClient(cc grpc.ClientConnInterface) OrderServerClient {
	return &orderServerClient{cc}
}

func (c *orderServerClient) Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferReply, error) {
	out := new(TransferReply)
	err := c.cc.Invoke(ctx, OrderServer_Transfer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServerClient) GetAssetsByAccountID(ctx context.Context, in *GetAssetsByAccountIDRequest, opts ...grpc.CallOption) (*GetAssetsByAccountIDReply, error) {
	out := new(GetAssetsByAccountIDReply)
	err := c.cc.Invoke(ctx, OrderServer_GetAssetsByAccountID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServerClient) GetAssetsHistory(ctx context.Context, in *GetAssetsHistoryRequest, opts ...grpc.CallOption) (*GetAssetsHistoryReply, error) {
	out := new(GetAssetsHistoryReply)
	err := c.cc.Invoke(ctx, OrderServer_GetAssetsHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServerClient) GetOrderHistory(ctx context.Context, in *GetOrderHistoryRequest, opts ...grpc.CallOption) (*GetOrderHistoryReply, error) {
	out := new(GetOrderHistoryReply)
	err := c.cc.Invoke(ctx, OrderServer_GetOrderHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServerClient) WithDraw(ctx context.Context, in *WithDrawRequest, opts ...grpc.CallOption) (*WithDrawReply, error) {
	out := new(WithDrawReply)
	err := c.cc.Invoke(ctx, OrderServer_WithDraw_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServerServer is the server API for OrderServer service.
// All implementations must embed UnimplementedOrderServerServer
// for forward compatibility
type OrderServerServer interface {
	// 执行转账
	Transfer(context.Context, *TransferRequest) (*TransferReply, error)
	// 1. 账户资产查询
	GetAssetsByAccountID(context.Context, *GetAssetsByAccountIDRequest) (*GetAssetsByAccountIDReply, error)
	// 资产变动记录
	GetAssetsHistory(context.Context, *GetAssetsHistoryRequest) (*GetAssetsHistoryReply, error)
	// 订单历史（充值-提现-内部转账）
	GetOrderHistory(context.Context, *GetOrderHistoryRequest) (*GetOrderHistoryReply, error)
	// 执行提现api
	WithDraw(context.Context, *WithDrawRequest) (*WithDrawReply, error)
	mustEmbedUnimplementedOrderServerServer()
}

// UnimplementedOrderServerServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServerServer struct {
}

func (UnimplementedOrderServerServer) Transfer(context.Context, *TransferRequest) (*TransferReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transfer not implemented")
}
func (UnimplementedOrderServerServer) GetAssetsByAccountID(context.Context, *GetAssetsByAccountIDRequest) (*GetAssetsByAccountIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssetsByAccountID not implemented")
}
func (UnimplementedOrderServerServer) GetAssetsHistory(context.Context, *GetAssetsHistoryRequest) (*GetAssetsHistoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssetsHistory not implemented")
}
func (UnimplementedOrderServerServer) GetOrderHistory(context.Context, *GetOrderHistoryRequest) (*GetOrderHistoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderHistory not implemented")
}
func (UnimplementedOrderServerServer) WithDraw(context.Context, *WithDrawRequest) (*WithDrawReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithDraw not implemented")
}
func (UnimplementedOrderServerServer) mustEmbedUnimplementedOrderServerServer() {}

// UnsafeOrderServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServerServer will
// result in compilation errors.
type UnsafeOrderServerServer interface {
	mustEmbedUnimplementedOrderServerServer()
}

func RegisterOrderServerServer(s grpc.ServiceRegistrar, srv OrderServerServer) {
	s.RegisterService(&OrderServer_ServiceDesc, srv)
}

func _OrderServer_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServerServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderServer_Transfer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServerServer).Transfer(ctx, req.(*TransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderServer_GetAssetsByAccountID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAssetsByAccountIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServerServer).GetAssetsByAccountID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderServer_GetAssetsByAccountID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServerServer).GetAssetsByAccountID(ctx, req.(*GetAssetsByAccountIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderServer_GetAssetsHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAssetsHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServerServer).GetAssetsHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderServer_GetAssetsHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServerServer).GetAssetsHistory(ctx, req.(*GetAssetsHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderServer_GetOrderHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServerServer).GetOrderHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderServer_GetOrderHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServerServer).GetOrderHistory(ctx, req.(*GetOrderHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderServer_WithDraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithDrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServerServer).WithDraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderServer_WithDraw_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServerServer).WithDraw(ctx, req.(*WithDrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderServer_ServiceDesc is the grpc.ServiceDesc for OrderServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.OrderServer",
	HandlerType: (*OrderServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Transfer",
			Handler:    _OrderServer_Transfer_Handler,
		},
		{
			MethodName: "GetAssetsByAccountID",
			Handler:    _OrderServer_GetAssetsByAccountID_Handler,
		},
		{
			MethodName: "GetAssetsHistory",
			Handler:    _OrderServer_GetAssetsHistory_Handler,
		},
		{
			MethodName: "GetOrderHistory",
			Handler:    _OrderServer_GetOrderHistory_Handler,
		},
		{
			MethodName: "WithDraw",
			Handler:    _OrderServer_WithDraw_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}