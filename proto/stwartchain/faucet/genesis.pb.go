// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: stwartchain/faucet/genesis.proto

package faucet

import (
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// GenesisState defines the faucet module's genesis state.
type GenesisState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// params defines all the parameters of the module.
	Params      *Params   `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	TokensList  []*Tokens `protobuf:"bytes,2,rep,name=tokensList,proto3" json:"tokensList,omitempty"`
	TokensCount uint64    `protobuf:"varint,3,opt,name=tokensCount,proto3" json:"tokensCount,omitempty"`
}

func (x *GenesisState) Reset() {
	*x = GenesisState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stwartchain_faucet_genesis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisState) ProtoMessage() {}

func (x *GenesisState) ProtoReflect() protoreflect.Message {
	mi := &file_stwartchain_faucet_genesis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenesisState.ProtoReflect.Descriptor instead.
func (*GenesisState) Descriptor() ([]byte, []int) {
	return file_stwartchain_faucet_genesis_proto_rawDescGZIP(), []int{0}
}

func (x *GenesisState) GetParams() *Params {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *GenesisState) GetTokensList() []*Tokens {
	if x != nil {
		return x.TokensList
	}
	return nil
}

func (x *GenesisState) GetTokensCount() uint64 {
	if x != nil {
		return x.TokensCount
	}
	return 0
}

var File_stwartchain_faucet_genesis_proto protoreflect.FileDescriptor

var file_stwartchain_faucet_genesis_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x66, 0x61,
	0x75, 0x63, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x73, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e,
	0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x1a, 0x11, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2f, 0x61, 0x6d,
	0x69, 0x6e, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x73, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x66, 0x61, 0x75,
	0x63, 0x65, 0x74, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x73, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x66, 0x61,
	0x75, 0x63, 0x65, 0x74, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb1, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x2e, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x09,
	0xc8, 0xde, 0x1f, 0x00, 0xa8, 0xe7, 0xb0, 0x2a, 0x01, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x40, 0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x2e, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x61, 0x6c, 0x77, 0x61, 0x72, 0x74, 0x2d, 0x61, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x6d, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x69,
	0x73, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x77, 0x61, 0x72, 0x74,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stwartchain_faucet_genesis_proto_rawDescOnce sync.Once
	file_stwartchain_faucet_genesis_proto_rawDescData = file_stwartchain_faucet_genesis_proto_rawDesc
)

func file_stwartchain_faucet_genesis_proto_rawDescGZIP() []byte {
	file_stwartchain_faucet_genesis_proto_rawDescOnce.Do(func() {
		file_stwartchain_faucet_genesis_proto_rawDescData = protoimpl.X.CompressGZIP(file_stwartchain_faucet_genesis_proto_rawDescData)
	})
	return file_stwartchain_faucet_genesis_proto_rawDescData
}

var file_stwartchain_faucet_genesis_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_stwartchain_faucet_genesis_proto_goTypes = []interface{}{
	(*GenesisState)(nil), // 0: stwartchain.faucet.GenesisState
	(*Params)(nil),       // 1: stwartchain.faucet.Params
	(*Tokens)(nil),       // 2: stwartchain.faucet.Tokens
}
var file_stwartchain_faucet_genesis_proto_depIdxs = []int32{
	1, // 0: stwartchain.faucet.GenesisState.params:type_name -> stwartchain.faucet.Params
	2, // 1: stwartchain.faucet.GenesisState.tokensList:type_name -> stwartchain.faucet.Tokens
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_stwartchain_faucet_genesis_proto_init() }
func file_stwartchain_faucet_genesis_proto_init() {
	if File_stwartchain_faucet_genesis_proto != nil {
		return
	}
	file_stwartchain_faucet_params_proto_init()
	file_stwartchain_faucet_tokens_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_stwartchain_faucet_genesis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenesisState); i {
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
			RawDescriptor: file_stwartchain_faucet_genesis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_stwartchain_faucet_genesis_proto_goTypes,
		DependencyIndexes: file_stwartchain_faucet_genesis_proto_depIdxs,
		MessageInfos:      file_stwartchain_faucet_genesis_proto_msgTypes,
	}.Build()
	File_stwartchain_faucet_genesis_proto = out.File
	file_stwartchain_faucet_genesis_proto_rawDesc = nil
	file_stwartchain_faucet_genesis_proto_goTypes = nil
	file_stwartchain_faucet_genesis_proto_depIdxs = nil
}
