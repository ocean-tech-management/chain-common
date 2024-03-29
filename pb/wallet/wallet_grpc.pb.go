// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: wallet.proto

// 定义包名

package wallet

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
	WalletServer_WithDraw_FullMethodName            = "/wallet.WalletServer/WithDraw"
	WalletServer_FoundFee_FullMethodName            = "/wallet.WalletServer/FoundFee"
	WalletServer_Collect_FullMethodName             = "/wallet.WalletServer/Collect"
	WalletServer_AddChain_FullMethodName            = "/wallet.WalletServer/AddChain"
	WalletServer_GetChainByChainName_FullMethodName = "/wallet.WalletServer/GetChainByChainName"
	WalletServer_GetAllChain_FullMethodName         = "/wallet.WalletServer/GetAllChain"
	WalletServer_AddToken_FullMethodName            = "/wallet.WalletServer/AddToken"
	WalletServer_GetTokenBySymbol_FullMethodName    = "/wallet.WalletServer/GetTokenBySymbol"
	WalletServer_GetTokens_FullMethodName           = "/wallet.WalletServer/GetTokens"
)

// WalletServerClient is the client API for WalletServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletServerClient interface {
	WithDraw(ctx context.Context, in *WithDrawRequest, opts ...grpc.CallOption) (*WithDrawReply, error)
	FoundFee(ctx context.Context, in *FoundFeeRequest, opts ...grpc.CallOption) (*WithDrawReply, error)
	Collect(ctx context.Context, in *CollectRequest, opts ...grpc.CallOption) (*WithDrawReply, error)
	AddChain(ctx context.Context, in *AddChainRequest, opts ...grpc.CallOption) (*AddChainReply, error)
	GetChainByChainName(ctx context.Context, in *GetChainByChainNameRequest, opts ...grpc.CallOption) (*GetChainByChainNameReply, error)
	// 支持所有公链查询
	GetAllChain(ctx context.Context, in *GetAllChainRequest, opts ...grpc.CallOption) (*GetAllChainReply, error)
	AddToken(ctx context.Context, in *AddTokenRequest, opts ...grpc.CallOption) (*AddTokenReReply, error)
	GetTokenBySymbol(ctx context.Context, in *GetTokenBySymbolRequest, opts ...grpc.CallOption) (*GetTokenBySymbolReply, error)
	GetTokens(ctx context.Context, in *GetTokensRequest, opts ...grpc.CallOption) (*GetTokensReply, error)
}

type walletServerClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletServerClient(cc grpc.ClientConnInterface) WalletServerClient {
	return &walletServerClient{cc}
}

