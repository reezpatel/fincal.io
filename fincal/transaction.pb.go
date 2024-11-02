// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: proto/transaction.proto

package fincal

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TransactionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transactions []*Transaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *TransactionsResponse) Reset() {
	*x = TransactionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionsResponse) ProtoMessage() {}

func (x *TransactionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionsResponse.ProtoReflect.Descriptor instead.
func (*TransactionsResponse) Descriptor() ([]byte, []int) {
	return file_proto_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionsResponse) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type CreateTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *CreateTransactionRequest) Reset() {
	*x = CreateTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTransactionRequest) ProtoMessage() {}

func (x *CreateTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTransactionRequest.ProtoReflect.Descriptor instead.
func (*CreateTransactionRequest) Descriptor() ([]byte, []int) {
	return file_proto_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTransactionRequest) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

type CreateTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CreateTransactionResponse) Reset() {
	*x = CreateTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTransactionResponse) ProtoMessage() {}

func (x *CreateTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTransactionResponse.ProtoReflect.Descriptor instead.
func (*CreateTransactionResponse) Descriptor() ([]byte, []int) {
	return file_proto_transaction_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTransactionResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateTransactionResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type DeleteTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTransactionRequest) Reset() {
	*x = DeleteTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transaction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTransactionRequest) ProtoMessage() {}

func (x *DeleteTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transaction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTransactionRequest.ProtoReflect.Descriptor instead.
func (*DeleteTransactionRequest) Descriptor() ([]byte, []int) {
	return file_proto_transaction_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteTransactionRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeleteTransactionResponse) Reset() {
	*x = DeleteTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transaction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTransactionResponse) ProtoMessage() {}

func (x *DeleteTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transaction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTransactionResponse.ProtoReflect.Descriptor instead.
func (*DeleteTransactionResponse) Descriptor() ([]byte, []int) {
	return file_proto_transaction_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteTransactionResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DeleteTransactionResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type RecurringTransactionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transactions []*RecurringTransaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *RecurringTransactionsResponse) Reset() {
	*x = RecurringTransactionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transaction_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecurringTransactionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecurringTransactionsResponse) ProtoMessage() {}

func (x *RecurringTransactionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transaction_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecurringTransactionsResponse.ProtoReflect.Descriptor instead.
func (*RecurringTransactionsResponse) Descriptor() ([]byte, []int) {
	return file_proto_transaction_proto_rawDescGZIP(), []int{5}
}

func (x *RecurringTransactionsResponse) GetTransactions() []*RecurringTransaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type CreateRecurringTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *RecurringTransaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *CreateRecurringTransactionRequest) Reset() {
	*x = CreateRecurringTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transaction_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRecurringTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRecurringTransactionRequest) ProtoMessage() {}

func (x *CreateRecurringTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transaction_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRecurringTransactionRequest.ProtoReflect.Descriptor instead.
func (*CreateRecurringTransactionRequest) Descriptor() ([]byte, []int) {
	return file_proto_transaction_proto_rawDescGZIP(), []int{6}
}

func (x *CreateRecurringTransactionRequest) GetTransaction() *RecurringTransaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

type CreateRecurringTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CreateRecurringTransactionResponse) Reset() {
	*x = CreateRecurringTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transaction_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRecurringTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRecurringTransactionResponse) ProtoMessage() {}

func (x *CreateRecurringTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transaction_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRecurringTransactionResponse.ProtoReflect.Descriptor instead.
func (*CreateRecurringTransactionResponse) Descriptor() ([]byte, []int) {
	return file_proto_transaction_proto_rawDescGZIP(), []int{7}
}

func (x *CreateRecurringTransactionResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateRecurringTransactionResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type DeleteRecurringTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRecurringTransactionRequest) Reset() {
	*x = DeleteRecurringTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transaction_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRecurringTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRecurringTransactionRequest) ProtoMessage() {}

func (x *DeleteRecurringTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transaction_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRecurringTransactionRequest.ProtoReflect.Descriptor instead.
func (*DeleteRecurringTransactionRequest) Descriptor() ([]byte, []int) {
	return file_proto_transaction_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteRecurringTransactionRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteRecurringTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeleteRecurringTransactionResponse) Reset() {
	*x = DeleteRecurringTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transaction_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRecurringTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRecurringTransactionResponse) ProtoMessage() {}

func (x *DeleteRecurringTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transaction_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRecurringTransactionResponse.ProtoReflect.Descriptor instead.
func (*DeleteRecurringTransactionResponse) Descriptor() ([]byte, []int) {
	return file_proto_transaction_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteRecurringTransactionResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DeleteRecurringTransactionResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_proto_transaction_proto protoreflect.FileDescriptor

var file_proto_transaction_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x66, 0x69, 0x6e, 0x63, 0x61,
	0x6c, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4f, 0x0a, 0x14, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0c, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x66, 0x69, 0x6e, 0x63, 0x61, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x51, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35,
	0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x66, 0x69, 0x6e, 0x63, 0x61, 0x6c, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4b, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x2a, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4b,
	0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x61, 0x0a, 0x1d, 0x52,
	0x65, 0x63, 0x75, 0x72, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0c,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x66, 0x69, 0x6e, 0x63, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x63, 0x75,
	0x72, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x63,
	0x0a, 0x21, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x75, 0x72, 0x72, 0x69, 0x6e,
	0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x66, 0x69, 0x6e, 0x63, 0x61,
	0x6c, 0x2e, 0x52, 0x65, 0x63, 0x75, 0x72, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x54, 0x0a, 0x22, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63,
	0x75, 0x72, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x33, 0x0a, 0x21, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x63, 0x75, 0x72, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x54,
	0x0a, 0x22, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x63, 0x75, 0x72, 0x72, 0x69, 0x6e,
	0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x32, 0xdc, 0x04, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0c, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x66, 0x69, 0x6e, 0x63, 0x61, 0x6c, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x66, 0x69, 0x6e, 0x63, 0x61,
	0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x66, 0x69, 0x6e,
	0x63, 0x61, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x5a, 0x0a, 0x11, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x66, 0x69, 0x6e, 0x63, 0x61, 0x6c, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x66, 0x69, 0x6e, 0x63, 0x61, 0x6c, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x15, 0x52,
	0x65, 0x63, 0x75, 0x72, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x25, 0x2e, 0x66,
	0x69, 0x6e, 0x63, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x63, 0x75, 0x72, 0x72, 0x69, 0x6e, 0x67, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x75, 0x0a, 0x1a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x63, 0x75, 0x72, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x29, 0x2e, 0x66, 0x69, 0x6e, 0x63, 0x61, 0x6c, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x75, 0x72, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a,
	0x2e, 0x66, 0x69, 0x6e, 0x63, 0x61, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x63, 0x75, 0x72, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x75, 0x0a, 0x1a,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x63, 0x75, 0x72, 0x72, 0x69, 0x6e, 0x67, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x2e, 0x66, 0x69, 0x6e,
	0x63, 0x61, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x63, 0x75, 0x72, 0x72,
	0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x66, 0x69, 0x6e, 0x63, 0x61, 0x6c, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x63, 0x75, 0x72, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x3b, 0x66, 0x69, 0x6e, 0x63, 0x61, 0x6c,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_transaction_proto_rawDescOnce sync.Once
	file_proto_transaction_proto_rawDescData = file_proto_transaction_proto_rawDesc
)

func file_proto_transaction_proto_rawDescGZIP() []byte {
	file_proto_transaction_proto_rawDescOnce.Do(func() {
		file_proto_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_transaction_proto_rawDescData)
	})
	return file_proto_transaction_proto_rawDescData
}

var file_proto_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_transaction_proto_goTypes = []any{
	(*TransactionsResponse)(nil),               // 0: fincal.TransactionsResponse
	(*CreateTransactionRequest)(nil),           // 1: fincal.CreateTransactionRequest
	(*CreateTransactionResponse)(nil),          // 2: fincal.CreateTransactionResponse
	(*DeleteTransactionRequest)(nil),           // 3: fincal.DeleteTransactionRequest
	(*DeleteTransactionResponse)(nil),          // 4: fincal.DeleteTransactionResponse
	(*RecurringTransactionsResponse)(nil),      // 5: fincal.RecurringTransactionsResponse
	(*CreateRecurringTransactionRequest)(nil),  // 6: fincal.CreateRecurringTransactionRequest
	(*CreateRecurringTransactionResponse)(nil), // 7: fincal.CreateRecurringTransactionResponse
	(*DeleteRecurringTransactionRequest)(nil),  // 8: fincal.DeleteRecurringTransactionRequest
	(*DeleteRecurringTransactionResponse)(nil), // 9: fincal.DeleteRecurringTransactionResponse
	(*Transaction)(nil),                        // 10: fincal.Transaction
	(*RecurringTransaction)(nil),               // 11: fincal.RecurringTransaction
	(*emptypb.Empty)(nil),                      // 12: google.protobuf.Empty
}
var file_proto_transaction_proto_depIdxs = []int32{
	10, // 0: fincal.TransactionsResponse.transactions:type_name -> fincal.Transaction
	10, // 1: fincal.CreateTransactionRequest.transaction:type_name -> fincal.Transaction
	11, // 2: fincal.RecurringTransactionsResponse.transactions:type_name -> fincal.RecurringTransaction
	11, // 3: fincal.CreateRecurringTransactionRequest.transaction:type_name -> fincal.RecurringTransaction
	12, // 4: fincal.TransactionService.Transactions:input_type -> google.protobuf.Empty
	1,  // 5: fincal.TransactionService.createTransaction:input_type -> fincal.CreateTransactionRequest
	3,  // 6: fincal.TransactionService.deleteTransaction:input_type -> fincal.DeleteTransactionRequest
	12, // 7: fincal.TransactionService.RecurringTransactions:input_type -> google.protobuf.Empty
	6,  // 8: fincal.TransactionService.createRecurringTransaction:input_type -> fincal.CreateRecurringTransactionRequest
	8,  // 9: fincal.TransactionService.deleteRecurringTransaction:input_type -> fincal.DeleteRecurringTransactionRequest
	0,  // 10: fincal.TransactionService.Transactions:output_type -> fincal.TransactionsResponse
	2,  // 11: fincal.TransactionService.createTransaction:output_type -> fincal.CreateTransactionResponse
	4,  // 12: fincal.TransactionService.deleteTransaction:output_type -> fincal.DeleteTransactionResponse
	5,  // 13: fincal.TransactionService.RecurringTransactions:output_type -> fincal.RecurringTransactionsResponse
	7,  // 14: fincal.TransactionService.createRecurringTransaction:output_type -> fincal.CreateRecurringTransactionResponse
	9,  // 15: fincal.TransactionService.deleteRecurringTransaction:output_type -> fincal.DeleteRecurringTransactionResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_transaction_proto_init() }
func file_proto_transaction_proto_init() {
	if File_proto_transaction_proto != nil {
		return
	}
	file_proto_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_transaction_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TransactionsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_transaction_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateTransactionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_transaction_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateTransactionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_transaction_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteTransactionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_transaction_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteTransactionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_transaction_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*RecurringTransactionsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_transaction_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CreateRecurringTransactionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_transaction_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*CreateRecurringTransactionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_transaction_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteRecurringTransactionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_transaction_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteRecurringTransactionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_transaction_proto_goTypes,
		DependencyIndexes: file_proto_transaction_proto_depIdxs,
		MessageInfos:      file_proto_transaction_proto_msgTypes,
	}.Build()
	File_proto_transaction_proto = out.File
	file_proto_transaction_proto_rawDesc = nil
	file_proto_transaction_proto_goTypes = nil
	file_proto_transaction_proto_depIdxs = nil
}