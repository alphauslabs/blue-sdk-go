// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: kvstore/v1/kvstore.proto

package kvstore

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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
	mi := &file_kvstore_v1_kvstore_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValue) ProtoMessage() {}

func (x *KeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_kvstore_v1_kvstore_proto_msgTypes[0]
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
	return file_kvstore_v1_kvstore_proto_rawDescGZIP(), []int{0}
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

// Request message for the KvStore.Scan rpc.
type ScanRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Optional. If non-empty, translates to SQL's `LIKE` operand.
	// Return all keys by default.
	Like          string `protobuf:"bytes,1,opt,name=like,proto3" json:"like,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ScanRequest) Reset() {
	*x = ScanRequest{}
	mi := &file_kvstore_v1_kvstore_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanRequest) ProtoMessage() {}

func (x *ScanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kvstore_v1_kvstore_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanRequest.ProtoReflect.Descriptor instead.
func (*ScanRequest) Descriptor() ([]byte, []int) {
	return file_kvstore_v1_kvstore_proto_rawDescGZIP(), []int{1}
}

func (x *ScanRequest) GetLike() string {
	if x != nil {
		return x.Like
	}
	return ""
}

// Request message for the KvStore.Read rpc.
type ReadRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. The key to read.
	Key           string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	mi := &file_kvstore_v1_kvstore_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kvstore_v1_kvstore_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_kvstore_v1_kvstore_proto_rawDescGZIP(), []int{2}
}

func (x *ReadRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// Request message for the KvStore.Delete rpc.
type DeleteRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. The key to delete. To delete all keys, set this to `-` (hyphen).
	// You can also prefix this with `like:` which translates to SQL's LIKE operator.
	// For example, `like:%pattern%`.
	Key           string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	mi := &file_kvstore_v1_kvstore_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kvstore_v1_kvstore_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_kvstore_v1_kvstore_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_kvstore_v1_kvstore_proto protoreflect.FileDescriptor

var file_kvstore_v1_kvstore_proto_rawDesc = string([]byte{
	0x0a, 0x18, 0x6b, 0x76, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x76, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x62, 0x6c, 0x75, 0x65,
	0x61, 0x70, 0x69, 0x2e, 0x6b, 0x76, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x08, 0x4b, 0x65, 0x79,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x21, 0x0a,
	0x0b, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x69, 0x6b, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6b, 0x65,
	0x22, 0x1f, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x22, 0x21, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x32, 0x8a, 0x04, 0x0a, 0x07, 0x4b, 0x76, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x12, 0x59, 0x0a, 0x04, 0x53, 0x63, 0x61, 0x6e, 0x12, 0x1f, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61,
	0x70, 0x69, 0x2e, 0x6b, 0x76, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63,
	0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x62, 0x6c, 0x75, 0x65,
	0x61, 0x70, 0x69, 0x2e, 0x6b, 0x76, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4b,
	0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12,
	0x08, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x30, 0x01, 0x12, 0x5d, 0x0a, 0x04, 0x52,
	0x65, 0x61, 0x64, 0x12, 0x1f, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x76,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x6b,
	0x76, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x76, 0x31, 0x2f,
	0x6b, 0x65, 0x79, 0x73, 0x2f, 0x7b, 0x6b, 0x65, 0x79, 0x7d, 0x12, 0x4d, 0x0a, 0x05, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x76,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x0e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x08, 0x3a, 0x01, 0x2a, 0x22, 0x03, 0x2f, 0x76, 0x31, 0x12, 0x5b, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x76,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x16,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x2a, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x65, 0x79, 0x73,
	0x2f, 0x7b, 0x6b, 0x65, 0x79, 0x7d, 0x1a, 0x98, 0x01, 0x92, 0x41, 0x94, 0x01, 0x12, 0x41, 0x28,
	0x42, 0x45, 0x54, 0x41, 0x29, 0x20, 0x4b, 0x76, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x20, 0x41, 0x50,
	0x49, 0x2e, 0x20, 0x42, 0x61, 0x73, 0x65, 0x20, 0x55, 0x52, 0x4c, 0x3a, 0x20, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2f, 0x6b, 0x76,
	0x1a, 0x4f, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x75, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x74,
	0x72, 0x65, 0x65, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6b, 0x76, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2f, 0x42, 0x51, 0x0a, 0x19, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x75, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x76, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x42, 0x0c,
	0x4b, 0x76, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x26, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73,
	0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x76, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_kvstore_v1_kvstore_proto_rawDescOnce sync.Once
	file_kvstore_v1_kvstore_proto_rawDescData []byte
)

func file_kvstore_v1_kvstore_proto_rawDescGZIP() []byte {
	file_kvstore_v1_kvstore_proto_rawDescOnce.Do(func() {
		file_kvstore_v1_kvstore_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_kvstore_v1_kvstore_proto_rawDesc), len(file_kvstore_v1_kvstore_proto_rawDesc)))
	})
	return file_kvstore_v1_kvstore_proto_rawDescData
}

var file_kvstore_v1_kvstore_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_kvstore_v1_kvstore_proto_goTypes = []any{
	(*KeyValue)(nil),      // 0: blueapi.kvstore.v1.KeyValue
	(*ScanRequest)(nil),   // 1: blueapi.kvstore.v1.ScanRequest
	(*ReadRequest)(nil),   // 2: blueapi.kvstore.v1.ReadRequest
	(*DeleteRequest)(nil), // 3: blueapi.kvstore.v1.DeleteRequest
	(*emptypb.Empty)(nil), // 4: google.protobuf.Empty
}
var file_kvstore_v1_kvstore_proto_depIdxs = []int32{
	1, // 0: blueapi.kvstore.v1.KvStore.Scan:input_type -> blueapi.kvstore.v1.ScanRequest
	2, // 1: blueapi.kvstore.v1.KvStore.Read:input_type -> blueapi.kvstore.v1.ReadRequest
	0, // 2: blueapi.kvstore.v1.KvStore.Write:input_type -> blueapi.kvstore.v1.KeyValue
	3, // 3: blueapi.kvstore.v1.KvStore.Delete:input_type -> blueapi.kvstore.v1.DeleteRequest
	0, // 4: blueapi.kvstore.v1.KvStore.Scan:output_type -> blueapi.kvstore.v1.KeyValue
	0, // 5: blueapi.kvstore.v1.KvStore.Read:output_type -> blueapi.kvstore.v1.KeyValue
	4, // 6: blueapi.kvstore.v1.KvStore.Write:output_type -> google.protobuf.Empty
	4, // 7: blueapi.kvstore.v1.KvStore.Delete:output_type -> google.protobuf.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kvstore_v1_kvstore_proto_init() }
func file_kvstore_v1_kvstore_proto_init() {
	if File_kvstore_v1_kvstore_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_kvstore_v1_kvstore_proto_rawDesc), len(file_kvstore_v1_kvstore_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kvstore_v1_kvstore_proto_goTypes,
		DependencyIndexes: file_kvstore_v1_kvstore_proto_depIdxs,
		MessageInfos:      file_kvstore_v1_kvstore_proto_msgTypes,
	}.Build()
	File_kvstore_v1_kvstore_proto = out.File
	file_kvstore_v1_kvstore_proto_goTypes = nil
	file_kvstore_v1_kvstore_proto_depIdxs = nil
}
