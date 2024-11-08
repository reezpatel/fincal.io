// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: proto/account.proto

package fincal

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AccountService_DepositAccounts_FullMethodName         = "/fincal.AccountService/DepositAccounts"
	AccountService_CreateDepositAccount_FullMethodName    = "/fincal.AccountService/createDepositAccount"
	AccountService_DeleteDepositAccount_FullMethodName    = "/fincal.AccountService/deleteDepositAccount"
	AccountService_CreditCardAccounts_FullMethodName      = "/fincal.AccountService/CreditCardAccounts"
	AccountService_CreateCreditCardAccount_FullMethodName = "/fincal.AccountService/createCreditCardAccount"
	AccountService_DeleteCreditCardAccount_FullMethodName = "/fincal.AccountService/deleteCreditCardAccount"
	AccountService_SecurityAccounts_FullMethodName        = "/fincal.AccountService/SecurityAccounts"
	AccountService_CreateSecurityAccount_FullMethodName   = "/fincal.AccountService/createSecurityAccount"
	AccountService_DeleteSecurityAccount_FullMethodName   = "/fincal.AccountService/deleteSecurityAccount"
	AccountService_SecurityItems_FullMethodName           = "/fincal.AccountService/SecurityItems"
	AccountService_CreateSecurityItem_FullMethodName      = "/fincal.AccountService/createSecurityItem"
	AccountService_DeleteSecurityItem_FullMethodName      = "/fincal.AccountService/deleteSecurityItem"
)

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountServiceClient interface {
	DepositAccounts(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DepositAccountsResponse, error)
	CreateDepositAccount(ctx context.Context, in *CreateDepositAccountRequest, opts ...grpc.CallOption) (*CreateDepositAccountResponse, error)
	DeleteDepositAccount(ctx context.Context, in *DeleteDepositAccountRequest, opts ...grpc.CallOption) (*DeleteDepositAccountResponse, error)
	CreditCardAccounts(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CreditCardAccountsResponse, error)
	CreateCreditCardAccount(ctx context.Context, in *CreateCreditCardAccountRequest, opts ...grpc.CallOption) (*CreateCreditCardAccountResponse, error)
	DeleteCreditCardAccount(ctx context.Context, in *DeleteCreditCardAccountRequest, opts ...grpc.CallOption) (*DeleteCreditCardAccountResponse, error)
	SecurityAccounts(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SecurityAccountsResponse, error)
	CreateSecurityAccount(ctx context.Context, in *CreateSecurityAccountRequest, opts ...grpc.CallOption) (*CreateSecurityAccountResponse, error)
	DeleteSecurityAccount(ctx context.Context, in *DeleteSecurityAccountRequest, opts ...grpc.CallOption) (*DeleteSecurityAccountResponse, error)
	SecurityItems(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SecurityItemsResponse, error)
	CreateSecurityItem(ctx context.Context, in *CreateSecurityItemRequest, opts ...grpc.CallOption) (*CreateSecurityItemResponse, error)
	DeleteSecurityItem(ctx context.Context, in *DeleteSecurityItemRequest, opts ...grpc.CallOption) (*DeleteSecurityItemResponse, error)
}

type accountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountServiceClient(cc grpc.ClientConnInterface) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) DepositAccounts(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DepositAccountsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DepositAccountsResponse)
	err := c.cc.Invoke(ctx, AccountService_DepositAccounts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) CreateDepositAccount(ctx context.Context, in *CreateDepositAccountRequest, opts ...grpc.CallOption) (*CreateDepositAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDepositAccountResponse)
	err := c.cc.Invoke(ctx, AccountService_CreateDepositAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) DeleteDepositAccount(ctx context.Context, in *DeleteDepositAccountRequest, opts ...grpc.CallOption) (*DeleteDepositAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteDepositAccountResponse)
	err := c.cc.Invoke(ctx, AccountService_DeleteDepositAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) CreditCardAccounts(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CreditCardAccountsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreditCardAccountsResponse)
	err := c.cc.Invoke(ctx, AccountService_CreditCardAccounts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) CreateCreditCardAccount(ctx context.Context, in *CreateCreditCardAccountRequest, opts ...grpc.CallOption) (*CreateCreditCardAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCreditCardAccountResponse)
	err := c.cc.Invoke(ctx, AccountService_CreateCreditCardAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) DeleteCreditCardAccount(ctx context.Context, in *DeleteCreditCardAccountRequest, opts ...grpc.CallOption) (*DeleteCreditCardAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCreditCardAccountResponse)
	err := c.cc.Invoke(ctx, AccountService_DeleteCreditCardAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) SecurityAccounts(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SecurityAccountsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SecurityAccountsResponse)
	err := c.cc.Invoke(ctx, AccountService_SecurityAccounts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) CreateSecurityAccount(ctx context.Context, in *CreateSecurityAccountRequest, opts ...grpc.CallOption) (*CreateSecurityAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateSecurityAccountResponse)
	err := c.cc.Invoke(ctx, AccountService_CreateSecurityAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) DeleteSecurityAccount(ctx context.Context, in *DeleteSecurityAccountRequest, opts ...grpc.CallOption) (*DeleteSecurityAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteSecurityAccountResponse)
	err := c.cc.Invoke(ctx, AccountService_DeleteSecurityAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) SecurityItems(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SecurityItemsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SecurityItemsResponse)
	err := c.cc.Invoke(ctx, AccountService_SecurityItems_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) CreateSecurityItem(ctx context.Context, in *CreateSecurityItemRequest, opts ...grpc.CallOption) (*CreateSecurityItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateSecurityItemResponse)
	err := c.cc.Invoke(ctx, AccountService_CreateSecurityItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) DeleteSecurityItem(ctx context.Context, in *DeleteSecurityItemRequest, opts ...grpc.CallOption) (*DeleteSecurityItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteSecurityItemResponse)
	err := c.cc.Invoke(ctx, AccountService_DeleteSecurityItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
// All implementations must embed UnimplementedAccountServiceServer
// for forward compatibility.
type AccountServiceServer interface {
	DepositAccounts(context.Context, *emptypb.Empty) (*DepositAccountsResponse, error)
	CreateDepositAccount(context.Context, *CreateDepositAccountRequest) (*CreateDepositAccountResponse, error)
	DeleteDepositAccount(context.Context, *DeleteDepositAccountRequest) (*DeleteDepositAccountResponse, error)
	CreditCardAccounts(context.Context, *emptypb.Empty) (*CreditCardAccountsResponse, error)
	CreateCreditCardAccount(context.Context, *CreateCreditCardAccountRequest) (*CreateCreditCardAccountResponse, error)
	DeleteCreditCardAccount(context.Context, *DeleteCreditCardAccountRequest) (*DeleteCreditCardAccountResponse, error)
	SecurityAccounts(context.Context, *emptypb.Empty) (*SecurityAccountsResponse, error)
	CreateSecurityAccount(context.Context, *CreateSecurityAccountRequest) (*CreateSecurityAccountResponse, error)
	DeleteSecurityAccount(context.Context, *DeleteSecurityAccountRequest) (*DeleteSecurityAccountResponse, error)
	SecurityItems(context.Context, *emptypb.Empty) (*SecurityItemsResponse, error)
	CreateSecurityItem(context.Context, *CreateSecurityItemRequest) (*CreateSecurityItemResponse, error)
	DeleteSecurityItem(context.Context, *DeleteSecurityItemRequest) (*DeleteSecurityItemResponse, error)
	mustEmbedUnimplementedAccountServiceServer()
}

// UnimplementedAccountServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAccountServiceServer struct{}

func (UnimplementedAccountServiceServer) DepositAccounts(context.Context, *emptypb.Empty) (*DepositAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DepositAccounts not implemented")
}
func (UnimplementedAccountServiceServer) CreateDepositAccount(context.Context, *CreateDepositAccountRequest) (*CreateDepositAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDepositAccount not implemented")
}
func (UnimplementedAccountServiceServer) DeleteDepositAccount(context.Context, *DeleteDepositAccountRequest) (*DeleteDepositAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDepositAccount not implemented")
}
func (UnimplementedAccountServiceServer) CreditCardAccounts(context.Context, *emptypb.Empty) (*CreditCardAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreditCardAccounts not implemented")
}
func (UnimplementedAccountServiceServer) CreateCreditCardAccount(context.Context, *CreateCreditCardAccountRequest) (*CreateCreditCardAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCreditCardAccount not implemented")
}
func (UnimplementedAccountServiceServer) DeleteCreditCardAccount(context.Context, *DeleteCreditCardAccountRequest) (*DeleteCreditCardAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCreditCardAccount not implemented")
}
func (UnimplementedAccountServiceServer) SecurityAccounts(context.Context, *emptypb.Empty) (*SecurityAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SecurityAccounts not implemented")
}
func (UnimplementedAccountServiceServer) CreateSecurityAccount(context.Context, *CreateSecurityAccountRequest) (*CreateSecurityAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSecurityAccount not implemented")
}
func (UnimplementedAccountServiceServer) DeleteSecurityAccount(context.Context, *DeleteSecurityAccountRequest) (*DeleteSecurityAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSecurityAccount not implemented")
}
func (UnimplementedAccountServiceServer) SecurityItems(context.Context, *emptypb.Empty) (*SecurityItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SecurityItems not implemented")
}
func (UnimplementedAccountServiceServer) CreateSecurityItem(context.Context, *CreateSecurityItemRequest) (*CreateSecurityItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSecurityItem not implemented")
}
func (UnimplementedAccountServiceServer) DeleteSecurityItem(context.Context, *DeleteSecurityItemRequest) (*DeleteSecurityItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSecurityItem not implemented")
}
func (UnimplementedAccountServiceServer) mustEmbedUnimplementedAccountServiceServer() {}
func (UnimplementedAccountServiceServer) testEmbeddedByValue()                        {}

// UnsafeAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServiceServer will
// result in compilation errors.
type UnsafeAccountServiceServer interface {
	mustEmbedUnimplementedAccountServiceServer()
}

func RegisterAccountServiceServer(s grpc.ServiceRegistrar, srv AccountServiceServer) {
	// If the following call pancis, it indicates UnimplementedAccountServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AccountService_ServiceDesc, srv)
}

func _AccountService_DepositAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).DepositAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_DepositAccounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).DepositAccounts(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_CreateDepositAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDepositAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CreateDepositAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_CreateDepositAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CreateDepositAccount(ctx, req.(*CreateDepositAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_DeleteDepositAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDepositAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).DeleteDepositAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_DeleteDepositAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).DeleteDepositAccount(ctx, req.(*DeleteDepositAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_CreditCardAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CreditCardAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_CreditCardAccounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CreditCardAccounts(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_CreateCreditCardAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCreditCardAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CreateCreditCardAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_CreateCreditCardAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CreateCreditCardAccount(ctx, req.(*CreateCreditCardAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_DeleteCreditCardAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCreditCardAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).DeleteCreditCardAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_DeleteCreditCardAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).DeleteCreditCardAccount(ctx, req.(*DeleteCreditCardAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_SecurityAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).SecurityAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_SecurityAccounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).SecurityAccounts(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_CreateSecurityAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSecurityAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CreateSecurityAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_CreateSecurityAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CreateSecurityAccount(ctx, req.(*CreateSecurityAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_DeleteSecurityAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSecurityAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).DeleteSecurityAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_DeleteSecurityAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).DeleteSecurityAccount(ctx, req.(*DeleteSecurityAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_SecurityItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).SecurityItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_SecurityItems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).SecurityItems(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_CreateSecurityItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSecurityItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CreateSecurityItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_CreateSecurityItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CreateSecurityItem(ctx, req.(*CreateSecurityItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_DeleteSecurityItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSecurityItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).DeleteSecurityItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountService_DeleteSecurityItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).DeleteSecurityItem(ctx, req.(*DeleteSecurityItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountService_ServiceDesc is the grpc.ServiceDesc for AccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fincal.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DepositAccounts",
			Handler:    _AccountService_DepositAccounts_Handler,
		},
		{
			MethodName: "createDepositAccount",
			Handler:    _AccountService_CreateDepositAccount_Handler,
		},
		{
			MethodName: "deleteDepositAccount",
			Handler:    _AccountService_DeleteDepositAccount_Handler,
		},
		{
			MethodName: "CreditCardAccounts",
			Handler:    _AccountService_CreditCardAccounts_Handler,
		},
		{
			MethodName: "createCreditCardAccount",
			Handler:    _AccountService_CreateCreditCardAccount_Handler,
		},
		{
			MethodName: "deleteCreditCardAccount",
			Handler:    _AccountService_DeleteCreditCardAccount_Handler,
		},
		{
			MethodName: "SecurityAccounts",
			Handler:    _AccountService_SecurityAccounts_Handler,
		},
		{
			MethodName: "createSecurityAccount",
			Handler:    _AccountService_CreateSecurityAccount_Handler,
		},
		{
			MethodName: "deleteSecurityAccount",
			Handler:    _AccountService_DeleteSecurityAccount_Handler,
		},
		{
			MethodName: "SecurityItems",
			Handler:    _AccountService_SecurityItems_Handler,
		},
		{
			MethodName: "createSecurityItem",
			Handler:    _AccountService_CreateSecurityItem_Handler,
		},
		{
			MethodName: "deleteSecurityItem",
			Handler:    _AccountService_DeleteSecurityItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/account.proto",
}
