// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: api/ripple/payer.proto

package ripple

import (
	api "github.com/alphauslabs/blue-sdk-go/api"
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

// Payer resource definition.
type Payer struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The payer account id.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The payer account name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The attributes (key/value pair) of the account.
	Metadata []*api.KeyValue `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty"`
	// List of all members under this payer.
	Members       []*api.Account `protobuf:"bytes,4,rep,name=members,proto3" json:"members,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Payer) Reset() {
	*x = Payer{}
	mi := &file_api_ripple_payer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Payer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payer) ProtoMessage() {}

func (x *Payer) ProtoReflect() protoreflect.Message {
	mi := &file_api_ripple_payer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payer.ProtoReflect.Descriptor instead.
func (*Payer) Descriptor() ([]byte, []int) {
	return file_api_ripple_payer_proto_rawDescGZIP(), []int{0}
}

func (x *Payer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Payer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Payer) GetMetadata() []*api.KeyValue {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Payer) GetMembers() []*api.Account {
	if x != nil {
		return x.Members
	}
	return nil
}

var File_api_ripple_payer_proto protoreflect.FileDescriptor

var file_api_ripple_payer_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x61, 0x79,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x1a, 0x11, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x12, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x65, 0x79, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x05, 0x50, 0x61, 0x79, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x31, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x42, 0x66, 0x0a, 0x20, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x42, 0x13, 0x41, 0x70, 0x69, 0x52, 0x69, 0x70,
	0x70, 0x6c, 0x65, 0x50, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x2d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75,
	0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_ripple_payer_proto_rawDescOnce sync.Once
	file_api_ripple_payer_proto_rawDescData []byte
)

func file_api_ripple_payer_proto_rawDescGZIP() []byte {
	file_api_ripple_payer_proto_rawDescOnce.Do(func() {
		file_api_ripple_payer_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_ripple_payer_proto_rawDesc), len(file_api_ripple_payer_proto_rawDesc)))
	})
	return file_api_ripple_payer_proto_rawDescData
}

var file_api_ripple_payer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_ripple_payer_proto_goTypes = []any{
	(*Payer)(nil),        // 0: blueapi.api.ripple.Payer
	(*api.KeyValue)(nil), // 1: blueapi.api.KeyValue
	(*api.Account)(nil),  // 2: blueapi.api.Account
}
var file_api_ripple_payer_proto_depIdxs = []int32{
	1, // 0: blueapi.api.ripple.Payer.metadata:type_name -> blueapi.api.KeyValue
	2, // 1: blueapi.api.ripple.Payer.members:type_name -> blueapi.api.Account
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_ripple_payer_proto_init() }
func file_api_ripple_payer_proto_init() {
	if File_api_ripple_payer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_ripple_payer_proto_rawDesc), len(file_api_ripple_payer_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_ripple_payer_proto_goTypes,
		DependencyIndexes: file_api_ripple_payer_proto_depIdxs,
		MessageInfos:      file_api_ripple_payer_proto_msgTypes,
	}.Build()
	File_api_ripple_payer_proto = out.File
	file_api_ripple_payer_proto_goTypes = nil
	file_api_ripple_payer_proto_depIdxs = nil
}
