// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: proto/transaction.proto

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
	TransactionService_Transactions_FullMethodName               = "/fincal.TransactionService/Transactions"
	TransactionService_CreateTransaction_FullMethodName          = "/fincal.TransactionService/createTransaction"
	TransactionService_DeleteTransaction_FullMethodName          = "/fincal.TransactionService/deleteTransaction"
	TransactionService_RecurringTransactions_FullMethodName      = "/fincal.TransactionService/RecurringTransactions"
	TransactionService_CreateRecurringTransaction_FullMethodName = "/fincal.TransactionService/createRecurringTransaction"
	TransactionService_DeleteRecurringTransaction_FullMethodName = "/fincal.TransactionService/deleteRecurringTransaction"
)

// TransactionServiceClient is the client API for TransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionServiceClient interface {
	Transactions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TransactionsResponse, error)
	CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error)
	DeleteTransaction(ctx context.Context, in *DeleteTransactionRequest, opts ...grpc.CallOption) (*DeleteTransactionResponse, error)
	RecurringTransactions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*RecurringTransactionsResponse, error)
	CreateRecurringTransaction(ctx context.Context, in *CreateRecurringTransactionRequest, opts ...grpc.CallOption) (*CreateRecurringTransactionResponse, error)
	DeleteRecurringTransaction(ctx context.Context, in *DeleteRecurringTransactionRequest, opts ...grpc.CallOption) (*DeleteRecurringTransactionResponse, error)
}

type transactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionServiceClient(cc grpc.ClientConnInterface) TransactionServiceClient {
	return &transactionServiceClient{cc}
}

func (c *transactionServiceClient) Transactions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TransactionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransactionsResponse)
	err := c.cc.Invoke(ctx, TransactionService_Transactions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTransactionResponse)
	err := c.cc.Invoke(ctx, TransactionService_CreateTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) DeleteTransaction(ctx context.Context, in *DeleteTransactionRequest, opts ...grpc.CallOption) (*DeleteTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTransactionResponse)
	err := c.cc.Invoke(ctx, TransactionService_DeleteTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) RecurringTransactions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*RecurringTransactionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RecurringTransactionsResponse)
	err := c.cc.Invoke(ctx, TransactionService_RecurringTransactions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) CreateRecurringTransaction(ctx context.Context, in *CreateRecurringTransactionRequest, opts ...grpc.CallOption) (*CreateRecurringTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateRecurringTransactionResponse)
	err := c.cc.Invoke(ctx, TransactionService_CreateRecurringTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) DeleteRecurringTransaction(ctx context.Context, in *DeleteRecurringTransactionRequest, opts ...grpc.CallOption) (*DeleteRecurringTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteRecurringTransactionResponse)
	err := c.cc.Invoke(ctx, TransactionService_DeleteRecurringTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServiceServer is the server API for TransactionService service.
// All implementations must embed UnimplementedTransactionServiceServer
// for forward compatibility.
type TransactionServiceServer interface {
	Transactions(context.Context, *emptypb.Empty) (*TransactionsResponse, error)
	CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error)
	DeleteTransaction(context.Context, *DeleteTransactionRequest) (*DeleteTransactionResponse, error)
	RecurringTransactions(context.Context, *emptypb.Empty) (*RecurringTransactionsResponse, error)
	CreateRecurringTransaction(context.Context, *CreateRecurringTransactionRequest) (*CreateRecurringTransactionResponse, error)
	DeleteRecurringTransaction(context.Context, *DeleteRecurringTransactionRequest) (*DeleteRecurringTransactionResponse, error)
	mustEmbedUnimplementedTransactionServiceServer()
}

// UnimplementedTransactionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTransactionServiceServer struct{}

func (UnimplementedTransactionServiceServer) Transactions(context.Context, *emptypb.Empty) (*TransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transactions not implemented")
}
func (UnimplementedTransactionServiceServer) CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}
func (UnimplementedTransactionServiceServer) DeleteTransaction(context.Context, *DeleteTransactionRequest) (*DeleteTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTransaction not implemented")
}
func (UnimplementedTransactionServiceServer) RecurringTransactions(context.Context, *emptypb.Empty) (*RecurringTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecurringTransactions not implemented")
}
func (UnimplementedTransactionServiceServer) CreateRecurringTransaction(context.Context, *CreateRecurringTransactionRequest) (*CreateRecurringTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRecurringTransaction not implemented")
}
func (UnimplementedTransactionServiceServer) DeleteRecurringTransaction(context.Context, *DeleteRecurringTransactionRequest) (*DeleteRecurringTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRecurringTransaction not implemented")
}
func (UnimplementedTransactionServiceServer) mustEmbedUnimplementedTransactionServiceServer() {}
func (UnimplementedTransactionServiceServer) testEmbeddedByValue()                            {}

// UnsafeTransactionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionServiceServer will
// result in compilation errors.
type UnsafeTransactionServiceServer interface {
	mustEmbedUnimplementedTransactionServiceServer()
}

func RegisterTransactionServiceServer(s grpc.ServiceRegistrar, srv TransactionServiceServer) {
	// If the following call pancis, it indicates UnimplementedTransactionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TransactionService_ServiceDesc, srv)
}

func _TransactionService_Transactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).Transactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_Transactions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).Transactions(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_CreateTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).CreateTransaction(ctx, req.(*CreateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_DeleteTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).DeleteTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_DeleteTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).DeleteTransaction(ctx, req.(*DeleteTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_RecurringTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).RecurringTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_RecurringTransactions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).RecurringTransactions(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_CreateRecurringTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRecurringTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).CreateRecurringTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_CreateRecurringTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).CreateRecurringTransaction(ctx, req.(*CreateRecurringTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_DeleteRecurringTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRecurringTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).DeleteRecurringTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_DeleteRecurringTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).DeleteRecurringTransaction(ctx, req.(*DeleteRecurringTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransactionService_ServiceDesc is the grpc.ServiceDesc for TransactionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fincal.TransactionService",
	HandlerType: (*TransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Transactions",
			Handler:    _TransactionService_Transactions_Handler,
		},
		{
			MethodName: "createTransaction",
			Handler:    _TransactionService_CreateTransaction_Handler,
		},
		{
			MethodName: "deleteTransaction",
			Handler:    _TransactionService_DeleteTransaction_Handler,
		},
		{
			MethodName: "RecurringTransactions",
			Handler:    _TransactionService_RecurringTransactions_Handler,
		},
		{
			MethodName: "createRecurringTransaction",
			Handler:    _TransactionService_CreateRecurringTransaction_Handler,
		},
		{
			MethodName: "deleteRecurringTransaction",
			Handler:    _TransactionService_DeleteRecurringTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/transaction.proto",
}
