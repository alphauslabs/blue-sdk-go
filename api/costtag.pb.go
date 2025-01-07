// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: api/costtag.proto

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

type CostTag struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The costtag id.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The parent organization id.
	OrgId string `protobuf:"bytes,2,opt,name=orgId,proto3" json:"orgId,omitempty"`
	// The parent billing internal id.
	BillingInternalId string `protobuf:"bytes,3,opt,name=billingInternalId,proto3" json:"billingInternalId,omitempty"`
	// The vendor.
	Vendor string `protobuf:"bytes,4,opt,name=vendor,proto3" json:"vendor,omitempty"`
	// The account id.
	AccountId string `protobuf:"bytes,5,opt,name=accountId,proto3" json:"accountId,omitempty"`
	// The logic.(and/or)
	Logic string `protobuf:"bytes,6,opt,name=logic,proto3" json:"logic,omitempty"`
	// The attributes (key/value pair) of the costtag.
	Tags          []*KeyValue `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CostTag) Reset() {
	*x = CostTag{}
	mi := &file_api_costtag_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CostTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostTag) ProtoMessage() {}

func (x *CostTag) ProtoReflect() protoreflect.Message {
	mi := &file_api_costtag_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostTag.ProtoReflect.Descriptor instead.
func (*CostTag) Descriptor() ([]byte, []int) {
	return file_api_costtag_proto_rawDescGZIP(), []int{0}
}

func (x *CostTag) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CostTag) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *CostTag) GetBillingInternalId() string {
	if x != nil {
		return x.BillingInternalId
	}
	return ""
}

func (x *CostTag) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *CostTag) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *CostTag) GetLogic() string {
	if x != nil {
		return x.Logic
	}
	return ""
}

func (x *CostTag) GetTags() []*KeyValue {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_api_costtag_proto protoreflect.FileDescriptor

var file_api_costtag_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69,
	0x1a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x65, 0x79, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f,
	0x67, 0x69, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x63,
	0x12, 0x29, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4b, 0x65, 0x79,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x42, 0x54, 0x0a, 0x19, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x2e, 0x62, 0x6c, 0x75,
	0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x69, 0x42, 0x0f, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x73,
	0x74, 0x54, 0x61, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x75, 0x73, 0x6c, 0x61, 0x62,
	0x73, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_costtag_proto_rawDescOnce sync.Once
	file_api_costtag_proto_rawDescData = file_api_costtag_proto_rawDesc
)

func file_api_costtag_proto_rawDescGZIP() []byte {
	file_api_costtag_proto_rawDescOnce.Do(func() {
		file_api_costtag_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_costtag_proto_rawDescData)
	})
	return file_api_costtag_proto_rawDescData
}

var file_api_costtag_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_costtag_proto_goTypes = []any{
	(*CostTag)(nil),  // 0: blueapi.api.CostTag
	(*KeyValue)(nil), // 1: blueapi.api.KeyValue
}
var file_api_costtag_proto_depIdxs = []int32{
	1, // 0: blueapi.api.CostTag.tags:type_name -> blueapi.api.KeyValue
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_costtag_proto_init() }
func file_api_costtag_proto_init() {
	if File_api_costtag_proto != nil {
		return
	}
	file_api_keyvalue_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_costtag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_costtag_proto_goTypes,
		DependencyIndexes: file_api_costtag_proto_depIdxs,
		MessageInfos:      file_api_costtag_proto_msgTypes,
	}.Build()
	File_api_costtag_proto = out.File
	file_api_costtag_proto_rawDesc = nil
	file_api_costtag_proto_goTypes = nil
	file_api_costtag_proto_depIdxs = nil
}
