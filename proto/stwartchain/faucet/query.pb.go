// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: stwartchain/faucet/query.proto

package faucet

import (
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QueryParamsRequest) Reset() {
	*x = QueryParamsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stwartchain_faucet_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryParamsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryParamsRequest) ProtoMessage() {}

func (x *QueryParamsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stwartchain_faucet_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryParamsRequest.ProtoReflect.Descriptor instead.
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return file_stwartchain_faucet_query_proto_rawDescGZIP(), []int{0}
}

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// params holds all the parameters of this module.
	Params *Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *QueryParamsResponse) Reset() {
	*x = QueryParamsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stwartchain_faucet_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryParamsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryParamsResponse) ProtoMessage() {}

func (x *QueryParamsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stwartchain_faucet_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryParamsResponse.ProtoReflect.Descriptor instead.
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return file_stwartchain_faucet_query_proto_rawDescGZIP(), []int{1}
}

func (x *QueryParamsResponse) GetParams() *Params {
	if x != nil {
		return x.Params
	}
	return nil
}

type QueryGetTokensRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *QueryGetTokensRequest) Reset() {
	*x = QueryGetTokensRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stwartchain_faucet_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryGetTokensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryGetTokensRequest) ProtoMessage() {}

func (x *QueryGetTokensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stwartchain_faucet_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryGetTokensRequest.ProtoReflect.Descriptor instead.
func (*QueryGetTokensRequest) Descriptor() ([]byte, []int) {
	return file_stwartchain_faucet_query_proto_rawDescGZIP(), []int{2}
}

func (x *QueryGetTokensRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type QueryGetTokensResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tokens *Tokens `protobuf:"bytes,1,opt,name=Tokens,proto3" json:"Tokens,omitempty"`
}

func (x *QueryGetTokensResponse) Reset() {
	*x = QueryGetTokensResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stwartchain_faucet_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryGetTokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryGetTokensResponse) ProtoMessage() {}

func (x *QueryGetTokensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stwartchain_faucet_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryGetTokensResponse.ProtoReflect.Descriptor instead.
func (*QueryGetTokensResponse) Descriptor() ([]byte, []int) {
	return file_stwartchain_faucet_query_proto_rawDescGZIP(), []int{3}
}

func (x *QueryGetTokensResponse) GetTokens() *Tokens {
	if x != nil {
		return x.Tokens
	}
	return nil
}

type QueryAllTokensRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryAllTokensRequest) Reset() {
	*x = QueryAllTokensRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stwartchain_faucet_query_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAllTokensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAllTokensRequest) ProtoMessage() {}

func (x *QueryAllTokensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stwartchain_faucet_query_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAllTokensRequest.ProtoReflect.Descriptor instead.
func (*QueryAllTokensRequest) Descriptor() ([]byte, []int) {
	return file_stwartchain_faucet_query_proto_rawDescGZIP(), []int{4}
}

func (x *QueryAllTokensRequest) GetPagination() *query.PageRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type QueryAllTokensResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tokens     []*Tokens           `protobuf:"bytes,1,rep,name=Tokens,proto3" json:"Tokens,omitempty"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryAllTokensResponse) Reset() {
	*x = QueryAllTokensResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stwartchain_faucet_query_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAllTokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAllTokensResponse) ProtoMessage() {}

func (x *QueryAllTokensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stwartchain_faucet_query_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAllTokensResponse.ProtoReflect.Descriptor instead.
func (*QueryAllTokensResponse) Descriptor() ([]byte, []int) {
	return file_stwartchain_faucet_query_proto_rawDescGZIP(), []int{5}
}

func (x *QueryAllTokensResponse) GetTokens() []*Tokens {
	if x != nil {
		return x.Tokens
	}
	return nil
}

func (x *QueryAllTokensResponse) GetPagination() *query.PageResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_stwartchain_faucet_query_proto protoreflect.FileDescriptor

var file_stwartchain_faucet_query_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x66, 0x61,
	0x75, 0x63, 0x65, 0x74, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x73, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x66, 0x61,
	0x75, 0x63, 0x65, 0x74, 0x1a, 0x11, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2f, 0x61, 0x6d, 0x69, 0x6e,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x73, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2f, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2f, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x73, 0x74, 0x77, 0x61, 0x72, 0x74,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x54, 0x0a, 0x13, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2e, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x42, 0x09, 0xc8, 0xde, 0x1f, 0x00, 0xa8, 0xe7, 0xb0, 0x2a, 0x01, 0x52, 0x06, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x27, 0x0a, 0x15, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x65,
	0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x52,
	0x0a, 0x16, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x74, 0x77, 0x61, 0x72,
	0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x06, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x22, 0x5f, 0x0a, 0x15, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6c, 0x6c, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x9b, 0x01, 0x0a, 0x16, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6c, 0x6c,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38,
	0x0a, 0x06, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x73, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x66, 0x61, 0x75,
	0x63, 0x65, 0x74, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00,
	0x52, 0x06, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x47, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x32, 0xb6, 0x03, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x86, 0x01, 0x0a, 0x06,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x26, 0x2e, 0x73, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2e, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x73, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x66, 0x61, 0x75,
	0x63, 0x65, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x12,
	0x23, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x73, 0x74, 0x77, 0x61, 0x72, 0x74,
	0x2d, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2f, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x12, 0x91, 0x01, 0x0a, 0x06, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12,
	0x29, 0x2e, 0x73, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x66, 0x61,
	0x75, 0x63, 0x65, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x74, 0x77,
	0x61, 0x72, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x12, 0x28,
	0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x73, 0x74, 0x77, 0x61, 0x72, 0x74, 0x2d,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x8f, 0x01, 0x0a, 0x09, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x41, 0x6c, 0x6c, 0x12, 0x29, 0x2e, 0x73, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2e, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e,
	0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6c, 0x6c, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x25, 0x12, 0x23, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f,
	0x73, 0x74, 0x77, 0x61, 0x72, 0x74, 0x2d, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x66, 0x61, 0x75,
	0x63, 0x65, 0x74, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x61, 0x6c, 0x77, 0x61, 0x72,
	0x74, 0x2d, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x6d, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x2f,
	0x63, 0x61, 0x6c, 0x6c, 0x69, 0x73, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73,
	0x74, 0x77, 0x61, 0x72, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x66, 0x61, 0x75, 0x63, 0x65,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stwartchain_faucet_query_proto_rawDescOnce sync.Once
	file_stwartchain_faucet_query_proto_rawDescData = file_stwartchain_faucet_query_proto_rawDesc
)

func file_stwartchain_faucet_query_proto_rawDescGZIP() []byte {
	file_stwartchain_faucet_query_proto_rawDescOnce.Do(func() {
		file_stwartchain_faucet_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_stwartchain_faucet_query_proto_rawDescData)
	})
	return file_stwartchain_faucet_query_proto_rawDescData
}

var file_stwartchain_faucet_query_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_stwartchain_faucet_query_proto_goTypes = []interface{}{
	(*QueryParamsRequest)(nil),     // 0: stwartchain.faucet.QueryParamsRequest
	(*QueryParamsResponse)(nil),    // 1: stwartchain.faucet.QueryParamsResponse
	(*QueryGetTokensRequest)(nil),  // 2: stwartchain.faucet.QueryGetTokensRequest
	(*QueryGetTokensResponse)(nil), // 3: stwartchain.faucet.QueryGetTokensResponse
	(*QueryAllTokensRequest)(nil),  // 4: stwartchain.faucet.QueryAllTokensRequest
	(*QueryAllTokensResponse)(nil), // 5: stwartchain.faucet.QueryAllTokensResponse
	(*Params)(nil),                 // 6: stwartchain.faucet.Params
	(*Tokens)(nil),                 // 7: stwartchain.faucet.Tokens
	(*query.PageRequest)(nil),      // 8: cosmos.base.query.v1beta1.PageRequest
	(*query.PageResponse)(nil),     // 9: cosmos.base.query.v1beta1.PageResponse
}
var file_stwartchain_faucet_query_proto_depIdxs = []int32{
	6, // 0: stwartchain.faucet.QueryParamsResponse.params:type_name -> stwartchain.faucet.Params
	7, // 1: stwartchain.faucet.QueryGetTokensResponse.Tokens:type_name -> stwartchain.faucet.Tokens
	8, // 2: stwartchain.faucet.QueryAllTokensRequest.pagination:type_name -> cosmos.base.query.v1beta1.PageRequest
	7, // 3: stwartchain.faucet.QueryAllTokensResponse.Tokens:type_name -> stwartchain.faucet.Tokens
	9, // 4: stwartchain.faucet.QueryAllTokensResponse.pagination:type_name -> cosmos.base.query.v1beta1.PageResponse
	0, // 5: stwartchain.faucet.Query.Params:input_type -> stwartchain.faucet.QueryParamsRequest
	2, // 6: stwartchain.faucet.Query.Tokens:input_type -> stwartchain.faucet.QueryGetTokensRequest
	4, // 7: stwartchain.faucet.Query.TokensAll:input_type -> stwartchain.faucet.QueryAllTokensRequest
	1, // 8: stwartchain.faucet.Query.Params:output_type -> stwartchain.faucet.QueryParamsResponse
	3, // 9: stwartchain.faucet.Query.Tokens:output_type -> stwartchain.faucet.QueryGetTokensResponse
	5, // 10: stwartchain.faucet.Query.TokensAll:output_type -> stwartchain.faucet.QueryAllTokensResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_stwartchain_faucet_query_proto_init() }
func file_stwartchain_faucet_query_proto_init() {
	if File_stwartchain_faucet_query_proto != nil {
		return
	}
	file_stwartchain_faucet_params_proto_init()
	file_stwartchain_faucet_tokens_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_stwartchain_faucet_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryParamsRequest); i {
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
		file_stwartchain_faucet_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryParamsResponse); i {
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
		file_stwartchain_faucet_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryGetTokensRequest); i {
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
		file_stwartchain_faucet_query_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryGetTokensResponse); i {
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
		file_stwartchain_faucet_query_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAllTokensRequest); i {
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
		file_stwartchain_faucet_query_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAllTokensResponse); i {
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
			RawDescriptor: file_stwartchain_faucet_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stwartchain_faucet_query_proto_goTypes,
		DependencyIndexes: file_stwartchain_faucet_query_proto_depIdxs,
		MessageInfos:      file_stwartchain_faucet_query_proto_msgTypes,
	}.Build()
	File_stwartchain_faucet_query_proto = out.File
	file_stwartchain_faucet_query_proto_rawDesc = nil
	file_stwartchain_faucet_query_proto_goTypes = nil
	file_stwartchain_faucet_query_proto_depIdxs = nil
}
