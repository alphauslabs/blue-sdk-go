// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/keyvalue.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type KeyValue struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KeyValue) Reset() {
	*x = KeyValue{}
	mi := &file_api_keyvalue_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValue) ProtoMessage() {}

func (x *KeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_api_keyvalue_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValue.ProtoReflect.Descriptor instead.
func (*KeyValue) Descriptor() ([]byte, []int) {
	return file_api_keyvalue_proto_rawDescGZIP(), []int{0}
}

func (x *KeyValue) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// KeyValueMap resource definition.
type KeyValueMap struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Key   string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Types that are valid to be assigned to Value:
	//
	//	*KeyValueMap_StringValue
	//	*KeyValueMap_BoolValue
	Value         isKeyValueMap_Value `protobuf_oneof:"value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KeyValueMap) Reset() {
	*x = KeyValueMap{}
	mi := &file_api_keyvalue_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KeyValueMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValueMap) ProtoMessage() {}

func (x *KeyValueMap) ProtoReflect() protoreflect.Message {
	mi := &file_api_keyvalue_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValueMap.ProtoReflect.Descriptor instead.
func (*KeyValueMap) Descriptor() ([]byte, []int) {
	return file_api_keyvalue_proto_rawDescGZIP(), []int{1}
}

func (x *KeyValueMap) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyValueMap) GetValue() isKeyValueMap_Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *KeyValueMap) GetStringValue() string {
	if x != nil {
		if x, ok := x.Value.(*KeyValueMap_StringValue); ok {
			return x.StringValue
		}
	}
	return ""
}

func (x *KeyValueMap) GetBoolValue() bool {
	if x != nil {
		if x, ok := x.Value.(*KeyValueMap_BoolValue); ok {
			return x.BoolValue
		}
	}
	return false
}

type isKeyValueMap_Value interface {
	isKeyValueMap_Value()
}

type KeyValueMap_StringValue struct {
	StringValue string `protobuf:"bytes,2,opt,name=stringValue,proto3,oneof"`
}

type KeyValueMap_BoolValue struct {
	BoolValue bool `protobuf:"varint,3,opt,name=boolValue,proto3,oneof"`
}

func (*KeyValueMap_StringValue) isKeyValueMap_Value() {}

func (*KeyValueMap_BoolValue) isKeyValueMap_Value() {}

var File_api_keyvalue_proto protoreflect.FileDescriptor

const file_api_keyvalue_proto_rawDesc = "" +
	"\n" +
	"\x12api/keyvalue.proto\x12\vblueapi.api\"2\n" +
	"\bKeyValue\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value\"l\n" +
	"\vKeyValueMap\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\"\n" +
	"\vstringValue\x18\x02 \x01(\tH\x00R\vstringValue\x12\x1e\n" +
	"\tboolValue\x18\x03 \x01(\bH\x00R\tboolValueB\a\n" +
	"\x05valueBU\n" +
	"\x19cloud.alphaus.blueapi.apiB\x10ApiKeyValueProtoZ&github.com/alphauslabs/blue-sdk-go/apib\x06proto3"

var (
	file_api_keyvalue_proto_rawDescOnce sync.Once
	file_api_keyvalue_proto_rawDescData []byte
)

func file_api_keyvalue_proto_rawDescGZIP() []byte {
	file_api_keyvalue_proto_rawDescOnce.Do(func() {
		file_api_keyvalue_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_keyvalue_proto_rawDesc), len(file_api_keyvalue_proto_rawDesc)))
	})
	return file_api_keyvalue_proto_rawDescData
}

var file_api_keyvalue_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_keyvalue_proto_goTypes = []any{
	(*KeyValue)(nil),    // 0: blueapi.api.KeyValue
	(*KeyValueMap)(nil), // 1: blueapi.api.KeyValueMap
}
var file_api_keyvalue_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_keyvalue_proto_init() }
func file_api_keyvalue_proto_init() {
	if File_api_keyvalue_proto != nil {
		return
	}
	file_api_keyvalue_proto_msgTypes[1].OneofWrappers = []any{
		(*KeyValueMap_StringValue)(nil),
		(*KeyValueMap_BoolValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_keyvalue_proto_rawDesc), len(file_api_keyvalue_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_keyvalue_proto_goTypes,
		DependencyIndexes: file_api_keyvalue_proto_depIdxs,
		MessageInfos:      file_api_keyvalue_proto_msgTypes,
	}.Build()
	File_api_keyvalue_proto = out.File
	file_api_keyvalue_proto_goTypes = nil
	file_api_keyvalue_proto_depIdxs = nil
}
