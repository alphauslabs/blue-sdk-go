// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: api/accountgroup.proto

package api

import (
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

type AccountGroup struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The AccountGroup id.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The AccountGroup name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The attributes (key/value pair) of the AccountGroup.
	Metadata      []*KeyValue `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty"`
	Accounts      []*Account  `protobuf:"bytes,4,rep,name=accounts,proto3" json:"accounts,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AccountGroup) Reset() {
	*x = AccountGroup{}
	mi := &file_api_accountgroup_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccountGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountGroup) ProtoMessage() {}

func (x *AccountGroup) ProtoReflect() protoreflect.Message {
	mi := &file_api_accountgroup_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountGroup.ProtoReflect.Descriptor instead.
func (*AccountGroup) Descriptor() ([]byte, []int) {
	return file_api_accountgroup_proto_rawDescGZIP(), []int{0}
}

func (x *AccountGroup) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AccountGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AccountGroup) GetMetadata() []*KeyValue {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *AccountGroup) GetAccounts() []*Account {
	if x != nil {
		return x.Accounts
	}
	return nil
}

var File_api_accountgroup_proto protoreflect.FileDescriptor

var file_api_accountgroup_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x65, 0x79, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a,
	0x0c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x31, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x30, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x42, 0x59, 0x0a, 0x19, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x70, 0x69, 0x42, 0x14, 0x41, 0x70, 0x69, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62,
	0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_accountgroup_proto_rawDescOnce sync.Once
	file_api_accountgroup_proto_rawDescData = file_api_accountgroup_proto_rawDesc
)

func file_api_accountgroup_proto_rawDescGZIP() []byte {
	file_api_accountgroup_proto_rawDescOnce.Do(func() {
		file_api_accountgroup_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_accountgroup_proto_rawDescData)
	})
	return file_api_accountgroup_proto_rawDescData
}

var file_api_accountgroup_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_accountgroup_proto_goTypes = []any{
	(*AccountGroup)(nil), // 0: blueapi.api.AccountGroup
	(*KeyValue)(nil),     // 1: blueapi.api.KeyValue
	(*Account)(nil),      // 2: blueapi.api.Account
}
var file_api_accountgroup_proto_depIdxs = []int32{
	1, // 0: blueapi.api.AccountGroup.metadata:type_name -> blueapi.api.KeyValue
	2, // 1: blueapi.api.AccountGroup.accounts:type_name -> blueapi.api.Account
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_accountgroup_proto_init() }
func file_api_accountgroup_proto_init() {
	if File_api_accountgroup_proto != nil {
		return
	}
	file_api_keyvalue_proto_init()
	file_api_account_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_accountgroup_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_accountgroup_proto_goTypes,
		DependencyIndexes: file_api_accountgroup_proto_depIdxs,
		MessageInfos:      file_api_accountgroup_proto_msgTypes,
	}.Build()
	File_api_accountgroup_proto = out.File
	file_api_accountgroup_proto_rawDesc = nil
	file_api_accountgroup_proto_goTypes = nil
	file_api_accountgroup_proto_depIdxs = nil
}