func (c *walletServerClient) WithDraw(ctx context.Context, in *WithDrawRequest, opts ...grpc.CallOption) (*WithDrawReply, error) {
	out := new(WithDrawReply)
	err := c.cc.Invoke(ctx, WalletServer_WithDraw_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServerClient) FoundFee(ctx context.Context, in *FoundFeeRequest, opts ...grpc.CallOption) (*WithDrawReply, error) {
	out := new(WithDrawReply)
	err := c.cc.Invoke(ctx, WalletServer_FoundFee_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServerClient) Collect(ctx context.Context, in *CollectRequest, opts ...grpc.CallOption) (*WithDrawReply, error) {
	out := new(WithDrawReply)
	err := c.cc.Invoke(ctx, WalletServer_Collect_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServerClient) AddChain(ctx context.Context, in *AddChainRequest, opts ...grpc.CallOption) (*AddChainReply, error) {
	out := new(AddChainReply)
	err := c.cc.Invoke(ctx, WalletServer_AddChain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServerClient) GetChainByChainName(ctx context.Context, in *GetChainByChainNameRequest, opts ...grpc.CallOption) (*GetChainByChainNameReply, error) {
	out := new(GetChainByChainNameReply)
	err := c.cc.Invoke(ctx, WalletServer_GetChainByChainName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServerClient) GetAllChain(ctx context.Context, in *GetAllChainRequest, opts ...grpc.CallOption) (*GetAllChainReply, error) {
	out := new(GetAllChainReply)
	err := c.cc.Invoke(ctx, WalletServer_GetAllChain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServerClient) AddToken(ctx context.Context, in *AddTokenRequest, opts ...grpc.CallOption) (*AddTokenReReply, error) {
	out := new(AddTokenReReply)
	err := c.cc.Invoke(ctx, WalletServer_AddToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServerClient) GetTokenBySymbol(ctx context.Context, in *GetTokenBySymbolRequest, opts ...grpc.CallOption) (*GetTokenBySymbolReply, error) {
	out := new(GetTokenBySymbolReply)
	err := c.cc.Invoke(ctx, WalletServer_GetTokenBySymbol_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServerClient) GetTokens(ctx context.Context, in *GetTokensRequest, opts ...grpc.CallOption) (*GetTokensReply, error) {
	out := new(GetTokensReply)
	err := c.cc.Invoke(ctx, WalletServer_GetTokens_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletServerServer is the server API for WalletServer service.
// All implementations must embed UnimplementedWalletServerServer
// for forward compatibility
type WalletServerServer interface {
	WithDraw(context.Context, *WithDrawRequest) (*WithDrawReply, error)
	FoundFee(context.Context, *FoundFeeRequest) (*WithDrawReply, error)
	Collect(context.Context, *CollectRequest) (*WithDrawReply, error)
	AddChain(context.Context, *AddChainRequest) (*AddChainReply, error)
	GetChainByChainName(context.Context, *GetChainByChainNameRequest) (*GetChainByChainNameReply, error)
	// 支持所有公链查询
	GetAllChain(context.Context, *GetAllChainRequest) (*GetAllChainReply, error)
	AddToken(context.Context, *AddTokenRequest) (*AddTokenReReply, error)
	GetTokenBySymbol(context.Context, *GetTokenBySymbolRequest) (*GetTokenBySymbolReply, error)
	GetTokens(context.Context, *GetTokensRequest) (*GetTokensReply, error)
	mustEmbedUnimplementedWalletServerServer()
}

// UnimplementedWalletServerServer must be embedded to have forward compatible implementations.
type UnimplementedWalletServerServer struct {
}

func (UnimplementedWalletServerServer) WithDraw(context.Context, *WithDrawRequest) (*WithDrawReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithDraw not implemented")
}
func (UnimplementedWalletServerServer) FoundFee(context.Context, *FoundFeeRequest) (*WithDrawReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FoundFee not implemented")
}
func (UnimplementedWalletServerServer) Collect(context.Context, *CollectRequest) (*WithDrawReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Collect not implemented")
}
func (UnimplementedWalletServerServer) AddChain(context.Context, *AddChainRequest) (*AddChainReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddChain not implemented")
}
func (UnimplementedWalletServerServer) GetChainByChainName(context.Context, *GetChainByChainNameRequest) (*GetChainByChainNameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChainByChainName not implemented")
}
func (UnimplementedWalletServerServer) GetAllChain(context.Context, *GetAllChainRequest) (*GetAllChainReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllChain not implemented")
}
func (UnimplementedWalletServerServer) AddToken(context.Context, *AddTokenRequest) (*AddTokenReReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToken not implemented")
}
func (UnimplementedWalletServerServer) GetTokenBySymbol(context.Context, *GetTokenBySymbolRequest) (*GetTokenBySymbolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenBySymbol not implemented")
}
func (UnimplementedWalletServerServer) GetTokens(context.Context, *GetTokensRequest) (*GetTokensReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokens not implemented")
}
func (UnimplementedWalletServerServer) mustEmbedUnimplementedWalletServerServer() {}

// UnsafeWalletServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletServerServer will
// result in compilation errors.
type UnsafeWalletServerServer interface {
	mustEmbedUnimplementedWalletServerServer()
}

func RegisterWalletServerServer(s grpc.ServiceRegistrar, srv WalletServerServer) {
	s.RegisterService(&WalletServer_ServiceDesc, srv)
}

func _WalletServer_WithDraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithDrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServerServer).WithDraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletServer_WithDraw_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServerServer).WithDraw(ctx, req.(*WithDrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletServer_FoundFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FoundFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServerServer).FoundFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletServer_FoundFee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServerServer).FoundFee(ctx, req.(*FoundFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletServer_Collect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServerServer).Collect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletServer_Collect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServerServer).Collect(ctx, req.(*CollectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletServer_AddChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddChainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServerServer).AddChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletServer_AddChain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServerServer).AddChain(ctx, req.(*AddChainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletServer_GetChainByChainName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChainByChainNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServerServer).GetChainByChainName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletServer_GetChainByChainName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServerServer).GetChainByChainName(ctx, req.(*GetChainByChainNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletServer_GetAllChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllChainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServerServer).GetAllChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletServer_GetAllChain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServerServer).GetAllChain(ctx, req.(*GetAllChainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletServer_AddToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServerServer).AddToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletServer_AddToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServerServer).AddToken(ctx, req.(*AddTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletServer_GetTokenBySymbol_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenBySymbolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServerServer).GetTokenBySymbol(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletServer_GetTokenBySymbol_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServerServer).GetTokenBySymbol(ctx, req.(*GetTokenBySymbolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletServer_GetTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServerServer).GetTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletServer_GetTokens_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServerServer).GetTokens(ctx, req.(*GetTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WalletServer_ServiceDesc is the grpc.ServiceDesc for WalletServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WalletServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wallet.WalletServer",
	HandlerType: (*WalletServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WithDraw",
			Handler:    _WalletServer_WithDraw_Handler,
		},
		{
			MethodName: "FoundFee",
			Handler:    _WalletServer_FoundFee_Handler,
		},
		{
			MethodName: "Collect",
			Handler:    _WalletServer_Collect_Handler,
		},
		{
			MethodName: "AddChain",
			Handler:    _WalletServer_AddChain_Handler,
		},
		{
			MethodName: "GetChainByChainName",
			Handler:    _WalletServer_GetChainByChainName_Handler,
		},
		{
			MethodName: "GetAllChain",
			Handler:    _WalletServer_GetAllChain_Handler,
		},
		{
			MethodName: "AddToken",
			Handler:    _WalletServer_AddToken_Handler,
		},
		{
			MethodName: "GetTokenBySymbol",
			Handler:    _WalletServer_GetTokenBySymbol_Handler,
		},
		{
			MethodName: "GetTokens",
			Handler:    _WalletServer_GetTokens_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallet.proto",
}
