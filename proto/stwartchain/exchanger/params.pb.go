// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: stwartchain/exchanger/params.proto

package exchanger

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

// Params defines the parameters for the module.
type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stwartchain_exchanger_params_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

func (x *Params) ProtoReflect() protoreflect.Message {
	mi := &file_stwartchain_exchanger_params_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_stwartchain_exchanger_params_proto_rawDescGZIP(), []int{0}
}

var File_stwartchain_exchanger_params_proto protoreflect.FileDescriptor

var file_stwartchain_exchanger_params_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x65, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x73, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x1a, 0x11, 0x61, 0x6d, 0x69,
	0x6e, 0x6f, 0x2f, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14,
	0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x3a, 0x27,
	0xe8, 0xa0, 0x1f, 0x01, 0x8a, 0xe7, 0xb0, 0x2a, 0x1e, 0x73, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2f, 0x78, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x72,
	0x2f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x61, 0x6c, 0x77, 0x61, 0x72, 0x74, 0x2d, 0x61,
	0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x6d, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x2f, 0x63, 0x61, 0x6c,
	0x6c, 0x69, 0x73, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x77, 0x61,
	0x72, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stwartchain_exchanger_params_proto_rawDescOnce sync.Once
	file_stwartchain_exchanger_params_proto_rawDescData = file_stwartchain_exchanger_params_proto_rawDesc
)

func file_stwartchain_exchanger_params_proto_rawDescGZIP() []byte {
	file_stwartchain_exchanger_params_proto_rawDescOnce.Do(func() {
		file_stwartchain_exchanger_params_proto_rawDescData = protoimpl.X.CompressGZIP(file_stwartchain_exchanger_params_proto_rawDescData)
	})
	return file_stwartchain_exchanger_params_proto_rawDescData
}

var file_stwartchain_exchanger_params_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_stwartchain_exchanger_params_proto_goTypes = []interface{}{
	(*Params)(nil), // 0: stwartchain.exchanger.Params
}
var file_stwartchain_exchanger_params_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stwartchain_exchanger_params_proto_init() }
func file_stwartchain_exchanger_params_proto_init() {
	if File_stwartchain_exchanger_params_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stwartchain_exchanger_params_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params); i {
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
			RawDescriptor: file_stwartchain_exchanger_params_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_stwartchain_exchanger_params_proto_goTypes,
		DependencyIndexes: file_stwartchain_exchanger_params_proto_depIdxs,
		MessageInfos:      file_stwartchain_exchanger_params_proto_msgTypes,
	}.Build()
	File_stwartchain_exchanger_params_proto = out.File
	file_stwartchain_exchanger_params_proto_rawDesc = nil
	file_stwartchain_exchanger_params_proto_goTypes = nil
	file_stwartchain_exchanger_params_proto_depIdxs = nil
}
