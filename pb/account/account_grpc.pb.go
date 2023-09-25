// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: account.proto

package account

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

// AccountServerClient is the client API for AccountServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountServerClient interface {
	// 从 apiKey 获取到对应的商户信息
	GetMerchantInfo(ctx context.Context, in *GetMerchantInfoRequest, opts ...grpc.CallOption) (*GetMerchantInfoReply, error)
	GetAccountInfo(ctx context.Context, in *GetAccountInfoRequest, opts ...grpc.CallOption) (*GetAccountInfoReply, error)
	GetAssetStatus(ctx context.Context, in *GetAssetStatusRequest, opts ...grpc.CallOption) (*GetAssetStatusReply, error)
}

type accountServerClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountServerClient(cc grpc.ClientConnInterface) AccountServerClient {
	return &accountServerClient{cc}
}

func (c *accountServerClient) GetMerchantInfo(ctx context.Context, in *GetMerchantInfoRequest, opts ...grpc.CallOption) (*GetMerchantInfoReply, error) {
	out := new(GetMerchantInfoReply)
	err := c.cc.Invoke(ctx, "/account.AccountServer/GetMerchantInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServerClient) GetAccountInfo(ctx context.Context, in *GetAccountInfoRequest, opts ...grpc.CallOption) (*GetAccountInfoReply, error) {
	out := new(GetAccountInfoReply)
	err := c.cc.Invoke(ctx, "/account.AccountServer/GetAccountInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServerClient) GetAssetStatus(ctx context.Context, in *GetAssetStatusRequest, opts ...grpc.CallOption) (*GetAssetStatusReply, error) {
	out := new(GetAssetStatusReply)
	err := c.cc.Invoke(ctx, "/account.AccountServer/GetAssetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServerServer is the server API for AccountServer service.
// All implementations must embed UnimplementedAccountServerServer
// for forward compatibility
type AccountServerServer interface {
	// 从 apiKey 获取到对应的商户信息
	GetMerchantInfo(context.Context, *GetMerchantInfoRequest) (*GetMerchantInfoReply, error)
	GetAccountInfo(context.Context, *GetAccountInfoRequest) (*GetAccountInfoReply, error)
	GetAssetStatus(context.Context, *GetAssetStatusRequest) (*GetAssetStatusReply, error)
	mustEmbedUnimplementedAccountServerServer()
}

// UnimplementedAccountServerServer must be embedded to have forward compatible implementations.
type UnimplementedAccountServerServer struct {
}

func (UnimplementedAccountServerServer) GetMerchantInfo(context.Context, *GetMerchantInfoRequest) (*GetMerchantInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMerchantInfo not implemented")
}
func (UnimplementedAccountServerServer) GetAccountInfo(context.Context, *GetAccountInfoRequest) (*GetAccountInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountInfo not implemented")
}
func (UnimplementedAccountServerServer) GetAssetStatus(context.Context, *GetAssetStatusRequest) (*GetAssetStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssetStatus not implemented")
}
func (UnimplementedAccountServerServer) mustEmbedUnimplementedAccountServerServer() {}

// UnsafeAccountServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServerServer will
// result in compilation errors.
type UnsafeAccountServerServer interface {
	mustEmbedUnimplementedAccountServerServer()
}

func RegisterAccountServerServer(s grpc.ServiceRegistrar, srv AccountServerServer) {
	s.RegisterService(&AccountServer_ServiceDesc, srv)
}

func _AccountServer_GetMerchantInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMerchantInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServerServer).GetMerchantInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountServer/GetMerchantInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServerServer).GetMerchantInfo(ctx, req.(*GetMerchantInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountServer_GetAccountInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServerServer).GetAccountInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountServer/GetAccountInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServerServer).GetAccountInfo(ctx, req.(*GetAccountInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountServer_GetAssetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAssetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServerServer).GetAssetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountServer/GetAssetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServerServer).GetAssetStatus(ctx, req.(*GetAssetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountServer_ServiceDesc is the grpc.ServiceDesc for AccountServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "account.AccountServer",
	HandlerType: (*AccountServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMerchantInfo",
			Handler:    _AccountServer_GetMerchantInfo_Handler,
		},
		{
			MethodName: "GetAccountInfo",
			Handler:    _AccountServer_GetAccountInfo_Handler,
		},
		{
			MethodName: "GetAssetStatus",
			Handler:    _AccountServer_GetAssetStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account.proto",
}
